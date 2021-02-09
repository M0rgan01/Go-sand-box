package main

import (
	"fmt"
	"time"
)

func roll(name string, loops int, sec int, count *int) {
	for i := 1; i <= loops; i++ {
		fmt.Printf("%s ---> %d \n", name, i)
		time.Sleep((time.Second / 2) * (time.Duration(sec)))
		*count++
	}
}

func main() {
	fmt.Println("Début")

	count := 0

	// async
	go roll("Boucle 1", 10, 1, &count)
	go roll("Boucle 2", 10, 1, &count)

	// simulation d'attente pour le process
	for count != 20 {
	}
	//time.Sleep((time.Second / 2) * 10)

	fmt.Println("Fin")
}
