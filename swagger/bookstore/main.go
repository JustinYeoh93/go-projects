package main

import (
	controller "bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/books", controller.GetBooks)
	r.PUT("/books", controller.UpdateBook)
	r.GET("/books/:isbn", controller.GetBookByISBN)
	r.POST("/books", controller.AddBook)
	r.DELETE("books/:isbn", controller.DeleteBookByISBN)
	return r
}

func main() {
	router := setupRouter()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
