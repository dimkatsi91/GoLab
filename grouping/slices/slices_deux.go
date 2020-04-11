package main

import (
	"fmt"
)


func main() {
	// a slice of values
	//
	fmt.Println("********************************************************")
	slice_une := []int{1,2,3,4,5}
	fmt.Println(slice_une)

	slice_deux := []int{}
	new_values := []int {6,7,8,9,10}
	slice_deux = append(slice_une, new_values...)
	fmt.Println("********************************************************")
	fmt.Println(slice_deux)

	// append again
	//
	fmt.Println("********************************************************")
	slice_trois := []int {}
	for times:=0; times<2; times++ {
		for i:=0; i<10; i++ {
			// no ternary operator is allowed in golang, so we have to stick with is{}else{}
			//
			if times==0 {
				slice_trois = append(slice_trois, slice_deux[i])
			} else {
				slice_trois = append(slice_trois, slice_deux[9-i])
			}
		}
	}
	fmt.Println(slice_trois)
}
