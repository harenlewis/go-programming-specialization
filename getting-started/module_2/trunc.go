package main

import (
	"fmt"
	"strconv"
)

func main() {
	var value string
	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&value)
	fmt.Println("Entered Number is", value)

	float_value, err := strconv.ParseFloat(value, 64)
	trunc_value := int(float_value)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Truncated Number is %s", strconv.Itoa(trunc_value))

}
