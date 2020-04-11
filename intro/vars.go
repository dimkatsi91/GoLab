package main

import (
	"fmt"
)

func main() {
	// var keywoard
	var pi float64 = 3.141593
	var integer int = 8
	var myName string = "Dimos"

	// update a string
	myName += " " + "Kacimardos"

	fmt.Printf("INteger = %d\n", integer)
	fmt.Printf("Floating-64 PI = %f\n", pi)
	fmt.Printf("Hello, my name is %s. Nice to meet You!\n", myName)
}