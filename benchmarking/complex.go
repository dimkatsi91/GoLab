package complex

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
func (cplx *Complex) print() string {
	if cplx.imag < 0 {
		return fmt.Sprint("Complex: Y = ", cplx.real, " ", cplx.imag, " * j\n")
	}
	return fmt.Sprint("Complex: Y = ", cplx.real, " + ", cplx.imag, " * j\n")
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

//// Run the command : 'go mod init example/complex' inside benchamrk folder
//// Then 'go test -v .'
/* OUTPUT ::
	=== RUN   TestNorm
	--- PASS: TestNorm (0.00s)
	=== RUN   TestPrint
	--- PASS: TestPrint (0.00s)
	PASS
	ok      example/complex 0.002s
*/
//// Also 'go test -bench .'
//// Also 'go test -cover'
/// OUTPUT OF COVERAGE :: 
/*
	PASS
	coverage: 66.7% of statements
	ok      example/complex 0.001s
*/
// ALso check in web browser the coverage via : 
// [1] 'go test -coverprofile a.out'
// [2] 'go tool cover -html=a.out'