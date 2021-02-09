package main

import (
	"fmt"
	"os/exec"
)

func main() {

	// ouverture explorer

	// cmd := exec.Command("nautilus",  "--browser",  "./")
	// cmd.Run()

	// execution d'une commande

	out, _ := exec.Command("ls", "..").Output()
	fmt.Println(string(out))
}
