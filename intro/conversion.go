package main

import (
	"fmt"
)

// pi is of Type float64
//
var pi float64 = 3.141593

// MyFloat64 is another type for loat64 numbers | Not the same as a float64
//
type MyFloat64 float64

// declare a MyFloat type variable and assign its value too
//
var myPI MyFloat64 = 3.1416

func main() {
	// print-out some info regarding pi variable
	//
	fmt.Printf("Pi is of type [ %T ] and has value: %f\n", pi, pi)
	// print out some info regarding myPI variable
	//
	fmt.Printf("myPI is of type [ %T ] and has value: %f\n", myPI, myPI)

	// Convert myPI to type float64 and add two pi's
	//
	double_pi := pi + float64(myPI)
	fmt.Printf("pi + float(myPI) = %f\n", double_pi)
}
