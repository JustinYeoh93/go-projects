package main

import (
	"log"
	"net/http"

	"github.com/JustinYeoh93/go-projects/fcc-projects/curd-api/pkg/routes"
	"github.com/gorilla/mux"
)

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
