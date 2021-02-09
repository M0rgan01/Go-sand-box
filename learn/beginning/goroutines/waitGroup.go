package main

import (
	"fmt"
	"sync"
)

func main() {
	a := sync.WaitGroup{}

	// Nombre de wait attendu
	a.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("test")
			a.Done()
		}()
	}

	// ici la commande wait attend 5 Done
	a.Wait()
	fmt.Println("Done")
}
