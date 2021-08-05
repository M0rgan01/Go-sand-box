package configuration

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var EnvKey = DevEnv

const (
	DevEnv  = "development"
	ProdEnv = "production"
)

func SetupConfiguration() {
	EnvKey = *flag.String("env", EnvKey, "environment")
	flag.Parse()

	log.Print("Actual env : " + EnvKey)

	setKeycloakInfosByEnv()
	setGinMode()
}

func setGinMode() {
	if EnvKey == ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}
}
