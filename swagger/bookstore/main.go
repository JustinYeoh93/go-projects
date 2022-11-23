package main

import (
	controller "bookstore/controllers"

	_ "bookstore/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gr := r.Group("/api/v1")
	gr.GET("/books", controller.GetBooks)
	gr.PUT("/books", controller.UpdateBook)
	gr.GET("/books/:isbn", controller.GetBookByISBN)
	gr.POST("/books", controller.AddBook)
	gr.DELETE("books/:isbn", controller.DeleteBookByISBN)
	return r
}

// @title Gingo Book Service
// @version 1.0
// @description A book management service API in Go using Gin framework.
// @termsOfService https://abc.com

// @contact.name Justin
// @contact.url https://facebook.com
// @contact.email justin@justinyeoh.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := setupRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
