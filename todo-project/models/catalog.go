package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model
	ID       uuid.UUID
	Name     string
	Products []Product
}
