package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const PORT string = ":9090"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)
		//fmt.Println("Data: ", string(d))
		fmt.Fprintf(os.Stdout, "Hello %s\n", d)

		rw.Write([]byte("HOME PAGE"))
	})

	fmt.Println("Server started running on port:", PORT[1:])
	http.ListenAndServe(PORT, nil)
}