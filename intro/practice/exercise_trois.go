package main

import (
	"fmt"
)

// declare & assign variables in the package main scope level
//
var x int = 88
var y string = "Giovanni Silva Oliveira"
var z bool = false

func main() {

	line := fmt.Sprintf("%d\n%s\n%t\n", x, y, !z)

	fmt.Println(line)
}
