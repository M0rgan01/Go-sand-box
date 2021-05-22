package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/controller"
	"github.com/morgan/Go-sand-box/todo-project/fixtures"
	"github.com/morgan/Go-sand-box/todo-project/security"
	"log"
	"strconv"
)

const port = 8080

func main() {
	configuration.SetupConfiguration()
	// init public key for security
	security.FetchPublicKey()
	// create keycloak fixtures
	security.CreateKeycloakFixtures()
	// create database fixtures
	fixtures.InitDatabaseSeed()

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

	// Route Handlers / Endpoints
	r.GET("/todoAPI/todos", controller.GetTodos)
	r.POST("/todoAPI/todo", controller.SaveTodo)
	r.GET("/todoAPI/todo/:id", controller.GetTodoById)
	r.DELETE("/todoAPI/todo/:id", controller.DeleteTodo)

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
