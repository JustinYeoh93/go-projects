package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startRouter(ctrl *userController) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", ctrl.getAllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", ctrl.deleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", ctrl.updateUser).Methods("PUT")
	router.HandleFunc("/user/{name}/{email}", ctrl.newUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	log.Println("starting server at 8080")

	database := newDatabase()
	database.initialMigration()

	userController := newUserController(database)

	startRouter(userController)
}
