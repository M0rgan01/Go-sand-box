package security

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/morgan/Go-sand-box/todo-project/configs"
	"github.com/morgan/Go-sand-box/todo-project/logger"
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

func FetchPublicKey() {
	logger.Info("Fetching public key...")
	client := gocloak.NewClient(configs.BaseKeycloakInfo.BaseURL)
	ctx := context.Background()
	issuerInfo, err := client.GetIssuer(ctx, configs.BaseKeycloakInfo.Realm)

	if err != nil {

		logger.Infof("Error when fetching public key : %s", err)
		if retryFetchPublicKey < 4 {
			logger.Infof("--- Retry in %d sec ---", retryFetchPublicKey)
			time.Sleep(time.Duration(retryFetchPublicKey) * time.Second)
			retryFetchPublicKey++
			FetchPublicKey()
		} else {
			logger.Fatal("Please check availability of keycloak service")
		}
	} else {
		publicKey = *issuerInfo.PublicKey
		logger.Info("Fetching public key done !")
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
		logger.Debug("Error when parsing token : ", err)
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, true
	} else {
		logger.Debug("Invalid JWT Token")
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
