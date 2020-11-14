package main

import "fmt"

func main() {

	msg1 := make(chan string)
	msg2 := make(chan string)

	go func() {
		msg1 <- "Message chan 1"
		msg2 <- "Message chan 2"
		msg1 <- "Message chan 1"
		msg2 <- "Message chan 2"
	}()

	for i := 0; i < 4; i++ {
		// un switch de channel
		select {
		case m1 := <-msg1:
			fmt.Println(m1)
		case m2 := <-msg2:
			fmt.Println(m2)
		}
	}
}
