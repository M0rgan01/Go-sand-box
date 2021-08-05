package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func GetGormInstance(dsn string) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	db = database
	log.Println("Gorm instance initializing")
	return db, nil
}

func GetGormDBConnection() (*sql.DB, error) {
	return db.DB()
}

// Without GORM
func OpenDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configuration.DBHost, configuration.DBPort, configuration.DBUser, configuration.DBPassword, configuration.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
