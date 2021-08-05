package configuration

import "fmt"

const (
	DBHost                 = "localhost"
	DBPort                 = 5433
	DBUser                 = "admin"
	DBPassword             = "password"
	DBName                 = "app_database"
	MigrationsDirectory    = "./migrations"
	MigrationFileExtension = ".sql"
)

func GetDataBaseDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d %s",
		DBHost,
		DBUser,
		DBPassword,
		DBName,
		DBPort,
		"sslmode=disable",
	)
}
