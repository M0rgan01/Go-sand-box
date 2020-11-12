package main

import "fmt"

func main() {

	// IF ESLE IF ELSE....

	variable := 10

	if variable > 10 {
		fmt.Println("Supérieur à 10")
	} else if variable < 10 {
		fmt.Println("Inférieur à 10")
	} else {
		fmt.Println("Valeur égal à 10")
	}

	// SWITCH CASE

	var result int

	switch variable {
	case 10:
		result = 10
	case 11:
		result = 11
	// default non obligatoire
	default:
		result = 0
	}

	fmt.Println("switch :", result)

	// il est possible de ne pas mettre de valeur sur le switch
	switch {
	case result == 10:
		result = 0
	case result > 10:
		result = 11
	case result >= 10:
		result = 11
	}

	fmt.Println("switch :", result)

}
