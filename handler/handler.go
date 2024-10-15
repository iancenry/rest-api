package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book

var jsonResponse = struct {
	key   string
	value string
}{
	key:   "Content-Type",
	value: "application/json",
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	for _, book := range Books {
		if book.ID == id {
			w.Header().Set(jsonResponse.key, jsonResponse.value)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)

}

// curl -X POST http://localhost:8080/books -d '{"title":"The Rebel","author":"Camus"}' -H "Content-Type: application/json"
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.Author == "" || book.Title == "" {
		http.Error(w, "Please provide a title and an author", http.StatusBadRequest)
		return
	}
	book.ID = len(Books) + 1
	Books = append(Books, book)
	w.Header().Set(jsonResponse.key, jsonResponse.value)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set(jsonResponse.key, jsonResponse.value)
	w.WriteHeader(http.StatusOK)

	encoder.Encode(Books)
}

// encode a JSON payload on request and decode a JSON body from the response
