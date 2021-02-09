package main

import (
	"fmt"
	"time"
)

type Personne struct {
	firstName, lastName string
	age                 int
}

// Méthode de structure Personne
func (p Personne) getBirthYear() int {
	return time.Now().Year() - p.age
}

// les setters doivent fournir une référence
func (p *Personne) setAge(age int) {
	// pas besoin de ciblé la valeur avec *p.age comme dans un méthode classique, içi go fait le travaille
	p.age = age
}

func (p Personne) sayHello() {
	fmt.Println("Hello world")
}

func main() {

	// Déclaration 1
	var personne1 Personne
	personne1.age = 10
	personne1.firstName = "BiBi"
	personne1.lastName = "PoPo"

	// Déclaration 2
	personne2 := Personne{"BiBi2", "PoPo2", 11}

	// Déclaration 3
	personne3 := Personne{
		age:       12,
		lastName:  "PoPo3",
		firstName: "BiBi3",
	}

	fmt.Println(personne1)
	fmt.Println(personne2)
	fmt.Println(personne3)

	fmt.Println(personne3.getBirthYear())

	// Une forme de méthode statique sur une structure
	Personne{}.sayHello()
}
