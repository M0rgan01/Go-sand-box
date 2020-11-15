package main

import (
	"fmt"
	"sort"
)

func main() {

	coll1 := []int{1, 2, 5, 3, 6, 4}
	coll2 := []string{"z", "a", "x", "g", "b", "t"}

	sort.Ints(coll1)
	sort.Strings(coll2)

	fmt.Println(coll1)
	fmt.Println(coll2)
}
