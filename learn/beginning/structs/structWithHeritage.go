package main

import "fmt"

type Animal struct {
	firstName string
	age       int
}

// Il est possible de déclarer une structure dans une autre
type Chien struct {
	id int
	// forme d'héritage
	Animal
}

func main() {

	animal := Animal{age: 12, firstName: "MonChien"}

	labrador := Chien{id: 50, Animal: animal}

	// les props sont directement accessible
	fmt.Println(labrador.age)
}
