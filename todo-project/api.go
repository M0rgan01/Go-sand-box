package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	writeJSON(todos, w, 201)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // get params

	var uuidFromString, err = uuid.Parse(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for _, item := range todos {
		if item.Id == uuidFromString {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Todo
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = createUuid() // Mock Id - not safe
	todos = append(todos, book)
	json.NewEncoder(w).Encode(book)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {

		var uuidFromString, err = uuid.Parse(params["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if item.Id == uuidFromString {
			todos = append(todos[:index], todos[index+1:]...)
			var book Todo
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = uuidFromString
			todos = append(todos, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var uuidFromString, err = uuid.Parse(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for index, item := range todos {

		if item.Id == uuidFromString {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

// writeJSON marshals data into JSON then outputs it on the response writer
// with appropriate status code.
func writeJSON(data interface{}, w http.ResponseWriter, statusCode int) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("err :", err, "data :", data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; encoding=\"utf-8\"")
	w.WriteHeader(statusCode)
}
