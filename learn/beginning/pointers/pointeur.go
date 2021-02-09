package main

import "fmt"

func increment(a int) {
	a++
}

func incrementWithRef(a *int) {
	*a++
}

func main() {

	test := 0

	increment(test)

	// Le résultat attendu est 1, mais les paramètres normaux sont 'copié' et ne modifies pas les attributs externe
	fmt.Println(test)

	// La méthode incrementWithRef prend une REF d'attribut (* pour un paramètre de méthode)
	// Pour donner la ref (adresse en mémoire) d'un attribut -> &attribut
	// Pour récupérer la valeur d'une ref -> *attribut
	incrementWithRef(&test)

	// içi le résultat est 1
	fmt.Println(test)

	// autres méthode
	testPointer := &test
	incrementWithRef(testPointer)
	fmt.Println(test)

	// Création Direct d'une ref
	var pointer = new(int)
	*pointer++
	fmt.Println(*pointer)
}
