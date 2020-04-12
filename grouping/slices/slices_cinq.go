package main

import (
	"fmt"
	"strings"
)

// This function accepts a slice of strings as input
// and returns a slice of strings of the same length as output
//
func capit(names []string) []string {
	// Fetch the length of the 'names' slice using
	// the built-in Go's func len()
	// Btw cap() returns the capacity of a slice in GoLang
	//
	orig_slice_len := len(names)
	orig_slice_cap := cap(names)
	// Create the slice to be returned ... capacity & length should be the same
	// as the original slice's
	//
	capd_slice := make([]string, orig_slice_len, orig_slice_cap)
	// loop in range of orig_slice and capitalize element
	//
	for index, name := range names {
		if name == "" {
			capd_slice[index] = strings.ToUpper("uknown name")
		} else {
			capd_slice[index] = strings.ToUpper(name)
		}
	}
	// return the new slice with the capitalized names
	return capd_slice
}

func main() {
	fmt.Println("#########################################")
	// A slice is the top layer abstract structure - the underlying
	// structure of a slice is an array
	// Create a slice of length=6 and capacity=10
	// Length is the size of the slice | Capacity is the size of the underlying array
	//
	names := make([]string, 6, 10)
	names[0] = "Dim"
	names[5] = "Kacim"
	// Print-out the values only - no need to print out the indices
	//
	for _, text := range names {
		if text == "" {
			fmt.Println("No value ... Skipping ...")
		} else {
			fmt.Println(text)
		}
	}
	fmt.Println("#########################################")

	// check the capit() function ...
	//
	capd_names := capit(names)
	for _, name := range capd_names {
		fmt.Println(name)
	}
	fmt.Println("#########################################")
}

/* Smaple Output ::

#########################################
Dim
No value ... Skipping ...
No value ... Skipping ...
No value ... Skipping ...
No value ... Skipping ...
Kacim
#########################################
DIM
UKNOWN NAME
UKNOWN NAME
UKNOWN NAME
UKNOWN NAME
KACIM
#########################################

*/
