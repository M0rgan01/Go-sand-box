package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string
}

type Product struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string
}
