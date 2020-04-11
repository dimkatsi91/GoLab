package main

import "fmt"

////////////////////////////////////
var integer int = 10

type myInteger int

var myInt myInteger = 20
/////////////////////////////////


func main() {
	
	fmt.Println(myInt)
	fmt.Printf("Type[integer] : %T\n", myInt)

	fmt.Println("############ Conversion IN Golang ############")

	integer = int(myInt) + 80
	fmt.Println(integer)
	fmt.Printf("Type[integer] : %T\n", integer)
}