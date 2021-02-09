package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var todos []Todo

func main() {
	// init router
	r := mux.NewRouter()

	r.Use(HandleAuth)

	// Mock data
	todos = append(todos, Todo{
		Id:       createUuid(),
		Title:    "Harry potter",
		Complete: false,
	})

	todos = append(todos, Todo{
		Id:       createUuid(),
		Title:    "Harry potter 2",
		Complete: false,
	})

	// Route Handlers / Endpoints
	r.HandleFunc("/todoAPI/todos", getTodos).Methods("GET")
	r.HandleFunc("/todoAPI/books/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todoAPI/book", createTodo).Methods("POST")
	r.HandleFunc("/todoAPI/book/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todoAPI/book/{id}", deleteTodo).Methods("DELETE")

	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
