package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iancenry/go-rest-api/handler"
	"github.com/iancenry/go-rest-api/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	router := *mux.NewRouter()

	handler.Books = append(handler.Books, handler.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})

	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/books/", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.Handle("/books", middleware.IsAuthenticated(http.HandlerFunc(handler.AddBook))).Methods("POST")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &router))

}
