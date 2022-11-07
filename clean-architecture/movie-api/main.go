package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	repo := NewInmemoyMovieRepository()
	svc := NewDefaultMovieSerives(repo)
	hdlr := NewMovieHandler(svc)

	router := httprouter.New()

	router.GET("/movies", hdlr.GetMovie)
	router.PATCH("/movies/:id", hdlr.UpdateMovie)

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
