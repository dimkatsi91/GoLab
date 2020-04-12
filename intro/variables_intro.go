package main

import (
	"fmt"
)

func main() {
	//
	// short declaration operator
	//
	integer := 10        // int
	float := 3.141593    // float64
	first := "Dimos"     // string
	last := "Kacimardos" // string

	wholeName := first + " " + last // concatenation two strings

	fmt.Printf("INteger = %d\n", integer)
	fmt.Printf("Floating-64 PI = %f\n", float)
	fmt.Printf("Hello, my name is %s. Nice to meet You!\n", wholeName)
}
