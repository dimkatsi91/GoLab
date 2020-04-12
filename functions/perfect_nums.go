package main

import "fmt"

func isPerfect(num int) bool {
	// initialized to zero
	var sum int
	// 6 = 1 + 2 + 3 (sum of its divisors)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	if sum == num {
		return true
	}
	return false
}

func main() {
	var perfect_nums_sum int
	for i := 1; i < 10000; i++ {
		if isPerfect(i) == true {
			fmt.Printf("Number %d is a perfect number\n", i)
			perfect_nums_sum += 1
		}
	}
	fmt.Printf("Perfect number in [1,10000] = %d\n", perfect_nums_sum)
}
