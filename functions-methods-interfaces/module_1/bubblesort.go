package main

import (
	"fmt"
)

func Swap(num_slice []int, i int) {
	// swap with i+1
	temp := num_slice[i]
	num_slice[i] = num_slice[i+1]
	num_slice[i+1] = temp
}

func BubbleSort(num_slice []int) {

	for i := 0; i < len(num_slice); i++ {
		for j := 0; j < len(num_slice)-i-1; j++ {
			if num_slice[j+1] < num_slice[j] {
				Swap(num_slice, j)
			}
		}
	}
}

func main() {
	num_slice := make([]int, 10)

	fmt.Printf("Enter 10 integers space separated: ")

	for i := range 10 {
		fmt.Scanf("%d", &num_slice[i])
	}

	fmt.Println("Entered numbers are:", num_slice)

	// bubble sort
	BubbleSort(num_slice)

	fmt.Println("Sorted numbers are:", num_slice)
}
