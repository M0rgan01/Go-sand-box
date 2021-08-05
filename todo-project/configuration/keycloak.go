package configuration

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
