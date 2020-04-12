package main

import (
	"fmt"
)

type Circle struct {
	radius int
}

// Factory of a Circle | Takes a parameter and returns a Circle Object ref
//
func CircleFactory(radius int) *Circle {
	return &Circle{
		radius: radius,
	}
}

// Attach next method to Circle struct
//
func (c *Circle) print() {
	fmt.Printf("This is a Circle with radius : %d\n", c.radius)
}

// ANother function to return a float64 2*PI*r , that is perimeter of a circle
// Attach this function with the Circle struct
//
func (c *Circle) getPerimeter() float64 {
	var PI float64 = 2.141593
	// Or more compact way to declare and initialize this variable ::
	// PI := 2.141593
	return float64(2*c.radius) * PI
}

///////////
// MAIN //
/////////
func main() {
	circle := CircleFactory(1)
	circle.print()
	fmt.Printf("Circle perimeter is (2*PI*radius) = %f\n", circle.getPerimeter())
}
