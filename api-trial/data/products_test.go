package data

import "testing"

func TestChecksValidation(t *testing.T) {
	pList := []*Product{
		&Product{
			Name: "Tea",
			Price: 1.00,
			SKU: "aaaa-aaa-aaa",
		},
		&Product{
			Name: "Coffee 2",
			Price: 3.75,
			SKU: "aax-34",
		},
	}

	for _, p := range pList {
		err := p.Validate()

		if err != nil {
			t.Fatal(err)
		}
	}
}