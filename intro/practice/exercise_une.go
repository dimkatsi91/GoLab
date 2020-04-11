package main

import (
	"fmt"
)

func main() {
	// short declaration operator usage
	x := 42
	y := "Giovanni" + " " + "Silva"
	z := false

	fmt.Printf("x = %d\ny = %s\nz = %t\n", x, y, !z)
}