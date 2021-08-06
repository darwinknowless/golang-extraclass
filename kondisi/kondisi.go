package main

import "fmt"

func main() {
	// if else
	var1 := 1
	var2 := 2
	if var1 != var2 {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}

	// switch case 01
	switch var1 {
	case 1:
		fmt.Println("Benar")
	case 2:
		fmt.Println("Salah")
	}

	// switch case 02
	switch {
	case (var1 == var2):
		fmt.Println("Salah")
	case (var2 != var1):
		fmt.Println("Benar")
	}

}
