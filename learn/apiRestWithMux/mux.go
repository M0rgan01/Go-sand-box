package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var books []Book

func main() {
	// init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{
		ID: createUuid(),
		Author: &Author{
			FirstName: "TestFirstName",
			LastName:  "TestLastName",
		},
		Title: "Harry potter",
	})

	books = append(books, Book{
		ID: createUuid(),
		Author: &Author{
			FirstName: "TestFirstName2",
			LastName:  "TestLastName2",
		},
		Title: "Harry potter2",
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
