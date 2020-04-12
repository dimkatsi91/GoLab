package main

import (
	"fmt"
)

func add(a, b int) int {
	fmt.Printf("Adding %d + %d from standard function ... ", a, b)
	return a + b
}

func main() {
	var sum int = add(5, 15)
	fmt.Printf(" is --> %d\n", sum)

	// Anonymous function usage
	// ENJOY dimos !! fourty said ! Mf@@@r sick b@@@d
	new_sum := func(a, b int) int {
		fmt.Printf("Adding %d + %d from anonymous function ... ", a, b)
		return a + b
	}(75, 25)
	fmt.Printf(" is --> %d\n", new_sum)
}
