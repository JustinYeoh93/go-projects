package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type movieHanlder struct {
	service IMovieService
}

func NewMovieHandler(ms IMovieService) *movieHanlder {
	return &movieHanlder{service: ms}
}

func (mh *movieHanlder) UpdateMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))

	var movie movieModel
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = mh.service.UpdateMovie(id, movie)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Successfully updated"))
}

func (mh *movieHanlder) GetMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	movies := mh.service.GetMovies()

	b, _ := json.Marshal(movies)
	w.WriteHeader(200)
	w.Write(b)
}
