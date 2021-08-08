package main

import (
	"github.com/morgan/Go-sand-box/todo-project/configs"
	"github.com/morgan/Go-sand-box/todo-project/controllers"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/repositories"
	"github.com/morgan/Go-sand-box/todo-project/security"
	services "github.com/morgan/Go-sand-box/todo-project/services"
	"log"
	"strconv"
)

const port = 8080

func main() {
	configs.SetupConfiguration()
	// create keycloak fixtures
	security.CreateKeycloakFixtures()
	// init public key for security
	security.FetchPublicKey()

	gorm, err := database.GetGormInstance(configs.GetDataBaseDSN())
	if err != nil {
		log.Fatalf("Failed to init gorm connection : %s", err.Error())
	}

	dbInstance, err := gorm.DB()
	if err != nil {
		log.Fatalf("Failed to get gorm DB instance : %s", err.Error())
	}

	err = database.Migrate(dbInstance, configs.MigrationsDirectory)
	if err != nil {
		log.Fatalf("Failed to migrate database : %s", err.Error())
	}

	instantiatedRepositories := repositories.InitRepositoriesInstances(gorm)
	instantiatedServices := services.InitDAOSInstances(instantiatedRepositories)

	r := controllers.SetupRoutes(instantiatedServices)

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err = r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
