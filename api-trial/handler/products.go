package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AlpKarar/package/tree/master/api-trial/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


/*
func getProducts() []*data.Product {
	return data.ProductList
}

func (p *Products) ServeHTTP(rw  http.ResponseWriter, r *http.Request) {
	products := getProducts()
	encodedProducts, err := json.Marshal(products)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(encodedProducts)
}
*/

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Failed to encode", http.StatusInternalServerError)
	} 
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Request got triggered ")

	// Middleware handles the part below
	/*
	newProd := &data.Product{}
	err := newProd.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	*/

	newProd := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&newProd)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT Request got triggered")

	/*
	newProd := &data.Product{}
	err := newProd.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to read JSON", http.StatusBadRequest)
		return
	}
	
	rx, _ := regexp.Compile("/([0-9]+)")
	g := rx.FindAllStringSubmatch(r.URL.Path, -1)

	if len(g) != 1 {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)
		return
	}

	if len(g[0]) != 2 {
		http.Error(rw, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(g[0][1])
	*/

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "String convert fault", http.StatusBadRequest)
		return
	}

	newProd := r.Context().Value(KeyProduct{}).(data.Product)

	//data.UpdateProduct(id, newProd)
	data.UpdateProduct(id, &newProd)
}

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("DELETE request got trigerred")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "String convert fault", http.StatusBadRequest)
	}

	errDelete := data.DeleteProduct(id)

	if errDelete != nil {
		http.Error(rw , errDelete.Error(), http.StatusBadRequest)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error in reading product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()

		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}