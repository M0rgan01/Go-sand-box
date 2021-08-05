package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/configuration"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/model"
	"github.com/morgan/Go-sand-box/todo-project/routes"
	"github.com/morgan/Go-sand-box/todo-project/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func structToString(a interface{}) string {
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func uuidFromString(s string) uuid.UUID {
	Uuid, _ := uuid.Parse(s)
	return Uuid
}

var todos = []model.Todo{
	{
		Id:       uuidFromString("47ad1ced-df3d-461b-b914-04213331cc36"),
		Title:    "Harry potter",
		Complete: false,
	},
	{
		Id:       uuidFromString("294d6fe3-9cc9-4fa3-9eda-9f70d84e83a6"),
		Title:    "Star wars",
		Complete: true,
	},
}

func TestCatalogController(t *testing.T) {

	migrationDir := "../../migrations"

	db, err := database.GetGormInstance(configuration.GetDataBaseDSN())

	if err != nil {
		t.FailNow()
	}

	dbConnection, err := db.DB()
	if err != nil {
		t.FailNow()
	}

	err = database.Migrate(dbConnection, migrationDir)
	if err != nil {
		t.FailNow()
	}

	t.Run("Get todo list", GetTodoListTest)
	t.Run("Get todo by ID", GetTodoByIdTest)
	t.Run("Add todo", AddTodoTest)
	t.Run("Update todo", UpdateTodoTest)
	t.Run("Delete todo", DeleteTodoTest)

	_, migrationInstance, err := database.GetDBMigrationInstance(dbConnection, migrationDir)

	err = migrationInstance.Drop()
	if err != nil {
		t.FailNow()
	}
}

func GetTodoListTest(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todoAPI/todo", nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(todos), w.Body.String())
}

func GetTodoByIdTest(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	todoToAssert := todos[0]

	req, _ := http.NewRequest("GET", "/todoAPI/todo/"+todoToAssert.Id.String(), nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(todoToAssert), w.Body.String())
}

func AddTodoTest(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	todoToAdd := model.Todo{
		Id:       utils.CreateUuid(),
		Title:    "test",
		Complete: true,
	}

	c := strings.NewReader(structToString(todoToAdd))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 3, len(todos))
	assert.Equal(t, todos[2], todoToAdd)
}

func UpdateTodoTest(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	todoToUpdate := todos[1]

	todoToUpdate.Title = "TestUpdate"
	todoToUpdate.Complete = true

	c := strings.NewReader(structToString(todoToUpdate))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(todos))
	assert.Equal(t, todos[1], todoToUpdate)
}

func DeleteTodoTest(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	id := todos[1].Id.String()

	req, _ := http.NewRequest("DELETE", "/todoAPI/todo/"+id, nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(todos))
	assert.NotEqual(t, todos[0].Id.String(), id)
}
