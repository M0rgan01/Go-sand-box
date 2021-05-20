package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/controller"
	"github.com/morgan/Go-sand-box/todo-project/fixtures"
	"github.com/morgan/Go-sand-box/todo-project/keycloak"
	"log"
	"strconv"
)

const port = 8080

func main() {
	configuration.SetupConfiguration()
	// init public key for security
	keycloak.FetchPublicKey()
	// create keycloak fixtures
	keycloak.CreateKeycloakFixtures()
	// create database fixtures
	fixtures.InitDatabaseSeed()

	// Creates a router without any middleware by default
	r := gin.New()

	// CORS handling
	r.Use(CorsHandler())

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// init security handling
	r.Use(keycloak.TokenAuthMiddleware())

	// Route Handlers / Endpoints
	r.GET("/todoAPI/todos", controller.GetTodos)
	r.POST("/todoAPI/todo", controller.CreateTodo)
	/*r.GET("/todoAPI/books/{id}", getTodo)

	r.PUT("/todoAPI/book/{id}", updateTodo)
	r.DELETE("/todoAPI/book/{id}", deleteTodo)*/

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
