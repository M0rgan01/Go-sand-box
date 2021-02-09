package main

import "fmt"

////////// FONCTION AVEC RETOUR MULTIPLE
func compute(a, b int) (int, int, int) {
	somme := a + b
	multiplication := a * b
	soustraction := a - b
	return somme, multiplication, soustraction
}

func main() {
	////////// RETOUR DE FONCTION MULTIPLE
	fmt.Println(compute(5, 10))

	somme, multiplication, soustraction := compute(5, 10)

	fmt.Println(somme, multiplication, soustraction)

	somme2, _, _ := compute(5, 10)

	fmt.Println(somme2)
}
