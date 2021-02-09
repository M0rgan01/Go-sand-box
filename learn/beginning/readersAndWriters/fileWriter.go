package main

import "os"

func main() {
	// si le fichier n'existe pas, il le créé (os.O_CREATE)
	file, _ := os.OpenFile("./beginning/readersAndWriters/test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	defer file.Close()

	file.WriteString("Un test \n")
}
