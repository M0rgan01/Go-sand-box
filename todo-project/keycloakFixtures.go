package main

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	log "github.com/sirupsen/logrus"
	"strings"
)

func createKeycloakFixtures() {
	updateAdminUser()
	createTodoRealm()
	createDevBaseUser()
}

func createDevBaseUser() {
	if envKey == devEnv {
		client, ctx, accessToken := getClient()
		count, err := client.GetUserCount(ctx, accessToken, keycloakInfo.Realm, gocloak.GetUsersParams{})
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
			_, err = client.CreateUser(ctx, accessToken, keycloakInfo.Realm, user)

			if err != nil {
				panic("Something wrong when create user : " + err.Error())
			} else {
				log.Println("Create dev user done !")
			}
		}
	}
}

func getClient() (gocloak.GoCloak, context.Context, string) {
	client := gocloak.NewClient(keycloakInfo.BaseURL)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, keycloakAdminInfo.username, keycloakAdminInfo.password, keycloakAdminInfo.baseRealm)
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
			Email:     gocloak.StringP(keycloakSMTPInfo.from),
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
	_, err := client.GetRealm(ctx, accessToken, keycloakInfo.Realm)

	if err != nil {
		if !strings.Contains(err.Error(), "404") {
			panic("Something wrong when fetching realm : " + err.Error())
		} else {
			log.Println("Create Todo realm...")

			roles := buildRoles([]string{userRole})
			uiClient := buildUiClient()
			smtServer := buildSMTPServer()

			todoRealm := gocloak.RealmRepresentation{
				ID:                  gocloak.StringP(keycloakInfo.Realm),
				Realm:               gocloak.StringP(keycloakInfo.Realm),
				Enabled:             gocloak.BoolP(true),
				RegistrationAllowed: gocloak.BoolP(true),
				DefaultLocale:       gocloak.StringP("en"),
				EmailTheme:          gocloak.StringP(keycloakInfo.baseTheme),
				AdminTheme:          gocloak.StringP(keycloakInfo.baseTheme),
				LoginTheme:          gocloak.StringP(keycloakInfo.loginTheme),
				Roles:               &roles,
				DefaultRoles:        &[]string{userRole},
				Clients:             &[]gocloak.Client{uiClient},
				PasswordPolicy:      gocloak.StringP(keycloakInfo.passwordPolicy),
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
		Name:         gocloak.StringP(keycloakAdminInfo.uiClient),
		ID:           gocloak.StringP(keycloakAdminInfo.uiClient),
		BaseURL:      gocloak.StringP(keycloakAdminInfo.uiClientUrl),
		Enabled:      gocloak.BoolP(true),
		PublicClient: gocloak.BoolP(true),
		RedirectURIs: &[]string{keycloakAdminInfo.uiClientUrl + "/*"},
		WebOrigins:   &[]string{keycloakAdminInfo.uiClientUrl},
	}
}

func buildSMTPServer() map[string]string {
	return map[string]string{
		"auth":            keycloakSMTPInfo.auth,
		"host":            keycloakSMTPInfo.host,
		"port":            keycloakSMTPInfo.port,
		"user":            keycloakSMTPInfo.user,
		"password":        keycloakSMTPInfo.password,
		"from":            keycloakSMTPInfo.from,
		"fromDisplayName": keycloakSMTPInfo.fromDisplayName,
		"ssl":             keycloakSMTPInfo.ssl,
		"starttls":        keycloakSMTPInfo.starttls,
	}
}
