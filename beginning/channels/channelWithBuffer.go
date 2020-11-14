package main

import "fmt"

func main() {

	// le 2eme est un buffer, qui indique le nombre maximum d'élément dans un channel à la fois
	// un channel avec buffer n'est pas obliger d'être incrémenter dans une goroutine
	messages := make(chan string, 2)

	messages <- "Un"
	messages <- "Deux"

	fmt.Println(<-messages)

	messages <- "Trois"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
