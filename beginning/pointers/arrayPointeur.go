package main

import "fmt"

func main() {
	var array = [2]int{10, 20}

	// L'adresse en mémoire du tableau est le même que celui du 1er élément
	println(&array)
	fmt.Println(&array[0])

	fmt.Println(&array[1])

	arrayPointer := &array

	println(arrayPointer)

	// syntaxe pour récupérer le x élément d'un pointeur sur un tableau
	fmt.Println((*arrayPointer)[0])
	// autre syntaxe... le comportement d'un pointeur sur un tableau est le même en réalité qu'un tableau normal
	fmt.Println(arrayPointer[0])

	//////// LES TABLEAUX D'ADRESSE

	adresse1 := new(int)
	adresse2 := new(int)

	arrayWithPointers := [2]*int{adresse1, adresse2}

	// le comportement est l'inverse d'un pointer sur un tableau normal
	fmt.Println(arrayWithPointers)
	fmt.Println(arrayWithPointers[0])
	fmt.Println(*arrayWithPointers[0])
}
