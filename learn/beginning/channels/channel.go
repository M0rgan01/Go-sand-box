package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	// içi un canal qui n'accepte que des sorties
	//messages := make(<- chan string)
	// içi un canal qui n'accepte que des entrées
	//messages := make(chan <- string)

	go func() {
		// ici nous assignons une valeur à un canal avec "canal <- valeur"
		// on ne peut ajouter un valeur dans un canal uniquement dans une goroutine
		messages <- "un"
		messages <- "deux"
	}()

	// ici nous récupérons la 1er valeur assigner à un canal,
	// si aucune valeur n'est stocker, le comportement est semblable à un await
	fmt.Println(<-messages)
	fmt.Println("-------")
	fmt.Println(<-messages)
	fmt.Println("-------")

	fmt.Println("Fin Channel messages")

	readOnlyChannel := createReadOnlyChannel()

	fmt.Println(<-readOnlyChannel)
	fmt.Println(<-readOnlyChannel)
}

func createReadOnlyChannel() <-chan int {
	channel := make(chan int)
	go func() {
		channel <- 1
		channel <- 2
	}()
	return channel
}
