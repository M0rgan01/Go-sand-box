package services

import (
	"github.com/morgan/Go-sand-box/todo-project/repository"
)

type ServiceInstances struct {
	CatalogService CatalogService
	TodoService    TodoService
}

func InitDAOSInstances(instances repository.RepositoriesInstances) ServiceInstances {
	return ServiceInstances{
		CatalogService: CatalogService{instances.CatalogRepository},
		TodoService:    TodoService{instances.TodoRepository},
	}
}
