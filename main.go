package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iancenry/go-rest-api/database"
	"github.com/iancenry/go-rest-api/handler"
	"github.com/iancenry/go-rest-api/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	db = database.ConnectDB()
	defer db.Close()

	
	router := *mux.NewRouter()

	handler.Books = append(handler.Books, handler.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})

	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/books/", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.Handle("/books", middleware.IsAuthenticated(http.HandlerFunc(handler.AddBook))).Methods("POST")

	fmt.Printf("Starting server on port %s...\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), &router))

}
