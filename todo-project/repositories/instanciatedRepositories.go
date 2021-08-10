package repositories

import "gorm.io/gorm"

type RepositoriesInstances struct {
	CatalogRepository CatalogRepository
	TodoRepository    TodoRepository
}

func InitRepositoriesInstances(db *gorm.DB) RepositoriesInstances {
	return RepositoriesInstances{
		CatalogRepository: CatalogRepository{db},
		TodoRepository:    TodoRepository{db},
	}
}
