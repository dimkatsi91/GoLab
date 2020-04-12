package main

import (
	"fmt"
)

func main() {
	// a slice of values
	//
	fmt.Println("********************************************************")
	slice_une := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice_une)

	// for loop in slices
	//
	fmt.Println("********************************************************")
	slice_deux := [10]int{}
	for index, value := range slice_une {
		slice_deux[index] = value * 2
	}

	for _, val := range slice_deux {
		fmt.Println(val)
	}

	// half of slice_deux is slice_trois
	//
	fmt.Println("********************************************************")
	slice_trois := slice_deux[5:]
	fmt.Println(slice_trois)
}
