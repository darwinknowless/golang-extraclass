package main // wajib ada

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// log.Println("Hello World")

	const var1 = "Constant" //? Const value

	var varNoValue string //? Deklarasi variabel tanpa nilai

	var var2 string = "Hello" //? Deklarasi variabel  1
	fmt.Println(var2)
	
	var3 := 3.14 //? Deklarasi variabel 2 (menyesuaikan dengan value)
	fmt.Println(var3)

	var2 = "Hello Lagi"
	var3 = 4.14

	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(varNoValue)
	fmt.Println(var1)
}