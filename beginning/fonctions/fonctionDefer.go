package main

import "fmt"

func premier() {
	fmt.Println("Un")
}

func deuxième() {
	fmt.Println("Deux")
}

func troisième() {
	fmt.Println("Trois")
}

func end() {
	fmt.Println("Fin de programme")
}

func executeAll() {
	// defer execute la méthode en dernier dans la méthode en cour
	// si il y a plusieurs defer, le 1er est dernier, le 2eme avant-dernier, ect...
	// utile pour directement close un fichier, une db... juste après un appel
	defer premier()
	defer deuxième()
	troisième()
}

func main() {
	executeAll()
	end()
}
