package services

import (
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/models"
	"github.com/morgan/Go-sand-box/todo-project/repositories"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
}

func (ts TodoService) GetTodos() ([]models.Todo, error) {
	todos, err := ts.TodoRepository.GetTodoList()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (ts TodoService) GetTodoById(id uuid.UUID) (models.Todo, error) {
	todo, err := ts.TodoRepository.GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (ts TodoService) SaveTodo(todo models.Todo) (bool, error) {
	isCreated, err := ts.TodoRepository.SaveTodo(todo)
	if err != nil {
		return false, err
	}
	return isCreated, nil
}

func (ts TodoService) DeleteTodo(id uuid.UUID) error {
	err := ts.TodoRepository.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}
