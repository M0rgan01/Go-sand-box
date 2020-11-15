package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var counter uint64 = 0

	var mutex = sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				// un mutex est présent pour bloquer les autres goroutines en cour,
				// dans le cas ou plusieurs goroutines traitent une même variable,
				// il y aura des conflits (accès simultanés à la même ressource)
				mutex.Lock()
				counter++
				// ne pas oublier de débloquer les autres goroutines
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter)
}
