package main

import (
	"fmt"
	"math"
)

func main() {
	println("Hello world")

	// dans une déclaration de var sans type -> le compilateur choisi le stockage (int8, int16..) et le type
	var numberWithoutType = 12

	// pareil que la déclaration avec var, mais sans typage possible
	numberWithoutTypeAndDeclaration := 12

	var numberWithType uint8 = 12

	println(numberWithoutType)
	println(numberWithoutTypeAndDeclaration)
	println(numberWithType)

	println("Valeur max d'un integer64 :", math.MaxInt64)

	// println est un output console de 'debug' ou de test, sur un affichage simple, utiliser fmt
	println(12.5)
	fmt.Println(12.5)
}
