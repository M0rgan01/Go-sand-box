package main

import (
	"net/http"
)

func main() {

	coasterHandlers := newCoasterHandler()

	http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/coasters", coasterHandlers.getCoaster)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
