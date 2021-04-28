package main

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	userRole  = "user"
	adminRole = "admin"
)

var publicKey string

func fetchPublicKey() {
	log.Println("Fetching public key...")
	client := gocloak.NewClient(keycloakInfo.BaseURL)
	ctx := context.Background()
	issuerInfo, err := client.GetIssuer(ctx, keycloakInfo.Realm)

	if err != nil {
		log.Fatal("Error when fetching public key :", err)
	}

	publicKey = *issuerInfo.PublicKey
	log.Println("Fetching public key done !")
}

func HandleAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken := getAccessToken(r)
		claims, valid := extractClaims(accessToken)

		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else if !isClaimsContainRole(claims, userRole) {
			w.WriteHeader(http.StatusForbidden)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func isClaimsContainRole(claims jwt.MapClaims, role string) bool {
	for key, value := range claims {
		if key == "realm_access" {
			for _, realmAccessValue := range value.(map[string]interface{}) {
				for _, realmRole := range realmAccessValue.([]interface{}) {
					if realmRole == role {
						return true
					}
				}
			}
		}
	}
	return false
}

func getAccessToken(r *http.Request) string {
	// Strip the prefix from the header.
	headerValue := r.Header.Get("Authorization")
	accessToken := strings.Replace(headerValue, "Bearer ", "", 1)
	return accessToken
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {

	token, err := jwt.Parse(tokenStr, validateToken)

	if err != nil {
		log.Println("Error when parsing token : ", err)
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, true
	} else {
		log.Println("Invalid JWT Token")
		return nil, false
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {

	base64Data := []byte("-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----")

	key, err := jwt.ParseRSAPublicKeyFromPEM(base64Data)

	if err != nil {
		return nil, nil
	}
	return key, nil
}
