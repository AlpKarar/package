package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlpKarar/package/tree/master/api-trial/handler"
)

const PORT = ":9090"

func main() {
	l := log.New(os.Stdout, "api-trial", log.LstdFlags)
	hh := handler.NewHello(l)
	gh := handler.NewGoodbye(l)

	mx := http.NewServeMux()

	mx.Handle("/", hh)
	mx.Handle("/goodbye", gh)

	server := &http.Server{
		Addr: PORT,
		Handler: mx,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
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