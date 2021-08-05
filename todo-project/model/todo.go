package model

import "github.com/google/uuid"

// Todo Model
type Todo struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Complete bool      `json:"complete"`
}
