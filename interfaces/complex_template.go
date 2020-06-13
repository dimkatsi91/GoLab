package main


import (
	"fmt"
	"errors"
)

// define complex number struct
// real and imginary part can be of any type, 
// making use of Go's empty interface feature
//
type Complex struct {
	real interface{}
	imag interface{}
}

// Factory
//
func NewComplex(re interface{}, img interface{}) Complex {
	return Complex{
		real: re,
		imag: img,
	}
}

func info(c interface{}) (string, error) {
	switch cplx := c.(type) {
	case Complex:
		return fmt.Sprintf("Y = %v + %v*j", cplx.real, cplx.imag), nil
	default:
		return "[]", errors.New("Only integer accepted ...")
	}
}


func main() {
	c1 := NewComplex(3,4)
	c2 := NewComplex(3.3, 4.4)
	// Use c3 Complex's address to store real and imaginary part
	// of c3 complex number
	// 
	c3 := &Complex{11, 33.33,}
	c4 := Complex{44.44, 8}
	
	// Dereference c3 to access its values, i.e. real + imaginary part
	// Same as c3:=Complex{11,33.33} ; c_nums:=[]Complex{c1,c2,c3,c4}
	//
	c_nums := []Complex{c1, c2, *c3, c4}
	// Both c_info and error need to be initialized
	//
	err := errors.New("Initialized error")
	var c_info string

	for _, c_num := range c_nums {
		c_info, err = info(c_num)
		if err != nil {
			fmt.Println("Error :: ", err)
		} else {
			fmt.Println(c_info)
		}
	}
}
