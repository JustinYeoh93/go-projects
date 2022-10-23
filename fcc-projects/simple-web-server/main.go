package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// forms are not parsed by default
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "PraseForm() err: %v", err)
		return
	}

	name := r.Form.Get("name")
	address := r.Form.Get("address")

	fmt.Printf("a new form submission with name '%s' and address '%s'\n", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Hello there!")
}
