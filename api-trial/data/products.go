package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       float32 `json:"price" validate:*gt=0`
	SKU         string `json:"sku" validate:"required,sku"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return ProductList
}

func AddProduct(p *Product) {
	p.ID = getNewProdId()
	ProductList = append(ProductList, p)
}

func UpdateProduct(id int, p *Product) {
	p.ID = id
	i, err := findProduct(id)

	if err != nil {
		fmt.Printf("Err: %v", err)
		return
	}

	ProductList[i] = p
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return i, nil
		}
	}
	
	return -1, ErrProductNotFound
}

func getNewProdId() int {
	return ProductList[len(ProductList) - 1].ID + 1
}

var ErrProductNotAllowedToDelete = fmt.Errorf("Product Not Allowed To Delete")

func DeleteProduct(id int) error {
	isId := false

	for _, prod := range ProductList {
		if prod.ID == id {
			isId = true
		}
	}

	if !isId {
		return ErrProductNotAllowedToDelete
	}

	tmpProductList := []*Product{}

	for _, prod := range ProductList {
		if prod.ID != id {
			tmpProductList = append(tmpProductList, prod)
		}
	}

	ProductList = tmpProductList

	return nil
}

var ProductList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}