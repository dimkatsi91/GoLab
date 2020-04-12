package main

import (
	"fmt"
)

var PI float64 = 3.141593

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	// area = pi*r^2
	return c.radius * c.radius * PI
}

func (s *square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println("Areas is: ", s.area())
}

func main() {
	c := circle{1}
	sq := square{2}

	info(&c)
	info(&sq)
}
