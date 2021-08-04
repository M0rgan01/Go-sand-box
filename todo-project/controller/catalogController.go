package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/repository"
	"net/http"
)

func GetCatalogs(c *gin.Context) {

	todos, err := repository.GetTodoList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}
