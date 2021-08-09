package main

import (
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/controller"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/repository"
	"github.com/morgan/Go-sand-box/todo-project/security"
	services "github.com/morgan/Go-sand-box/todo-project/service"
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

	gorm, err := database.GetGormInstance(configuration.GetDataBaseDSN())
	if err != nil {
		log.Fatalf("Failed to init gorm connection : %s", err.Error())
	}

	dbInstance, err := gorm.DB()
	if err != nil {
		log.Fatalf("Failed to get gorm DB instance : %s", err.Error())
	}

	err = database.Migrate(dbInstance, configuration.MigrationsDirectory)
	if err != nil {
		log.Fatalf("Failed to migrate database : %s", err.Error())
	}

	instantiatedRepositories := repository.InitRepositoriesInstances(gorm)
	instantiatedServices := services.InitDAOSInstances(instantiatedRepositories)

	r := controller.SetupRoutes(instantiatedServices)

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err = r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
