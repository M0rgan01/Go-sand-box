package main

import "fmt"

////////// FONCTION RECURSIVE
func recursion(a int) {
	if a > 10 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
		a++
		recursion(a)
	}
}

func main() {
	recursion(0)
}
