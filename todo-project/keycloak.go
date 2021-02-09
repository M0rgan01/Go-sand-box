package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
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
	ClientSecret: "41009360-201d-4747-aa96-0a5b71f17262",
}

func HandleAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Strip the prefix from the header.
		headerValue := r.Header.Get("Authorization")
		accessToken := strings.Replace(headerValue, "Bearer ", "", 1)

		claims, valid := extractClaims(accessToken)

		fmt.Println(claims)
		fmt.Println(valid)

		h.ServeHTTP(w, r)
	})
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	fmt.Println(tokenStr)

	token, err := jwt.Parse(tokenStr, validToken)

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Println("Invalid JWT Token")
		return nil, false
	}
}

func validToken(token *jwt.Token) (interface{}, error) {

	pubPEM := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkBQQL+BzdMyMTqPUjV8K\nAEKKHaR09o2tFbPZtskhgXbOfh0DyHXjD3oORCC6m9btqEILItPSO6+VdjJ/Kspp\nCTYs96WIj66ywoZxntvqWvSj+0q14TYYQ7AuxEhmRAqCJIn8+Zq+VNkA2WKwYwWV\nq5qsDKWB25hHuAsYHYxVIy3lQsPLnmIvvWFRDskV4LNCOZDlV5TfzaC59yalFM/E\n1mv9eYJ4Ji1N0fDgz73t9m5YGnLLRlhDeI/3cxx2HU3r+GasjQPc4U/lclgIoRAa\nPVvfWRAjvf+nLhqUlM77wnVWCD+RhodYl+FgYHYf2lZ/tlG/U7+/dmqkRc721eM5\nuwIDAQAB\n-----END PUBLIC KEY-----"

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubPEM))

	if err != nil {
		return nil, nil
	}
	return key, nil
}
