package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Todo Model
type Todo struct {
	gorm.Model
	ID       uuid.UUID
	Title    string
	Complete bool
}
