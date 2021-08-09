package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/model"
	services "github.com/morgan/Go-sand-box/todo-project/service"
	"net/http"
)

func SetupTodoRoutes(engine *gin.RouterGroup, todoService services.TodoService) {
	r := engine.Group("/todo")
	{
		r.GET("", func(context *gin.Context) {
			TodoController{context, todoService}.GetTodos()
		})
		r.POST("", func(context *gin.Context) {
			TodoController{context, todoService}.SaveTodo()
		})
		r.GET(":id", func(context *gin.Context) {
			TodoController{context, todoService}.GetTodoById()
		})
		r.DELETE(":id", func(context *gin.Context) {
			TodoController{context, todoService}.DeleteTodo()
		})
	}
}

type TodoController struct {
	Context     *gin.Context
	TodoService services.TodoService
}

func (tc TodoController) GetTodos() {

	todos, err := tc.TodoService.GetTodos()

	if err != nil {
		tc.Context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	tc.Context.JSON(http.StatusOK, todos)
}

func (tc TodoController) GetTodoById() {
	todoId, err := getTodoId(tc.Context)

	todo, err := tc.TodoService.GetTodoById(todoId)

	if err != nil {
		tc.Context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	tc.Context.JSON(http.StatusOK, todo)
}

func (tc TodoController) SaveTodo() {
	var todo model.Todo
	_ = json.NewDecoder(tc.Context.Request.Body).Decode(&todo)

	isCreated, err := tc.TodoService.SaveTodo(todo)
	if err != nil {
		tc.Context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if isCreated {
		tc.Context.Status(http.StatusCreated)
	} else {
		tc.Context.Status(http.StatusOK)
	}
}

func (tc TodoController) DeleteTodo() {
	todoId, err := getTodoId(tc.Context)

	err = tc.TodoService.DeleteTodo(todoId)

	if err != nil {
		tc.Context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	tc.Context.Status(http.StatusOK)
}

func getTodoId(c *gin.Context) (uuid.UUID, error) {
	id := c.Param("id")
	todoId, err := uuid.Parse(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return todoId, err
}
