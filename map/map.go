package main

import "fmt"

func main() {
	// var map1 = map [type data key]type data value {values}
	map1 := map[string]string{
		"name": "DK",
		"age":  "29",
	}
	fmt.Println(map1)
	fmt.Println(map1["name"])

	// var map2 = map [type data key]type data any {} {values}
	map2 := map[string]interface{}{
		"name": "Darwin",
		"age":  30,
	}
	fmt.Println(map2)
	fmt.Println(map2["age"])

	// array map
	map3 := map[string]string{
		"name": "Knowless",
		"age":  "31",
	}

	arr1 := []map[string]string{map1, map3}
	fmt.Println(arr1)

	arr2 := []interface{}{map1, map2}
	fmt.Println(arr2)

}
