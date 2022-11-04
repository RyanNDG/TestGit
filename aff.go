package main

import "fmt"

func main() {

	var x string
	var m = ""
	fmt.Print("Ecrire un mot: ")
	fmt.Scanln(&x)

	for i := len(x) - 1; i > len(x)+1; i-- {

		m = m + (string(x[i]))
		fmt.Println("", m)

	}

}
