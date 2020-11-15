package main

import "fmt"

func main() {

	////////////////////// Array

	var array [2]int8
	array[0] = 10
	array[1] = 20

	// autre déclaration
	array1 := [2]int8{5, 10}

	fmt.Println(array, array1)

	////////////////////// Slice

	array2 := [5]int8{5, 10, 15, 20, 25}

	// prend l'index 0 (inclus) à l'index 3 (exclu) du tableau
	slice1 := array2[0:3]
	// syntaxe équivalente sur un index 0
	slice2 := array2[:3]

	fmt.Println(slice1, slice2)

	// les slices sont redimensionnable
	slice2 = append(slice2, 123)

	// il est possible de créer un slice sans passer par un array
	slice := make([]int, 0)
	// autre déclaration
	slice = []int{}

	fmt.Println(slice)

	////////////////////// Map

	map1 := map[string]int{"Un": 1, "Deux": 2, "Trois": 3}
	// ou
	map1 = map[string]int{
		"Un":    1,
		"Deux":  2,
		"Trois": 3,
	}

	// autre déclaration
	map1 = make(map[string]int)
	map1["Un"] = 1
	map1["Deux"] = 2
	map1["Trois"] = 3

	fmt.Println(map1, map1["Un"])

	// suppression
	delete(map1, "Deux")

	fmt.Println(map1)
}
