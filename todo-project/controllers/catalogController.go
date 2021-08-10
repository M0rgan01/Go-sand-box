package controllers

import (
	"github.com/gin-gonic/gin"
	services "github.com/morgan/Go-sand-box/todo-project/services"
)

type CatalogController struct {
	Context        *gin.Context
	CatalogService services.CatalogService
}

/*func GetCatalogs(c *gin.Context) {

	todos, err := repository.GetTodoList()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}*/
