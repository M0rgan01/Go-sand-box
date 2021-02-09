package main

import (
	"fmt"
	"time"
)

// Reprise/amélioration du scénario dans goroutine.go

func roll(name string, loops int, sec int, c *chan bool) {
	for i := 1; i <= loops; i++ {
		fmt.Printf("%s ---> %d \n", name, i)
		time.Sleep((time.Second / 2) * (time.Duration(sec)))
	}
	*c <- true
}

func main() {

	interrupt := make(chan bool)

	// async
	go roll("Boucle 1", 7, 1, &interrupt)
	go roll("Boucle 2", 5, 1, &interrupt)

	<-interrupt
	<-interrupt

	fmt.Println("Fin")
}
