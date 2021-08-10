package configs

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
