package main

import (
	"bookstore/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBooksRoute(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestBooksbyISBNRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/9781612680194", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.MatchRegex(t, w.Body.String(), "Rich Dad Poor Dad")
}

func TestPostBookRoute(t *testing.T) {
	router := setupRouter()

	book := models.Book{
		ISBN:   "1234567891012",
		Author: "Santosh Kumar",
		Title:  "Hello World",
	}

	body, _ := json.Marshal(book)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.MatchRegex(t, w.Body.String(), "Hello World")
}

func TestDeleteBookRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/9781612680194", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 202, w.Code)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/books/9781612680194", nil)
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 404, w2.Code)
}

func TestUpdateBookRoute(t *testing.T) {
	router := setupRouter()

	newBook := &models.Book{
		ISBN:   "9780593419052",
		Title:  "NewTitle",
		Author: "NewAuthor",
	}

	b, _ := json.Marshal(newBook)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books", bytes.NewReader(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.MatchRegex(t, w.Body.String(), "9780593419052")
	assert.MatchRegex(t, w.Body.String(), "NewTitle")
	assert.MatchRegex(t, w.Body.String(), "NewAuthor")
}
