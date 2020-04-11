package main

import (
	"fmt"
)

// A Circle struct
//
type Circle struct {
	x int
	y int
	radius int
}

// area(Circle) = p * r^2
//
func (c *Circle) area() float64 {
	return (3.141593 * float64(c.radius * c.radius))
}

// A Rectangle is also a Point
//
type Rectangle struct {
	side int
}

func (r *Rectangle) area() float64 {
	return float64(r.side * r.side)
}
// an interface for a Point
//
type Point interface {
	area() float64
}

func info(p Point) {
	fmt.Println(p)
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
	info(&r1)
	
	// a circle
	//
	fmt.Println("***************************************************************")
	c1 := Circle{
		x : 1,
		y : 1,
		radius : 1,
	}
	//fmt.Println(c1.area())
	info(&c1)
}