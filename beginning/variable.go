package main

import (
	"fmt"
	"math"
)

func main() {

	// dans une déclaration de var sans type -> le compilateur choisi le stockage (int8, int16..) et le type
	var numberWithoutType = 12
	// pareil que la déclaration avec var, mais sans typage possible
	numberWithoutTypeAndDeclaration := 12
	// déclaration avec type
	var numberWithType uint8 = 12
	// déclaration de constante
	const myConst = 12

	println(numberWithoutType, numberWithoutTypeAndDeclaration, numberWithType)

	// déclaration multiple
	var (
		variable1         = "coucou"
		variable2         = true
		variable3 float32 = 1.26
	)

	println(variable1, variable2, variable3)

	println("Valeur max d'un integer64 :", math.MaxInt64)

	// println est un output console de 'debug' ou de test, sur un affichage simple, utiliser fmt
	println(12.5)
	fmt.Println(12.5)

	// String

	var stringWithReturn = "Test\n"
	var stringWithBackTick = `Test\n`
	var stringWithBackTickAndText = `Test "coucou"`
	var stringLength = len(stringWithBackTickAndText)

	fmt.Print(stringWithReturn, stringWithBackTick, stringWithBackTickAndText)
	fmt.Println()
	fmt.Println(stringLength)

	// CAST

	var a float32 = 12.5
	var b float64 = 13.5
	var c = float64(a)

	fmt.Println(b * c)
}
