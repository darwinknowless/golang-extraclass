package main

import "fmt"

func main() {

	arr1 := [3]int{1, 2, 3}
	// string1 := "string"

	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(arr1[i])
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("perulangan ke-", i)
	// }

	// for i := 10; i >= 1; i-- { // => perulangan terbalik
	// 	fmt.Println("perulangan ke-", i)
	// }

	// for index, value := range string1 {
	// 	fmt.Println(index, string(value))
	// }

	// for index, value := range arr1 {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range arr1 {
	// 	fmt.Println(value)
	// }

	for _, value := range arr1 {
		if value == 1 {
			fmt.Println("Benar")
			break //  break:
			// or
			// continue: ngeskip perulangan ssat ini dan lanjut ke proses berikutnya
		}
	}
}
