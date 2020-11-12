package main

import "fmt"

func main() {
	////////// FONCTION "CLOSER"
	var result = func(a, b int) int {
		return a * b
	}(10, 20)

	fmt.Println(result)

	func() {
		fmt.Println("Hello world in closer")
	}()

	var function = func() {
		fmt.Println("Hello world in closer 2")
	}

	function()

	var functionMultiply = func(a, b int) {
		fmt.Println(a * b)
	}

	functionMultiply(10, 20)
}
