package main

import "fmt"

func main() {

	/////// boucle trad

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/////// boucle foreach

	// avec un array
	myList := []string{"dog", "cat", "hedgehog"}

	// for {key}, {value} := range {list}
	for _, animal := range myList {
		fmt.Println("My animal is:", animal)
	}

	// avec une map
	myMap := map[string]string{
		"dog":      "woof",
		"cat":      "meow",
		"hedgehog": "sniff",
	}

	for animal, noise := range myMap {
		fmt.Println("The", animal, "went", noise)
	}

	/////// boucle while

	whileVar := 0

	for whileVar < 10 {
		fmt.Println(whileVar)
		whileVar++
	}

	/////// boucle infini

	for {
		// infini
		break
	}
}
