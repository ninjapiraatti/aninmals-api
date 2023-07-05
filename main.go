package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", getBooks)

	books = append(books, Book{ID: "1", Title: "Book One", Author: "Author One"})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: "Author Two"})

	http.ListenAndServe(":8000", nil)
}
