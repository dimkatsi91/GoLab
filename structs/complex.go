package main

import (
	"fmt"
	"math"
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
		fmt.Printf("Complex: Y = %d %d * j\n", cplx.real, cplx.imag)
	} else {
		fmt.Printf("Complex: Y = %d + %d * j\n", cplx.real, cplx.imag)
	}
}

// Factory
//
func NewComplex(re int, img int) Complex {
	return Complex { 
		real : re,
		imag : img,
	}
}

// |y| = |re|^2 + |y|^2
//
func (cplx *Complex) norm() float64 {
	// explicit conversions ... between types ...
	var x float64 = float64(cplx.real)
	var y float64 = float64(cplx.imag)
	var z = math.Sqrt( x*x + y*y )
	//fmt.Printf("|%d + %d*j| = %f \n", cplx.real, cplx.imag, z)

	return z
}

// Complex conjugate
//
func conj(cplx *Complex) Complex {
	return Complex {
		real : cplx.real,
		imag : -cplx.imag,
	}
}

func main() {
	// 1st way ::
	cplx_i := &Complex{3, 4}
	cplx_i.print()
	var cplx_i_norm float64 = cplx_i.norm()
	fmt.Printf("|%d + %d*j| = %f \n", cplx_i.real, cplx_i.imag, cplx_i_norm)

	// 2nd ::
	cplx_ii := NewComplex(10, 20);
	cplx_ii.print()

	// Just conjugate the 2nd complex via a func(*obj) ::
	cplx_iii := conj(&cplx_ii) // &Complex{33, 44}
	cplx_iii.print()

	// 3rd way ::
	cplx_iv := new(Complex)
	cplx_iv.real = 1
	cplx_iv.imag = 1
	cplx_iv.print()

	// 4th way ::
	cplx_v := &Complex {
		real : 300,
		imag : 400,
	}
	cplx_v.print()
	var cplx_v_norm float64 = cplx_v.norm()
	fmt.Printf("|%d + %d*j| = %f \n", cplx_v.real, cplx_v.imag, cplx_v_norm)
}
