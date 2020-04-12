package main

import "fmt"

func evenOrNot(number int) bool {
	if res := number % 2; res == 0 {
		return true
	}
	return false
}

func main() {
	for num := 1; num < 11; num++ {
		//var condition bool = evenOrNot(num)
		//switch condition {
		switch evenOrNot(num) {
		case true:
			fmt.Printf("Number: %d is even\n", num)
		case false:
			fmt.Printf("Number: %d is odd\n", num)
		default:
			fmt.Println("Never reached ...")
		}
	}
}
