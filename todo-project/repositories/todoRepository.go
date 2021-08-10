package repositories

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (tr TodoRepository) GetTodoList() ([]models.Todo, error) {
	db, err := tr.db.DB()
	if err != nil {
		return nil, err
	}

	selDB, err := db.Query("SELECT * FROM todos ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	todo := models.Todo{}
	var todos []models.Todo
	for selDB.Next() {
		err := buildTodo(selDB, &todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (tr TodoRepository) GetTodoById(id uuid.UUID) (models.Todo, error) {
	db, err := tr.db.DB()
	if err != nil {
		return models.Todo{}, err
	}

	selDB, err := db.Query(`select * from todos where id = $1`, id)
	if err != nil {
		return models.Todo{}, err
	}

	todo := models.Todo{}
	for selDB.Next() {
		err := buildTodo(selDB, &todo)
		if err != nil {
			return models.Todo{}, err
		}
	}

	return todo, nil
}

func buildTodo(selDB *sql.Rows, todo *models.Todo) error {
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

	*todo = models.Todo{ID: id, Title: title, Complete: complete}
	return nil
}

func (tr TodoRepository) GetTodosCount() (int, error) {
	db, err := tr.db.DB()
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

func (tr TodoRepository) SaveTodo(todo models.Todo) (bool, error) {
	isTodoExist, err := tr.GetTodoById(todo.ID)
	if err != nil {
		return false, err
	}

	if isTodoExist != (models.Todo{}) {
		err := tr.updateTodo(todo)
		if err != nil {
			return false, err
		}
		return false, nil
	} else {
		err := tr.insertTodo(todo)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (tr TodoRepository) insertTodo(todo models.Todo) error {

	db, err := tr.db.DB()
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

func (tr TodoRepository) updateTodo(todo models.Todo) error {
	db, err := tr.db.DB()
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

func (tr TodoRepository) DeleteTodo(id uuid.UUID) error {
	db, err := tr.db.DB()
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
