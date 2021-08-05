package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uuid.UUID
	Number    string
	Amount    int
	CatalogID uuid.UUID
}
