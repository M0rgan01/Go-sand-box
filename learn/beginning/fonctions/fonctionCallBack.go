package main

import "fmt"

////////// FONCTION AVEC AUTRE FONCTION EN PARAMS
func functionWithFunction(a, b int, function func(x, y int)) {
	function(a, b)
}

////////// FONCTION AVEC UN RETOUR FONCTION
func multiplyBuilder(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

func main() {
	functionWithFunction(10, 200, func(x, y int) { fmt.Println("Appel d'un fonction en params :", x*y) })

	multiplyByTen := multiplyBuilder(10)
	multiplyByTwenty := multiplyBuilder(20)

	fmt.Println(multiplyByTen(10), multiplyByTwenty(20))

	// OU

	multiplyByTen2 := multiplyBuilder(10)(10)
	multiplyByTwenty2 := multiplyBuilder(20)(20)

	fmt.Println(multiplyByTen2, multiplyByTwenty2)
}
