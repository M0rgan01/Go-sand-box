package main

import (
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/database"
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

	gorm, err := database.GetConnection()
	if err != nil {
		log.Fatalf("Failed to init gorm connection : %s", err.Error())
	}

	dbInstance, err := gorm.DB()
	if err != nil {
		log.Fatalf("Failed to get gorm DB instance : %s", err.Error())
	}

	err = database.Migrate(dbInstance)
	if err != nil {
		log.Fatalf("Failed to migrate database : %s", err.Error())
	}

	// TODO: à remplacer par des migrations
	// create database fixtures
	// fixtures.InitDatabaseSeed()

	r := routes.SetupRoutes()

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err = r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
