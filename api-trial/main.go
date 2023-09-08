package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlpKarar/package/tree/master/api-trial/handler"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const PORT = ":9090"

func main() {
	l := log.New(os.Stdout, "api-trial ", log.LstdFlags)
	//hh := handler.NewHello(l)
	//gh := handler.NewGoodbye(l)
	ph := handler.NewProducts(l)

	gsm := mux.NewRouter()

	getRouter := gsm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)

	postRouter := gsm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", ph.AddProducts)
	postRouter.Use(ph.MiddleWareProductValidation)

	putRouter := gsm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddleWareProductValidation)

	deleteRouter := gsm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	//getRouter.Handle("/goodbye", gh)
	//mx.Handle("/", hh)

	crs := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := &http.Server{
		Addr: PORT,
		Handler: crs(gsm),
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	go func () {
		l.Println("Server intialized running on port:", PORT[1:])

		err := server.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan

	l.Println("Received terminate, proper shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	server.Shutdown(tc)
}