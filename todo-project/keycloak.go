package main

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	userRole  = "user"
	adminRole = "admin"
)

var publicKey string
var retryFetchPublicKey = 1

func fetchPublicKey() {
	log.Println("Fetching public key...")
	client := gocloak.NewClient(keycloakInfo.BaseURL)
	ctx := context.Background()
	issuerInfo, err := client.GetIssuer(ctx, keycloakInfo.Realm)

	if err != nil {

		log.Printf("Error when fetching public key : %s", err)
		if retryFetchPublicKey < 4 {
			log.Printf("--- Retry in %d sec ---", retryFetchPublicKey)
			time.Sleep(time.Duration(retryFetchPublicKey) * time.Second)
			retryFetchPublicKey++
			fetchPublicKey()
		} else {
			log.Fatal("Please check availability of keycloak service")
		}
	} else {
		publicKey = *issuerInfo.PublicKey
		log.Println("Fetching public key done !")
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken := getAccessToken(c.Request)
		claims, valid := extractClaims(accessToken)

		if !valid {
			respondWithError(c, http.StatusUnauthorized, "Token not valid")
		} else if !isClaimsContainRole(claims, userRole) {
			respondWithError(c, http.StatusForbidden, "Forbidden")
		}

		c.Next()
	}
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
