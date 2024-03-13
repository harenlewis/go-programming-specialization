package main

import "fmt"

type struct {
	Name
}

func main() {
	var name string
	var address string

	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Name: ", name)

	fmt.Println("Enter address:")
	fmt.Scan(&address)
	fmt.Println("Address: ", address)

	// create a map
	// use Marshal() to create a JSON object
	// print the JSON object



}
