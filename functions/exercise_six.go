package main

import (
	"fmt"
)

func main() {
	fact := func(num int) {
		f := 1
		for i:=1; i<=num; i++ {
			f *= i
		}
		fmt.Printf("%d! = %d\n", num, f)
	}

	for number:=1; number<11; number++ {
		fact(number)
	}
}