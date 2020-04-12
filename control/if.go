package main

import "fmt"

func isDivisibleByThree(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}

func main() {

	var sum int = 0
	N := 33

	for i := 1; i <= N; i++ {
		if isDivisibleByThree(i) != false {
			fmt.Printf("Number %d is divisible by 3\n", i)
			sum++
		}
	}

	fmt.Printf("[1-%d] : There are %d numbers divisble by 3\n", N, sum)
}
