package main

import "fmt"

func main() {
	var value string
	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&value)
	fmt.Printf("Number is %s", value)
}
