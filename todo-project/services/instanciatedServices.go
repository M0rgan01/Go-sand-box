package services

import (
	"github.com/morgan/Go-sand-box/todo-project/repositories"
)

type ServiceInstances struct {
	CatalogService CatalogService
	TodoService    TodoService
}

func InitDAOSInstances(instances repositories.RepositoriesInstances) ServiceInstances {
	return ServiceInstances{
		CatalogService: CatalogService{instances.CatalogRepository},
		TodoService:    TodoService{instances.TodoRepository},
	}
}
