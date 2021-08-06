// Parsing Type Data
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var string1 string = "10"
	var string2 string = "4"

	numStr1, err := strconv.ParseInt(string1, 10, 32)
	if err != nil {
		fmt.Println("error parsing", err)
	}

	numStr2, _ := strconv.ParseInt(string2, 10, 32)
	if err != nil {
		fmt.Println("error parsing", err)
	}

	result1 := numStr1 + numStr2
	fmt.Println(result1)
}
