package main

import "fmt"

////////// FONCTION TRADITIONNEL
func multiply(a int, b int) int {
	return a * b
}

// autre syntaxe des params
func multiply2(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 10))
}
