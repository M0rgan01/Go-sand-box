package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/model"
	"github.com/morgan/Go-sand-box/todo-project/repository"
	"net/http"
)

func GetTodos(c *gin.Context) {

	todos, err := repository.GetTodoList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	todoId, err := getTodoId(c)

	todo, err := repository.GetTodoById(todoId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)
}

func SaveTodo(c *gin.Context) {
	var todo model.Todo
	_ = json.NewDecoder(c.Request.Body).Decode(&todo)

	isCreated, err := repository.SaveTodo(todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if isCreated {
		c.Status(http.StatusCreated)
	} else {
		c.Status(http.StatusOK)
	}
}

func DeleteTodo(c *gin.Context) {
	todoId, err := getTodoId(c)

	err = repository.DeleteTodo(todoId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func getTodoId(c *gin.Context) (uuid.UUID, error) {
	id := c.Param("id")
	todoId, err := uuid.Parse(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return todoId, err
}
