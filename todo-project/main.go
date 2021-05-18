package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var todos []Todo

const port = 8080

func main() {
	setupConfiguration()
	// init public key for security
	fetchPublicKey()
	// create keycloak fixtures
	createKeycloakFixtures()

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
	r.Use(TokenAuthMiddleware())

	// Mock data
	todos = append(todos, Todo{
		Id:       createUuid(),
		Title:    "Harry potter",
		Complete: false,
	})

	todos = append(todos, Todo{
		Id:       createUuid(),
		Title:    "Harry potter 2",
		Complete: false,
	})

	// Route Handlers / Endpoints
	r.GET("/todoAPI/todos", getTodos)
	/*r.GET("/todoAPI/books/{id}", getTodo)
	r.POST("/todoAPI/book", createTodo)
	r.PUT("/todoAPI/book/{id}", updateTodo)
	r.DELETE("/todoAPI/book/{id}", deleteTodo)*/

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
