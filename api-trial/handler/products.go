package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/AlpKarar/package/tree/master/api-trial/data"
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

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST Request got triggered ")

	newProd := &data.Product{}
	err := newProd.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(newProd)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "PUT Request got triggered")

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

	data.UpdateProduct(id, newProd)
}