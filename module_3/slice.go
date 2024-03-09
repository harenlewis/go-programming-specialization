package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	num_slice := make([]int, 0, 3)

	for {
		var number string

		fmt.Println("Enter the numbers:")
		fmt.Scan(&number)
		fmt.Println("Entered number:", number)

		if strings.ToUpper(number) == "X" {
			fmt.Println("Peace out!!!!")
			break
		}

		num, _ := strconv.Atoi(number)

		// add to slice
		num_slice = append(num_slice, num)

		// sort the slice in ascending order
		sort.Slice(num_slice, func(i, j int) bool {
			return num_slice[i] < num_slice[j]
		})

		// print the slice
		fmt.Println("Sorted Slice: ")
		fmt.Println(num_slice)
	}
}
