package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	FirstName string
	LastName  string
}

func main() {
	var filename string
	name_slice := make([]Name, 0, 1)

	fmt.Println("Enter file name:")

	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fullname := scanner.Text()
		name_arr := strings.Split(fullname, " ")
		name_slice = append(name_slice, Name{FirstName: name_arr[0], LastName: name_arr[1]})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("The names are:")
	for _, name := range name_slice {
		fmt.Println(name.FirstName, name.LastName)
	}
}
