package main

import (
	"fmt"
	"net/http"
	"os"

	"aninmals-api/pkg/handler"
	"aninmals-api/pkg/repository"
	"aninmals-api/pkg/router"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func testFunction() int {
	return 666
}

func main() {
	err := godotenv.Load("../../.env")
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

	aninmalRepository := repository.AninmalRepository{}
	aninmalHandler := handler.AninmalHandler{AninmalRepository: aninmalRepository}

	r := router.NewRouter(aninmalHandler)

	http.ListenAndServe(":8000", r)
}
