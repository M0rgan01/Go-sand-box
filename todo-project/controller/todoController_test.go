package controller

import (
	"encoding/json"
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

func TestCatalogController(t *testing.T) {
	t.Run("Get todo list", GetTodoList)
	t.Run("Get todo by ID", GetTodoById)
	t.Run("Add todo", AddTodo)
	t.Run("Update todo", UpdateTodo)
	t.Run("Delete todo", DeleteTodo)
}

func GetTodoList(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todoAPI/todo", nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(repository.Catalogs), w.Body.String())
}

func GetTodoById(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	catalogToAssert := repository.Catalogs[0]

	req, _ := http.NewRequest("GET", "/todoAPI/todo/"+catalogToAssert.Id.String(), nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, structToString(catalogToAssert), w.Body.String())
}

func AddTodo(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	catalogToAdd := model.Catalog{
		Id:           utils.CreateUuid(),
		SupplierName: "test",
		Enabled:      true,
		CreatedBy:    "admin",
		CreatedAt:    "07/08/2012",
	}

	c := strings.NewReader(structToString(catalogToAdd))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 3, len(repository.Catalogs))
	assert.Equal(t, repository.Catalogs[2], catalogToAdd)
}

func UpdateTodo(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	catalogToUpdate := repository.Catalogs[1]

	catalogToUpdate.SupplierName = "TestUpdate"
	catalogToUpdate.Enabled = true
	catalogToUpdate.CreatedBy = "TestAdmin"

	c := strings.NewReader(structToString(catalogToUpdate))
	req, _ := http.NewRequest("POST", "/todoAPI/todo", c)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 2, len(repository.Catalogs))
	assert.Equal(t, repository.Catalogs[1], catalogToUpdate)
}

func DeleteTodo(t *testing.T) {

	// given
	router := routes.SetupRoutes()
	w := httptest.NewRecorder()

	id := repository.Catalogs[1].Id.String()

	req, _ := http.NewRequest("DELETE", "/todoAPI/todo/"+id, nil)

	// when
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(repository.Catalogs))
	assert.NotEqual(t, repository.Catalogs[0].Id.String(), id)
}
