package main

import "sync"

type Coaster struct {
	Name   string
	ID     string
	Height int
}

type coasterHandler struct {
	sync.Mutex
	store map[string]Coaster
}
