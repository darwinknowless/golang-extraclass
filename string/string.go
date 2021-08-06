package main

import "fmt"

func main() {
	// hanya bisa menggunakan "" dan `` jika string
	var string1 string = "Menggunakan Double Quotes"
	var string2 string = `Menggunakan Backstick`

	fmt.Println(string1)
	fmt.Println(string2)
}
