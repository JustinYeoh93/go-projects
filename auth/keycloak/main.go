package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"example/src/config"
	"example/src/controllers"
	"example/src/services"
)

func main() {
	config.DbConnect()
	services.InitializeOauthServer()

	r := mux.NewRouter().StrictSlash(true)
	controllers.EventController{}.RegisterRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
