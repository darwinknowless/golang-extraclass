package main

import "fmt"

// create struct location
type location struct {
	province string
	city     string
}

// nested struct
type person struct {
	name    string
	age     int
	address struct {
		street string
	}
}

// create struct contact
type contact struct { // struct : model
	name    string
	email   string
	phone   string
	address location // => embeded struct
}

func main() {

	// inisialisasi object cara 1
	var c1 contact
	c1.name = "Darwin"
	c1.email = "darwin@gmail.com"
	c1.phone = "085750000049"
	c1.address.province = "DKI Jakarta"
	c1.address.city = "Jakarta Utara"

	// inisialisasi object cara 2
	var c2 = contact{
		name:  "Knowless",
		email: "knowless@gmail.com",
		phone: "089680000049",
		address: location{
			city:     "Yogjakarta",
			province: "DI Yogjakarta",
		},
	}
	// result
	fmt.Println(c1)
	fmt.Println(c2)

	// TODO : Anonymous struct
	var c3 = struct {
		name string
		age  int
	}{
		// name: "Knowless Darwin",
		// age: 30,
	}
	c3.name = "Darwin Knowless"
	c3.age = 29
	// result
	fmt.Println(c3)

	// TODO : nested struct
	c4 := person{
		name: "Darwin",
		age:  29,
		address: struct{ street string }{
			street: "Jl. Lodan Dalam",
		},
	}
	c4.name = "Marsya"
	c4.age = 6
	c4.address.street = " Jl. Ancol Barat"
	// result
	fmt.Println(c4)

	// TODO : list contact
	var listContact = []contact{
		{name: "Lala", email: "lala@gmail.com", phone: "123456",
			address: location{
				city:     "Jakarta Utara",
				province: "DKI Jakarta",
			}},
		{name: "Lulu", email: "lulu@gmail.com", phone: "123456",
			address: location{
				city:     "Jakarta Utara",
				province: "DKI Jakarta",
			}},
		{name: "Lili", email: "lili@gmail.com", phone: "123456",
			address: location{
				city:     "Jakarta Utara",
				province: "DKI Jakarta",
			}},
	}
	// menambahkan data
	listContact = append(listContact, contact{name: "Lolo", email: "lolo@gmail.com", phone: "123456",
		address: location{
			city:     "Jakarta Utara",
			province: "DKI Jakarta",
		}})
	// result
	fmt.Println("--------")
	fmt.Println(listContact)

	c1.Hello()
}

// method dari sebuah struct
func (c contact) Hello() {
	fmt.Println("Hello", c.name)
}
