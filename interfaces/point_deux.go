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
	x int
	y int
	radius int
}

// Circle implements Point
//
// area(Circle) = p * r^2
//
func (c Circle) area() float64 {
	return (3.141593 * float64(c.radius * c.radius))
}

// A Rectangle is also a Point
//
type Rectangle struct {
	side int
}

// Circle implements Point
//
func (r Rectangle) area() float64 {
	return float64(r.side * r.side)
}

func info(p Point) {
	switch p.(type) {
	case Rectangle:
		fmt.Printf("This is a Rectangle with side: %d\n", p.(Rectangle).side)
	case Circle:
		fmt.Printf("This is a Circle with radius: %d\n", p.(Circle).radius)
	default:
		fmt.Println("Please specify a Rectangle or a Circle and try again!")
	}
	fmt.Println(p.area())
}


//
// main
//
func main() {
	// a rect
	//
	fmt.Println("***************************************************************")
	r1 := Rectangle{10}
	info(r1)
	
	// a circle
	//
	fmt.Println("***************************************************************")
	c1 := Circle{
		x : 1,
		y : 1,
		radius : 1,
	}
	info(c1)
}