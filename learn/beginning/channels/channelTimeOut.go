package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		time.Sleep(time.Second * 4)
		c <- true
	}()

	select {
	case channelEntry := <-c:
		fmt.Println("Channel is complete", channelEntry)
	// ajout d'un timeOut dans le cas ou un des appel est trop long
	case <-time.After(time.Second * 2):
		fmt.Println("Too long call, closing channel")
	}

	fmt.Println("Finish")
}
