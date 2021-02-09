package main

import (
	"fmt"
	"github.com/google/uuid"
)

func createUuid() uuid.UUID {
	var random, err = uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	return random
}
