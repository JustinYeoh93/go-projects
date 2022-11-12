package controllers

import "github.com/gorilla/mux"

type Controllers interface {
	RegisterRoutes(router *mux.Router)
}
