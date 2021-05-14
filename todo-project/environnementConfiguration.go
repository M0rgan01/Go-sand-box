package main

import (
	"flag"
	"log"
)

var envKey = devEnv

const (
	devEnv  = "development"
	prodEnv = "production"
)

type KeycloakAdminInfo struct {
	username    string
	password    string
	baseRealm   string
	uiClient    string
	uiClientUrl string
}

type KeycloakInfo struct {
	BaseURL        string
	ClientID       string
	ClientSecret   string
	Realm          string
	baseTheme      string
	loginTheme     string
	passwordPolicy string
}

type KeycloakSMTPInfo struct {
	auth            string
	host            string
	port            string
	user            string
	password        string
	from            string
	fromDisplayName string
	ssl             string
	starttls        string
}

var KeycloakDevInfo = KeycloakInfo{
	BaseURL:        "http://localhost",
	ClientID:       "TodoApi",
	Realm:          "TodoRealm",
	baseTheme:      "base",
	loginTheme:     "todo",
	passwordPolicy: "length(4)",
}

var keycloakDevAdminInfo = KeycloakAdminInfo{
	username:    "admin",
	password:    "admin",
	baseRealm:   "master",
	uiClient:    "TodoUi",
	uiClientUrl: "http://localhost:3000",
}

var KeycloakProdInfo = KeycloakInfo{
	BaseURL:        "http://localhost",
	ClientID:       "TodoApi",
	Realm:          "TodoRealm",
	baseTheme:      "base",
	loginTheme:     "todo",
	passwordPolicy: "length(8) and digits(1) and notUsername(undefined) and upperCase(1) and lowerCase(1) and specialChars(1)",
}

var keycloakProdAdminInfo = KeycloakAdminInfo{
	username:    "admin",
	password:    "admin",
	baseRealm:   "master",
	uiClient:    "TodoUi",
	uiClientUrl: "http://localhost:3000",
}

var keycloakDevSMTPServer = KeycloakSMTPInfo{
	user:            "test@gmail.com",
	port:            "1025",
	host:            "mailhog",
	ssl:             "false",
	auth:            "false",
	fromDisplayName: "Todo contact",
	from:            "test@gmail.com",
}

var keycloakInfo KeycloakInfo
var keycloakAdminInfo KeycloakAdminInfo
var keycloakSMTPInfo KeycloakSMTPInfo

func setKeycloakInfosByEnv() {
	log.Print("Actual env : " + envKey)
	switch envKey {
	case prodEnv:
		keycloakInfo = KeycloakProdInfo
		keycloakAdminInfo = keycloakProdAdminInfo
	case devEnv:
		keycloakSMTPInfo = keycloakDevSMTPServer
		keycloakInfo = KeycloakDevInfo
		keycloakAdminInfo = keycloakDevAdminInfo
	default:
		keycloakSMTPInfo = keycloakDevSMTPServer
		keycloakInfo = KeycloakDevInfo
		keycloakAdminInfo = keycloakDevAdminInfo
	}
}

func setFlags() {
	envKey = *flag.String("env", envKey, "environment")
	flag.Parse()
}
