package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	reader, _ := os.OpenFile("./beginning/readersAndWriters/test.txt", os.O_RDONLY, 0777)
	defer reader.Close()

	// méthode 1 (pas très simple...)

	/*data := make([]byte, 512)
	reader.Read(data)*/
	// méthode 2
	data, _ := ioutil.ReadAll(reader)

	fmt.Println(string(data))
}
