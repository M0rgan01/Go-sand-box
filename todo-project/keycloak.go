package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v8"
	"net/http"
)

const (

	// headerPrefix is the prefix the JWT has in the "Authorization" header's value.
	headerPrefix = "Bearer "
)

var (

	// ErrJWTExpired indicates the JWT has expired.
	ErrJWTExpired = errors.New("JWT has expired")
)

type KeycloakInfo struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	Realm        string
}

var keycloakInfo = KeycloakInfo{
	BaseURL:      "http://localhost",
	ClientID:     "TodoAPI",
	Realm:        "TodoRealm",
	ClientSecret: "41009360-201d-4747-aa96-0a5b71f17262"}

func HandleAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client := gocloak.NewClient(keycloakInfo.BaseURL)
		ctx := context.Background()

		// Strip the prefix from the header.
		headerValue := r.Header.Get("Authorization")

		_, claims, err := client.DecodeAccessToken(ctx, headerValue, keycloakInfo.Realm, "")

		if err != nil {
			panic("Failed to decode access token:" + err.Error())
		}

		fmt.Println(claims["dsdqs"])

		h.ServeHTTP(w, r)
	})
}
