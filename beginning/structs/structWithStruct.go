package main

import "fmt"

type People struct {
	firstName, lastName string
	age                 int
}

// Il est possible de déclarer une structure dans une autre
type Employee struct {
	id     int
	people People
}

func main() {

	people1 := People{age: 10, firstName: "Morgan", lastName: "Pichat"}
	employee1 := Employee{id: 1, people: people1}

	fmt.Println(employee1.people.lastName)
}
