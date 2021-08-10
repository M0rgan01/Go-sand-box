package repositories

import "gorm.io/gorm"

type CatalogRepository struct {
	db *gorm.DB
}
