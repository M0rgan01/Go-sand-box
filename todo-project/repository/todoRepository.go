package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/model"
)

func GetTodoList() ([]model.Todo, error) {
	db, err := database.GetGormDBConnection()
	if err != nil {
		return nil, err
	}

	selDB, err := db.Query("SELECT * FROM todos ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	todo := model.Todo{}
	var todos []model.Todo
	for selDB.Next() {
		err := buildTodo(selDB, &todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoById(id uuid.UUID) (model.Todo, error) {
	db, err := database.GetGormDBConnection()
	if err != nil {
		return model.Todo{}, err
	}

	selDB, err := db.Query(`select * from todos where id = $1`, id)
	if err != nil {
		return model.Todo{}, err
	}

	todo := model.Todo{}
	for selDB.Next() {
		err := buildTodo(selDB, &todo)
		if err != nil {
			return model.Todo{}, err
		}
	}

	return todo, nil
}

func buildTodo(selDB *sql.Rows, todo *model.Todo) error {
	var id uuid.UUID
	var title string
	var complete bool
	var createdAt sql.NullString
	var updatedAt sql.NullString
	var deletedAt sql.NullString

	err := selDB.Scan(&id, &title, &complete, &createdAt, &updatedAt, &deletedAt)
	if err != nil {
		return err
	}

	*todo = model.Todo{ID: id, Title: title, Complete: complete}
	return nil
}

func GetTodosCount() (int, error) {
	db, err := database.GetGormDBConnection()
	if err != nil {
		return 0, err
	}

	rows, err := db.Query("SELECT count(*) FROM todos")
	if err != nil {
		return 0, err
	}

	var count int

	for rows.Next() {
		err = rows.Scan(&count)

		if err != nil {
			return 0, err
		}

	}

	return count, nil
}

func SaveTodo(todo model.Todo) (bool, error) {
	isTodoExist, err := GetTodoById(todo.ID)
	if err != nil {
		return false, err
	}

	if isTodoExist != (model.Todo{}) {
		err := updateTodo(todo)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		err := insertTodo(todo)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func insertTodo(todo model.Todo) error {

	db, err := database.GetGormDBConnection()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("INSERT INTO todos (id, title, complete) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = insForm.Exec(todo.ID, todo.Title, todo.Complete)
	if err != nil {
		return err
	}

	return nil
}

func updateTodo(todo model.Todo) error {
	db, err := database.GetGormDBConnection()
	if err != nil {
		return err
	}

	insForm, err := db.Prepare("UPDATE todos SET title=$1, complete=$2 WHERE id=$3")
	if err != nil {
		return err
	}

	_, err = insForm.Exec(todo.Title, todo.Complete, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id uuid.UUID) error {
	db, err := database.GetGormDBConnection()
	if err != nil {
		return err
	}

	delForm, err := db.Prepare("DELETE FROM todos WHERE id=$1")
	if err != nil {
		return err
	}

	_, err = delForm.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
