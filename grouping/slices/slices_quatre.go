package main

import (
	"fmt"
)


func main() {
	// a slice of values
	//
	fmt.Println("********************************************************")
	slice_une := make([]string, 0, 5)
	players := []string{"giovanni", "rivaldo", "karembue", "djordjevic", "galletti"}
	slice_une = append(slice_une, players...)
	fmt.Println(slice_une)
	fmt.Printf("slice_une length = %d | capacity = %d\n", len(slice_une), cap(slice_une))

	new_players := []string {"stelios", "zetterberg", "ze elias", "nikopolidis", "kovacevic"}
	slice_une = append(slice_une, new_players...)
	fmt.Println("********************************************************")
	fmt.Println(slice_une)
	fmt.Printf("slice_une length = %d | capacity = %d\n", len(slice_une), cap(slice_une))
}
