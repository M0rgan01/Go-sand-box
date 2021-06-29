package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/controller"
)

func SetupTodoRoutes(engine *gin.RouterGroup) {
	r := engine.Group("/todo")
	{
		r.GET("", controller.GetTodos)
		r.POST("", controller.SaveTodo)
		r.GET(":id", controller.GetTodoById)
		r.DELETE(":id", controller.DeleteTodo)
	}
}
