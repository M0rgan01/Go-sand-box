package repository

import (
	"github.com/google/uuid"
	"github.com/morgan/Go-sand-box/todo-project/database"
	"github.com/morgan/Go-sand-box/todo-project/model"
)

func GetTodoList() ([]model.Todo, error) {
	db, err := database.OpenDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	selDB, err := db.Query("SELECT * FROM Todo ORDER BY id DESC")

	if err != nil {
		return nil, err
	}

	todo := model.Todo{}
	var todos []model.Todo
	for selDB.Next() {
		var id uuid.UUID
		var title string
		var complete bool
		err = selDB.Scan(&id, &title, &complete)
		if err != nil {
			return nil, err
		}
		todo.Id = id
		todo.Title = title
		todo.Complete = complete
		todos = append(todos, todo)
	}

	return todos, nil
}

func InsertTodo(todo model.Todo) (model.Todo, error) {

	db, err := database.OpenDB()
	defer db.Close()

	if err != nil {
		return model.Todo{}, err
	}

	if todo.Id.String() == "" {
		id, _ := uuid.NewUUID()
		todo.Id = id
	}

	sqlStatement := `INSERT INTO Todo (id, title, complete) VALUES ($1, $2, $3) RETURNING id`

	err = db.QueryRow(sqlStatement, todo.Id, todo.Title, todo.Complete).Scan(&todo.Id)
	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func GetTodosCount() (int, error) {
	db, err := database.OpenDB()
	defer db.Close()

	if err != nil {
		return 0, err
	}

	rows, err := db.Query("SELECT count(*) FROM Todo")

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

/*func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}


func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
*/
