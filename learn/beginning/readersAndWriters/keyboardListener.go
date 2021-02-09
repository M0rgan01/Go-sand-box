package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// on initialise un scanner, içi sur os.Stdin (entrée standard, au clavier)
	scanner := bufio.NewScanner(os.Stdin)
	// on attend une entrée
	scanner.Scan()
	fmt.Println("Clavier --->", scanner.Text())
}
