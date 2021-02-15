package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var todos []Todo

const port = 8080

func main() {

	// init public key for security
	fetchPublicKey()

	// init router
	r := mux.NewRouter()

	// init security handling
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
	r.HandleFunc("/todoAPI/todos", getTodos).Methods("GET").Name(adminRole)
	r.HandleFunc("/todoAPI/books/{id}", getTodo).Methods("GET").Name(adminRole)
	r.HandleFunc("/todoAPI/book", createTodo).Methods("POST").Name(adminRole)
	r.HandleFunc("/todoAPI/book/{id}", updateTodo).Methods("PUT").Name(adminRole)
	r.HandleFunc("/todoAPI/book/{id}", deleteTodo).Methods("DELETE").Name(adminRole)

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), CorsHandler(r)))
}
