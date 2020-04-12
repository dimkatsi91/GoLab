package main

import (
	"fmt"
)

func main() {
	// a slice of values
	//
	fmt.Println("********************************************************")
	slice_une := []int{1, 2, 3, 4, 5}
	fmt.Println(slice_une)

	slice_deux := []int{}
	new_values := []int{6, 7, 8, 9, 10}
	slice_deux = append(slice_une, new_values...) // {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("********************************************************")
	fmt.Println(slice_deux)

	// delete from the slice_deux a few ints | Just #2 first and #2 last ints
	//
	slice_deux = append(slice_deux[:2], slice_deux[8:]...)
	fmt.Println(slice_deux)
	fmt.Printf("SLICE_DEUX LEN IS NOW ---> %d\n", len(slice_deux))
}
