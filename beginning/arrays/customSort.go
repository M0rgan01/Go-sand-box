package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return i > j
	} else {
		return len(s[i]) < len(s[j])
	}
}

func main() {
	coll := byLength{"zz", "aaaaaa", "xx", "g", "bbb", "tttt"}
	sort.Sort(coll)
	fmt.Println(coll)
}
