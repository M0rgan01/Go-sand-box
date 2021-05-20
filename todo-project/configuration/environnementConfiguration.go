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

type KeycloakAdminInfo struct {
	Username    string
	Password    string
	BaseRealm   string
	UiClient    string
	UiClientUrl string
}

type KeycloakInfo struct {
	BaseURL        string
	ClientID       string
	ClientSecret   string
	Realm          string
	BaseTheme      string
	LoginTheme     string
	PasswordPolicy string
}

type KeycloakSMTPInfo struct {
	Auth            string
	Host            string
	Port            string
	User            string
	Password        string
	From            string
	FromDisplayName string
	Ssl             string
	Starttls        string
}

var KeycloakDevInfo = KeycloakInfo{
	BaseURL:        "http://localhost",
	ClientID:       "TodoApi",
	Realm:          "TodoRealm",
	BaseTheme:      "base",
	LoginTheme:     "todo",
	PasswordPolicy: "length(4)",
}

var keycloakDevAdminInfo = KeycloakAdminInfo{
	Username:    "admin",
	Password:    "admin",
	BaseRealm:   "master",
	UiClient:    "TodoUi",
	UiClientUrl: "http://localhost:3000",
}

var KeycloakProdInfo = KeycloakInfo{
	BaseURL:        "http://localhost",
	ClientID:       "TodoApi",
	Realm:          "TodoRealm",
	BaseTheme:      "base",
	LoginTheme:     "todo",
	PasswordPolicy: "length(8) and digits(1) and notUsername(undefined) and upperCase(1) and lowerCase(1) and specialChars(1)",
}

var keycloakProdAdminInfo = KeycloakAdminInfo{
	Username:    "admin",
	Password:    "admin",
	BaseRealm:   "master",
	UiClient:    "TodoUi",
	UiClientUrl: "http://localhost:3000",
}

var keycloakDevSMTPServer = KeycloakSMTPInfo{
	User:            "test@gmail.com",
	Port:            "1025",
	Host:            "mailhog",
	Ssl:             "false",
	Auth:            "false",
	FromDisplayName: "Todo contact",
	From:            "test@gmail.com",
}

var BaseKeycloakInfo KeycloakInfo
var BaseKeycloakAdminInfo KeycloakAdminInfo
var BaseKeycloakSMTPInfo KeycloakSMTPInfo

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

func setKeycloakInfosByEnv() {
	switch EnvKey {
	case ProdEnv:
		BaseKeycloakInfo = KeycloakProdInfo
		BaseKeycloakAdminInfo = keycloakProdAdminInfo
	case DevEnv:
		BaseKeycloakSMTPInfo = keycloakDevSMTPServer
		BaseKeycloakInfo = KeycloakDevInfo
		BaseKeycloakAdminInfo = keycloakDevAdminInfo
	default:
		BaseKeycloakSMTPInfo = keycloakDevSMTPServer
		BaseKeycloakInfo = KeycloakDevInfo
		BaseKeycloakAdminInfo = keycloakDevAdminInfo
	}
}
