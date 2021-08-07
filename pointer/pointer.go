// pointer : reference
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	// var p1 int = 20
	// var p2 int = p1 // pass by reference

	// // *int: pointer integer
	// // &p1: ampersand
	// // variabel pointer
	// var p3 *int = &p1
	// // (&p1 != nilai p1 20)
	// // (&p1 alamat memory p1)

	// p1 = 30

	// fmt.Println(&p1) // cara print alamat memory
	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(*p3) // cara print pointer integer

	var p1 = person{"Darwin", 29}
	var p2 = p1

	var p3 *person = &p1

	p1.name = "Knowless"
	p3.name = "DK"

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
