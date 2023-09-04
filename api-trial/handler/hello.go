package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
} 

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Data couldn't be read", http.StatusBadRequest)
	}

	fmt.Printf("Data: %s", d)
	fmt.Fprintf(os.Stdout, "Data: %s", d)
} 