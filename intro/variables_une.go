package main

import (
	"fmt"
)

type Number struct {
	integer int
	double float64
}

func (num * Number) getNums() (int, float64) {
	return num.integer, num.double
}


func print_num(num * Number, choice int) {

	integer, double := num.getNums()

	switch choice {
	case 1:
		fmt.Printf("Integer Number = %d\n", integer)
	case 2:
		fmt.Printf("Float64 Number = %f\n", double)
	default:
		fmt.Println("Please choose 1 for integer | 2 for double number printout and try again ...")
	}
}

func main() {
	
	Object := Number {
		integer : 10,
		double : 3.141593,
	}

	for i:=0; i<4; i++ {
		print_num(&Object, i)
	}

}