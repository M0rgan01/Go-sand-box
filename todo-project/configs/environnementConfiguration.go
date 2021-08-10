package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/logger"
	"os"
)

const (
	Environnement = "APP_ENV"
)

const (
	DevEnvironnement  = "development"
	ProdEnvironnement = "production"
)

var envList = []string{DevEnvironnement, ProdEnvironnement}

func GetEnvironnement() string {
	env, exists := os.LookupEnv(Environnement)
	if !exists {
		env = DevEnvironnement
	} else {
		ok := isValueExist(env)
		if ok {
			return env
		} else {
			logger.Errorf("Error when setting env, value %s not correct", env)
			return DevEnvironnement
		}
	}
	return env
}

func isValueExist(value string) bool {
	for _, v := range envList {
		if v == value {
			return true
		}
	}
	return false
}

func SetupConfiguration() {

	logger.Infof("Actual env: %s", GetEnvironnement())

	setKeycloakInfosByEnv()
	setGinMode()
}

func setGinMode() {
	if GetEnvironnement() == ProdEnvironnement {
		gin.SetMode(gin.ReleaseMode)
	}
}
