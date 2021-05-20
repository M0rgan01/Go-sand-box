package fixtures

import (
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/model"
	"github.com/morgan/Go-sand-box/todo-project/repository"
	"log"
)

func InitDatabaseSeed() {
	databaseInitialization()
	initTodoSeed()
}

func databaseInitialization() {

	_, tableNotExist := repository.GetTodoList()

	if tableNotExist != nil {

		db, err := database.OpenDB()
		defer db.Close()
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("CREATE TABLE Todo ( id uuid, title varchar(255), complete bool )")

		if err != nil {
			log.Printf("Error when execute SQL command : " + err.Error())
			panic(err)
		}

		log.Println("Database initialization successful !")
	} else {
		log.Println("Database already initialized")
	}

}

func initTodoSeed() {

	todoCount, err := repository.GetTodosCount()

	if err != nil {
		panic(err)
	} else if todoCount == 0 {

		_, err := repository.InsertTodo(model.Todo{
			Title:    "Harry potter",
			Complete: false,
		})

		if err != nil {
			panic(err)
		}

		_, err = repository.InsertTodo(model.Todo{
			Title:    "Star wars",
			Complete: true,
		})

		if err != nil {
			panic(err)
		}

		log.Println("Todo seed initialization successful !")
	} else {
		log.Println("Todo seed already initialized")
	}
}
