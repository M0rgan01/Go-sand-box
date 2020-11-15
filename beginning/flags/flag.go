package main

import (
	"flag"
	"fmt"
)

func main() {

	// go run flag.go -val=10
	// go run flag.go -help

	flag1 := flag.Int("val", 1 /* valeur par défaut */, "Nombre à afficher")
	flag.Parse()
	fmt.Println(*flag1)
}
