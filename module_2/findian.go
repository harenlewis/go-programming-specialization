package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string
	fmt.Println("Enter a string")

	fmt.Scan(&word)

	fmt.Println("String:", word[:1])
	first_char := word[:1]
	last_char := word[len(word)-1:]

	if first_char == "i" && last_char == "n" && strings.Contains(word, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
