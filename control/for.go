package main

import (
	"fmt"
)

func isEven(number int) bool {
	if number % 2 == 0 {
		return true
	}
	return false
}

func main() {
	for i:=0; i<10; i++ {
		if isEven(i) == true {
			fmt.Printf("Number : %d is even\n", i)
		} else {
			fmt.Printf("Number : %d is odd\n", i)
		}
	}
}

