package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/controller"
	"github.com/morgan/Go-sand-box/todo-project/model"
)

func SetupWebsocketRoutes(engine *gin.RouterGroup) {
	hub := model.NewHub()
	go hub.Run()
	r := engine.Group("/ws")
	{
		r.GET("/chat", func(context *gin.Context) {
			controller.ServeWs(hub, context)
		})
	}
}
