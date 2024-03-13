package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string // `json:"tag_name"`
	Address string // `json:"tag_address"`
}

func main() {
	var name string
	var address string

	fmt.Println("Enter name:")
	fmt.Scan(&name)
	// fmt.Println("Name: ", name)

	fmt.Println("Enter address:")
	fmt.Scan(&address)
	// fmt.Println("Address: ", address)

	// create a map
	personMap := make(map[string]string)
	personMap["name"] = name
	personMap["address"] = address

	// use Marshal() to create a JSON object
	// u, err := json.Marshal(Person{Name: name, Address: address})
	u, err := json.Marshal(personMap)

	if err != nil {
		panic(err)
	}

	// print the JSON object
	fmt.Println(string(u))
}
