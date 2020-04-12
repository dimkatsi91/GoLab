package main

import (
	"fmt"
)

// define complex number struct
type Complex struct {
	real int
	imag int
}

// Print complex number
//
func (cplx *Complex) print() {
	if cplx.imag < 0 {
		fmt.Printf("Complex: Y = %d %d*j\n", cplx.real, cplx.imag)
	} else {
		fmt.Printf("Complex: Y = %d + %d*j\n", cplx.real, cplx.imag)
	}
}

// Factory
//
func NewComplex(re int, img int) Complex {
	return Complex{
		real: re,
		imag: img,
	}
}

// add a complex to an existing complex number
//
func (complex_premier *Complex) add(complex_deuxieme *Complex) {
	complex_premier.real += complex_deuxieme.real
	complex_premier.imag += complex_deuxieme.imag
}

// same as above member function of struct Complex
//
func (complex *Complex) add_new(rhs *Complex) *Complex {
	cplx := &Complex{
		real: complex.real + rhs.real,
		imag: complex.imag + rhs.imag,
	}
	return cplx
}

func main() {
	// Create a complex number and a second and add the second one to the first one
	//
	c1 := NewComplex(3, 4) // Y1 = 3 + 4*j
	c2 := NewComplex(7, 6) // Y2 = 7 + 6*j
	// print them both before and after ...
	//
	c1.print()
	c2.print()
	// add the first to the second
	//
	c2.add(&c1) // c2 is now Y2 = 10 + 10*j
	c2.print()

	// Test the add_new member function of Complex struct
	//
	c3 := c1.add_new(&c2)
	c3.print() // Y1 = (3 + 10) + (4 + 10)*j = 13 + 14*j
}
