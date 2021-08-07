package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Test struct {
	Title    string
	TestFunc func(t *testing.T)
}

var testGormInstance *gorm.DB

const testMigrationDir = "../../migrations"

func DataBaseSetup() {
	db := getGormDb()

	err := database.Migrate(db, testMigrationDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully setup")
}

func GetDbConnection() *gorm.DB {
	if testGormInstance != nil {
		return testGormInstance
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		configuration.DBHost,
		configuration.DBUser,
		configuration.DBPassword,
		configuration.DBName,
		"5433",
		"sslmode=disable",
	)

	db, err := database.GetGormInstance(dsn)
	if err != nil {
		log.Fatal(err)
	}

	testGormInstance = db
	return testGormInstance
}

func getGormDb() *sql.DB {
	db, err := GetDbConnection().DB()
	if err != nil {
		log.Fatalf("failed to retreive connection to the database: %v", err)
	}
	return db
}

func DropDataBase() {
	db := getGormDb()
	_, migrationInstance, err := database.GetDBMigrationInstance(db, testMigrationDir)
	err = migrationInstance.Drop()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully drop")
}

func StructToString(a interface{}) string {
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func UuidFromString(s string) uuid.UUID {
	Uuid, _ := uuid.Parse(s)
	return Uuid
}

func ExecuteIntegrationsTests(testList []Test, tt *testing.T, beforeEach func(), afterEach func()) {

	//db := GetDbConnection()
	/*instantiatedDaos := daos.InitDAOSInstances(db)
	instantiatedServices := services.InitDAOSInstances(instantiatedDaos)*/

	for _, test := range testList {
		tt.Run(test.Title, func(t *testing.T) {
			if beforeEach != nil {
				beforeEach()
			}
			test.TestFunc(t)
			t.Cleanup(func() {
				if afterEach != nil {
					afterEach()
				}
			})
		})
	}
}
