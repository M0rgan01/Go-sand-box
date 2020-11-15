package main

import (
	"fmt"
)

func main() {
	testPanic()

	boolean, err := testError("tests")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(boolean)
	}

	fmt.Println("Fin")
}

func testPanic() {
	defer func() {
		// recover fait en sorte que panic n'arrête que le bloc en cour
		// il doit être placé dans une fonction, avec defer
		recover()
	}()

	// force l'arrêt du programme
	panic("Une Error")
}

func testError(s string) (bool, error) {
	if s == "test" {
		return true, nil
	} else {
		return false, monErrorPerso{id: 10, message: "error"}
		// ou return false, fmt.Errorf("---> %s", "une erreur")
		// ou return false, errors.New("une erreur")
	}
}

type monErrorPerso struct {
	id      int
	message string
}

func (e monErrorPerso) Error() string {
	return fmt.Sprintf("Erreur type %d : %s", e.id, e.message)
}
