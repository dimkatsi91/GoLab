package main

import (
	"fmt"
)

// an interface for a Point
//
type Point interface {
	area() float64
}

// A Circle struct
//
type Circle struct {
	x      int
	y      int
	radius int
}

// Circle implements Point
//
// area(Circle) = p * r^2
//
func (c Circle) area() float64 {
	return (3.141593 * float64(c.radius*c.radius))
}


func info(p Point) {
	switch p.(type) {
	case Circle:
		fmt.Printf("This is a Circle with radius: %d\n", p.(Circle).radius)
	default:
		fmt.Println("Please specify a Rectangle or a Circle and try again!")
		return
	}
	fmt.Println(p.area())
}


//
// main
//
func main() {
	c1 := new(Circle)
	c1.x = 1
	c1.y = 1
	c1.radius = 1

	// c1 is a of type *Circle, so dereference it in order to be assigned in c2 
	c2 := Circle(*c1)
	fmt.Printf("Type of c1 is :: %T\n", c1)
	info(c2)

	c3 := *c1
	c3.radius = 3
	fmt.Printf("Type of c3 is :: %T\n", c3)
	info(c3)

	c4 := Point(*c1)
	// ./point_trois.go:62:4: c4.radius undefined (type Point has no field or method radius)
	// c4.radius = 3
	fmt.Printf("Type of c4 is :: %T\n", c4)
	info(c4)

}


/* Printout Sample ::

Type of c1 is :: *main.Circle
This is a Circle with radius: 1
3.141593
Type of c3 is :: main.Circle
This is a Circle with radius: 3
28.274337
Type of c4 is :: main.Circle
This is a Circle with radius: 1
3.141593

*/
