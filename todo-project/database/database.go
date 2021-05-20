package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5433
	dbUser     = "admin"
	dbPassword = "password"
	dbname     = "app_database"
)

func OpenDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	return db, nil
}
