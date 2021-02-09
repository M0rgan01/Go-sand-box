package main

import "fmt"

type inter interface {
	hello()
}

type Lapin struct {
	nom string
}

type Chat struct {
	age int
}

func (c Chat) hello() {
	fmt.Println("Hello from cat")
}

func (l Lapin) hello() {
	fmt.Println("Hello from rabbit")
}

func interfaceTest(i inter) {
	i.hello()
}

func main() {
	chat := Chat{10}
	lapin := Lapin{"Pirate"}

	interfaceTest(chat)
	interfaceTest(lapin)

}
