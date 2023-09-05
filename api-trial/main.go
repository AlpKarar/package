package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AlpKarar/package/tree/master/api-trial/handler"
)

func main() {
	l := log.New(os.Stdout, "api-trial", log.LstdFlags)
	hh := handler.NewHello(l)

	mx := http.NewServeMux()

	mx.Handle("/", hh)
	
	http.ListenAndServe(":9090", mx)
}