package keycloak

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"log"
	"strings"
)

func CreateKeycloakFixtures() {
	updateAdminUser()
	createTodoRealm()
	createDevBaseUser()
}

func createDevBaseUser() {
	if configuration.EnvKey == configuration.DevEnv {
		client, ctx, accessToken := getClient()
		count, err := client.GetUserCount(ctx, accessToken, configuration.BaseKeycloakInfo.Realm, gocloak.GetUsersParams{})
		if err != nil {
			panic("Something wrong when fetching user count : " + err.Error())
		} else if count == 0 {
			log.Println("Create dev user...")

			credential := gocloak.CredentialRepresentation{
				Temporary: gocloak.BoolP(false),
				Type:      gocloak.StringP("password"),
				SecretData: gocloak.StringP("{" +
					"\"value\": \"fa8akRKCigEtANBY+RM82A==\", " +
					"\"salt\": \"ExrLjxuU9ag+mmAmYLI8KA==\"" +
					"}"),
				CredentialData: gocloak.StringP("{\"hashIterations\": 27500, \"algorithm\": \"pbkdf2-sha256\"}"),
			}

			user := gocloak.User{
				Username:    gocloak.StringP("Test@email.com"),
				FirstName:   gocloak.StringP("Dev"),
				LastName:    gocloak.StringP("Test"),
				Email:       gocloak.StringP("Test@email.com"),
				Enabled:     gocloak.BoolP(true),
				Credentials: &[]gocloak.CredentialRepresentation{credential},
			}
			_, err = client.CreateUser(ctx, accessToken, configuration.BaseKeycloakInfo.Realm, user)

			if err != nil {
				panic("Something wrong when create user : " + err.Error())
			} else {
				log.Println("Create dev user done !")
			}
		}
	}
}

func getClient() (gocloak.GoCloak, context.Context, string) {
	client := gocloak.NewClient(configuration.BaseKeycloakInfo.BaseURL)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, configuration.BaseKeycloakAdminInfo.Username, configuration.BaseKeycloakAdminInfo.Password, configuration.BaseKeycloakAdminInfo.BaseRealm)
	if err != nil {
		panic("Something wrong with the credentials or url : " + err.Error())
	}
	return client, ctx, token.AccessToken
}

// set base ADMIN info
func updateAdminUser() {
	client, ctx, accessToken := getClient()
	adminUserParam := gocloak.GetUsersParams{Username: gocloak.StringP("admin")}
	users, err := client.GetUsers(ctx, accessToken, "master", adminUserParam)

	if err != nil {
		panic("Something wrong when fetching Admin user : " + err.Error())
	}

	adminUser := users[0]

	if adminUser.Email == nil {
		log.Println("Update keycloak ADMIN user...")

		user := gocloak.User{
			ID:        adminUser.ID,
			Username:  adminUser.Username,
			FirstName: gocloak.StringP("keycloakAdmin"),
			LastName:  gocloak.StringP("keycloakAdmin"),
			Email:     gocloak.StringP(configuration.BaseKeycloakSMTPInfo.From),
		}

		err = client.UpdateUser(ctx, accessToken, "master", user)

		if err != nil {
			panic("Something wrong when update Admin user : " + err.Error())
		} else {
			log.Println("Update keycloak ADMIN user done !")
		}
	}
}

func createTodoRealm() {
	client, ctx, accessToken := getClient()
	_, err := client.GetRealm(ctx, accessToken, configuration.BaseKeycloakInfo.Realm)

	if err != nil {
		if !strings.Contains(err.Error(), "404") {
			panic("Something wrong when fetching realm : " + err.Error())
		} else {
			log.Println("Create Todo realm...")

			roles := buildRoles([]string{userRole})
			uiClient := buildUiClient()
			smtServer := buildSMTPServer()

			todoRealm := gocloak.RealmRepresentation{
				ID:                  gocloak.StringP(configuration.BaseKeycloakInfo.Realm),
				Realm:               gocloak.StringP(configuration.BaseKeycloakInfo.Realm),
				Enabled:             gocloak.BoolP(true),
				RegistrationAllowed: gocloak.BoolP(true),
				DefaultLocale:       gocloak.StringP("en"),
				EmailTheme:          gocloak.StringP(configuration.BaseKeycloakInfo.BaseTheme),
				AdminTheme:          gocloak.StringP(configuration.BaseKeycloakInfo.BaseTheme),
				LoginTheme:          gocloak.StringP(configuration.BaseKeycloakInfo.LoginTheme),
				Roles:               &roles,
				DefaultRoles:        &[]string{userRole},
				Clients:             &[]gocloak.Client{uiClient},
				PasswordPolicy:      gocloak.StringP(configuration.BaseKeycloakInfo.PasswordPolicy),
				SMTPServer:          &smtServer,
			}

			_, err = client.CreateRealm(ctx, accessToken, todoRealm)
			if err != nil {
				panic("Error when create realm : " + err.Error())
			} else {
				log.Println("Create Todo realm done !")
			}
		}
	}
}

func buildRoles(roles []string) gocloak.RolesRepresentation {
	var keycloakRoles []gocloak.Role
	for _, roleAsString := range roles {
		role := gocloak.Role{
			Name: gocloak.StringP(roleAsString),
		}
		keycloakRoles = append(keycloakRoles, role)
	}
	return gocloak.RolesRepresentation{Realm: &keycloakRoles}
}

func buildUiClient() gocloak.Client {
	return gocloak.Client{
		Name:         gocloak.StringP(configuration.BaseKeycloakAdminInfo.UiClient),
		ID:           gocloak.StringP(configuration.BaseKeycloakAdminInfo.UiClient),
		BaseURL:      gocloak.StringP(configuration.BaseKeycloakAdminInfo.UiClientUrl),
		Enabled:      gocloak.BoolP(true),
		PublicClient: gocloak.BoolP(true),
		RedirectURIs: &[]string{configuration.BaseKeycloakAdminInfo.UiClientUrl + "/*"},
		WebOrigins:   &[]string{configuration.BaseKeycloakAdminInfo.UiClientUrl},
	}
}

func buildSMTPServer() map[string]string {
	return map[string]string{
		"auth":            configuration.BaseKeycloakSMTPInfo.Auth,
		"host":            configuration.BaseKeycloakSMTPInfo.Host,
		"port":            configuration.BaseKeycloakSMTPInfo.Port,
		"user":            configuration.BaseKeycloakSMTPInfo.User,
		"password":        configuration.BaseKeycloakSMTPInfo.Password,
		"from":            configuration.BaseKeycloakSMTPInfo.From,
		"fromDisplayName": configuration.BaseKeycloakSMTPInfo.FromDisplayName,
		"ssl":             configuration.BaseKeycloakSMTPInfo.Ssl,
		"starttls":        configuration.BaseKeycloakSMTPInfo.Starttls,
	}
}
