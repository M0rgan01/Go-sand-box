package database

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

const migrationsDirectory = "./migrations"
const migrationFileExtension = ".sql"

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.New("failed to init postgres driver : " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:"+migrationsDirectory,
		"postgres", driver)
	if err != nil {
		return errors.New("failed to init golang-migrate : " + err.Error())
	}

	lfv, isDirEmpty, err := getLastMigrationFileVersion(migrationsDirectory)
	if err != nil {
		return errors.New("failed to get last file migration version : " + err.Error())
	}

	if isDirEmpty {
		return nil
	}

	lastFileVersion, err := strconv.ParseUint(lfv, 10, 64)
	if err != nil {
		return errors.New("Error when converting last file version to int : " + err.Error())
	}

	actualDataBaseMigrationVersion, _, err := driver.Version()
	if err != nil {
		return errors.New("failed to get database migration version : " + err.Error())
	}

	if actualDataBaseMigrationVersion != -1 {
		actualMigrationVersion, _, err := m.Version()
		if err != nil {
			return errors.New("failed to get database migration version : " + err.Error())
		}
		if uint(lastFileVersion) > actualMigrationVersion {
			err = proceedMigrations(m)
		}
	} else {
		err = proceedMigrations(m)
	}

	if err != nil {
		return err
	}

	return nil
}

func proceedMigrations(m *migrate.Migrate) error {
	log.Println("Migrations detected, proceed database update...")
	err := m.Up()

	if err != nil {
		return errors.New("failed to migrate database : " + err.Error())
	}
	log.Println("Migrations succeed")
	return nil
}

func getLastMigrationFileVersion(dir string) (string, bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", false, err
	}

	var migrationFile []fs.FileInfo

	for _, file := range files {
		fileExtension := filepath.Ext(file.Name())
		if fileExtension == migrationFileExtension {
			migrationFile = append(migrationFile, file)
		}
	}

	if len(migrationFile) == 0 {
		return "", true, nil
	}

	lastFile := migrationFile[len(migrationFile)-1]
	lastFileVersion := strings.Split(lastFile.Name(), "_")[0]

	return lastFileVersion, false, nil
}
