package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/security"
)

func SetupRoutes() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// CORS handling
	r.Use(security.CorsHandler())

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// init security handling
	r.Use(security.TokenAuthMiddleware())

	apiGroup := r.Group("/todoAPI")

	SetupTodoRoutes(apiGroup)

	return r
}
