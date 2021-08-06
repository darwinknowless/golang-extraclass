package main

import "fmt"

func main() {
	//TODO : Operator Aritmatika
	// Declare var
	var penjumlahan int = 10 + 10
	var pengurangan int = 10 - 10
	var perkalian int = 10 * 10
	var modulo int = 5 % 2
	// khusus pembagian
	num1 := 10
	num2 := 3
	var result float32 = float32(num1) / float32(num2)
	// Cetak hasil
	fmt.Println("penjumlahan", penjumlahan)
	fmt.Println("pengurangan", pengurangan)
	fmt.Println("perkalian", perkalian)
	fmt.Println("sisa bagi", modulo)
	fmt.Println("pembagian", result)
	fmt.Println()

	//TODO : Operator Logika
	perbandingan1 := 10 != 10
	var perbandingan2 bool = 10 <= 1
	var perbandingan3 bool = 10 >= 1
	fmt.Println(perbandingan1)
	fmt.Println(perbandingan2)
	fmt.Println(perbandingan3)
}
