package repository

import "gorm.io/gorm"

type CatalogRepository struct {
	db *gorm.DB
}
