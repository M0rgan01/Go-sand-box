package main

import "fmt"

////////// FONCTION AVEC SPREAD OPERATOR
func computeSomme(args ...int) int {
	var somme int
	for i := 0; i < len(args); i++ {
		somme += args[i]
	}
	return somme
}

func main() {
	fmt.Println(computeSomme(5, 10, 20, 30))
}
