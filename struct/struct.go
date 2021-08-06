package main

import "fmt"

type contact struct { //struct : model
	name  string
	email string
	phone string
}

func main() {
	var c1 contact
	c1.name = "Darwin"
	c1.email = "darwin@gmail.com"
	c1.phone = "085750000049"

	var c2 = contact{"Knowless", "knowless@gmail.com", "089680000049"}

	var c3 = contact{
		email: "dk@gmail.com",
		phone: "085780000049",
	}

	var c4 = contact{
		name: "DK",
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)

}
