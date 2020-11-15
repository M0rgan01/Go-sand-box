package main

import (
	"flag"
	"fmt"
)

func main() {

	// go run arg.go "test" "test2" 123

	flag.Parse()
	flag1 := flag.Args()
	fmt.Println(flag1)
}
