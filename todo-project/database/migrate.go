package database

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"io/ioutil"
	"strconv"
	"strings"
)

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.New("failed to init postgres driver : " + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:./migrations",
		"postgres", driver)
	if err != nil {
		return errors.New("failed to init golang-migrate : " + err.Error())
	}

	lfv, err := getLastFileVersion()
	if err != nil {
		return errors.New("failed to get last file migration version : " + err.Error())
	}

	lastFileVersion, err := strconv.ParseUint(lfv, 10, 64)
	if err != nil {
		return errors.New("Error when converting last file version to int : " + err.Error())
	}

	dataBaseVersion, _, err := m.Version()
	if err != nil {
		return errors.New("failed to get database migration version : " + err.Error())
	}

	if uint(lastFileVersion) > dataBaseVersion {
		err = m.Up()

		if err != nil {
			return errors.New("failed to migrate database : " + err.Error())
		}
	}

	return nil
}

func getLastFileVersion() (string, error) {
	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		return "", err
	}

	lastFile := files[len(files)-1]

	return strings.Split(lastFile.Name(), "_")[0], nil
}
