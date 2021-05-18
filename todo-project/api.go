package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

/*func getTodo(c *gin.Context) {
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

func createTodo(c *gin.Context) {
	w.Header().Set("Content-Type", "application/json")
	var book Todo
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = createUuid() // Mock Id - not safe
	todos = append(todos, book)
	json.NewEncoder(w).Encode(book)
}

func updateTodo(c *gin.Context) {
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

func deleteTodo(c *gin.Context) {
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
}*/
