package main

import "github.com/google/uuid"

// Book Model
type Book struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
}

// Author Model
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
