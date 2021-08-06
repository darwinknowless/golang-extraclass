package main

import "fmt"

func main() {
	// call function
	var baru = hello2("Darwin")
	fmt.Println(baru)

	var name, age = hello3("Darwin", 29)
	fmt.Println(name, age)

	listNumber(1, 2, 3, 4, 5, 6, 7, 8)

	var closure1 = func(name string) {
		fmt.Println("Ini closure", "Hello", name)
	}
	closure1("Darwin")

}

//function hello1
func hello1(name string) {
	fmt.Println("Hello", name)
}

//function hello2
func hello2(name string) string {
	greetings := "Hello " + name
	return greetings
}

//function hello3
func hello3(name string, age int) (string, int) {
	return name, age
}

// variadic function: parameter bisa banyak
func listNumber(number ...int) {
	fmt.Println(number)
}
