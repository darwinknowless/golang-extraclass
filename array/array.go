package main

import "fmt"

func main() {
	// [array length] type data {values}
	var arr1 = [3]int{} // array statis
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	// [...] => operator spread / rest
	arr2 := [3]int{1, 2, 3}
	// => array kosong
	var arr3 = make([]int, 3) // => array slice/dinamis
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3

	// arr1 = append(arr1, 10)

	fmt.Println("array 01", arr1)
	fmt.Println("array 02", arr2)
	fmt.Println("array 03", arr3)

	fmt.Println("________")

	// slice : lebih dinamis daripada array
	//? [:] => [start : end]
	slice1 := arr1[:]
	fmt.Println("slice 01", slice1)
	fmt.Println(cap(slice1))
	// slice1 = append(slice1, 5) // append : tambah dkapasitas, cap awal dikali 2
	// fmt.Println(cap(slice1))
	fmt.Println("_________")

	slice2 := []int{1, 2, 3, 4}
	fmt.Println("slice 02", slice2)
	fmt.Println(len(slice2)) // len : panjang array
	fmt.Println(cap(slice2)) // cap : kapasitas yang bisa di tampung

	slice2 = append(slice2, 5)
	fmt.Println("_________")
	fmt.Println("slice 02", slice2)
	fmt.Println(len(slice2)) // len : panjang array
	fmt.Println(cap(slice2)) // cap : kapasitas yang bisa di tampung

	slice2 = append(slice2, 6)
	fmt.Println("_________")
	fmt.Println("slice 02", slice2)
	fmt.Println(len(slice2)) // len : panjang array
	fmt.Println(cap(slice2)) // cap : kapasitas yang bisa di tampung
}
