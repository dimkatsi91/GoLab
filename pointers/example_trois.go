package main

import (
	"fmt"
	// "math"
)

/* 
	2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
	What is the sum of the digits of the number 2^1000?
*/

func powerDigitSum(num int) *int {
	
	var sum int

	// keep a backup
	// num_bkp := num
	for(num != 0) {
		sum += num%10
		num /= 10
	}
	// fmt.Printf("%d | Power Digit Sum = %d\n", num_bkp, sum)
	return &sum
}


func main() {

	num := 32768
	var result_ptr *int = powerDigitSum(num)

	fmt.Printf("Power Digit Sum(%d) = %d\n", num, *result_ptr)
	fmt.Printf("Address of %d is --> %v\n", *result_ptr, result_ptr)

}