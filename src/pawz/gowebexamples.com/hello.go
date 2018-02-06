package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo gess, request : %s \n", r.URL.Path[1:])
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//})
	http.HandleFunc("/handler", handler)

	http.ListenAndServe(":12345", nil)
}
