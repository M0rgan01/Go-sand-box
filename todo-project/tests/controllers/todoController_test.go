package controllers

import (
	"github.com/morgan/Go-sand-box/todo-project/controllers"
	"github.com/morgan/Go-sand-box/todo-project/models"
	services "github.com/morgan/Go-sand-box/todo-project/services"
	"github.com/morgan/Go-sand-box/todo-project/tests"
	"github.com/morgan/Go-sand-box/todo-project/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCatalogController(t *testing.T) {

	testList := []tests.Test{
		{Title: "Get todo list", TestFunc: GetTodoListTest},
		{Title: "Get todo by ID", TestFunc: GetTodoByIdTest},
		{Title: "Add todo", TestFunc: AddTodoTest},
		{Title: "Update todo", TestFunc: UpdateTodoTest},
		{Title: "Delete todo", TestFunc: DeleteTodoTest},
	}

	tests.ExecuteIntegrationsTests(testList, t,
		func() {
			tests.DataBaseSetup()
		},
		func() {
			tests.DropDataBase()
		})
}

func GetTodoListTest(t *testing.T, services services.ServiceInstances) {

	// given
	router := controllers.SetupRoutes(services)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todoAPI/todo", nil)
	var todos []models.Todo
	tests.GetDbConnection().Find(&todos)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, tests.StructToString(todos), w.Body.String())
}

func GetTodoByIdTest(t *testing.T, services services.ServiceInstances) {

	// given
	router := controllers.SetupRoutes(services)
	w := httptest.NewRecorder()
	var todo models.Todo
	tests.GetDbConnection().First(&todo)

	req, _ := http.NewRequest("GET", "/todoAPI/todo/"+todo.ID.String(), nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, tests.StructToString(todo), w.Body.String())
}

func AddTodoTest(t *testing.T, services services.ServiceInstances) {

	// given
	router := controllers.SetupRoutes(services)
	w := httptest.NewRecorder()

	todoToAdd := models.Todo{
		ID:       utils.CreateUuid(),
		Title:    "test",
		Complete: true,
	}

	c := strings.NewReader(tests.StructToString(todoToAdd))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 201, w.Code)

	var count int64
	tests.GetDbConnection().Model(&models.Todo{}).Count(&count)
	assert.Equal(t, int64(3), count)

	var todo models.Todo
	tests.GetDbConnection().Find(&todo, todoToAdd.ID)
	assert.Equal(t, todo, todoToAdd)
}

func UpdateTodoTest(t *testing.T, services services.ServiceInstances) {

	// given
	router := controllers.SetupRoutes(services)
	w := httptest.NewRecorder()

	var todoToUpdate models.Todo
	tests.GetDbConnection().First(&todoToUpdate)
	todoToUpdate.Title = "TestUpdate"
	todoToUpdate.Complete = true

	c := strings.NewReader(tests.StructToString(todoToUpdate))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)

	var count int64
	tests.GetDbConnection().Model(&models.Todo{}).Count(&count)
	assert.Equal(t, int64(2), count)

	var todoUpdated models.Todo
	tests.GetDbConnection().First(&todoUpdated)
	assert.Equal(t, todoUpdated, todoToUpdate)
}

func DeleteTodoTest(t *testing.T, services services.ServiceInstances) {

	// given
	router := controllers.SetupRoutes(services)
	w := httptest.NewRecorder()

	var todoToDelete models.Todo
	tests.GetDbConnection().First(&todoToDelete)
	id := todoToDelete.ID.String()

	req, _ := http.NewRequest("DELETE", "/todoAPI/todo/"+id, nil)

	// when
	router.ServeHTTP(w, req)

	// then
	var count int64
	tests.GetDbConnection().Model(&models.Todo{}).Count(&count)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, int64(1), count)
}
