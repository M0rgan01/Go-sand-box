package main

import (
	"fmt"
)

func rollWithClose(loops int, c *chan string) {
	for i := 1; i <= loops; i++ {
		*c <- fmt.Sprintf("Message %d", i)
	}
	// ici l'instruction close permet d'indiquer à la boucle qu'il y a plus rien dans le chan
	close(*c)
}

func main() {

	messages := make(chan string)

	go rollWithClose(5, &messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}
