package controller

import (
	"bookstore/db"
	"bookstore/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, db.Books)
}

func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	for _, book := range db.Books {
		if book.ISBN == isbn {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, fmt.Sprintf("book with isbn '%s' not found", isbn))
}

func AddBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, "unable to parse JSON")
		return
	}

	db.Books = append(db.Books, book)

	c.JSON(http.StatusCreated, book)
}

func DeleteBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	for ind, book := range db.Books {
		if book.ISBN == isbn {
			db.Books = append(db.Books[:ind], db.Books[ind+1:]...)
			c.JSON(http.StatusAccepted, "deleted")
			return
		}
	}

	c.JSON(200, fmt.Sprintf("book with isbn '%s' not found", isbn))
}

func UpdateBook(c *gin.Context) {
	var newBook models.Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, "unable to parse json")
		return
	}

	for _, book := range db.Books {
		if newBook.ISBN == book.ISBN {
			book.Author = newBook.Author
			book.Title = newBook.Title
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, fmt.Sprintf("book with isbn '%s' not found", newBook.ISBN))
}
