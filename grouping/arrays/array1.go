package main

import (
	"fmt"
)

func main() {
	// arrays are not used in Go
	// slices are used instead
	//
	fmt.Println("********************************************************")
	var zeros [10]int
	fmt.Println(zeros)

	// ones
	fmt.Println("********************************************************")
	var ones [10]int
	for i := 0; i < 10; i++ {
		ones[i] = i + 10
		fmt.Printf("ONES[%d] = %d\n", i, ones[i])
	}
	fmt.Println("********************************************************")
	fmt.Println(ones)
}
