package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	pgDB := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: dbPassword,
		Database: "aninmals",
	})

	defer pgDB.Close()

	// test the connection
	_, connErr := pgDB.Exec("SELECT 1")
	if connErr != nil {
		fmt.Println(connErr)
		return
	}

	fmt.Println("Connected successfully!")

	http.HandleFunc("/books", getBooks)

	books = append(books, Book{ID: "1", Title: "Book One", Author: "Author One"})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: "Author Two"})

	http.ListenAndServe(":8000", nil)
}
