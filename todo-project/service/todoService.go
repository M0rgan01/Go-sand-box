package services

import (
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/model"
	"github.com/morgan/Go-sand-box/todo-project/repository"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

func (ts TodoService) GetTodos() ([]model.Todo, error) {
	todos, err := ts.TodoRepository.GetTodoList()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (ts TodoService) GetTodoById(id uuid.UUID) (model.Todo, error) {
	todo, err := ts.TodoRepository.GetTodoById(id)
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (ts TodoService) SaveTodo(todo model.Todo) (bool, error) {
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
