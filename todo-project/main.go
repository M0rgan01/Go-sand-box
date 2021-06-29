package main

import (
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/fixtures"
	"github.com/morgan/Go-sand-box/todo-project/routes"
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

	r := routes.SetupRoutes()

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
