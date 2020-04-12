package main

import (
	"fmt"
)

func fact(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fact(num-1) * num
}

func main() {
	for num := 0; num < 21; num++ {
		fmt.Println("Factorial[", num, "] = ", fact(num))
	}
}
