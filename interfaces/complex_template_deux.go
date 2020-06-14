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

func info(c Complex) (string, error) {

	_, ok1 := c.real.(float64)
	_, ok2 := c.imag.(float64)

	if ok1 && ok2 {
		return fmt.Sprintf("Y = %v + %v*j", c.real, c.imag), nil
	} else {
		return "[]", errors.New("Only float64 accepted ...")
	}
}


func main() {
	c1 := NewComplex(3,4)
	c2 := NewComplex(3.3, 4.4)
	
	c_nums := []Complex{c1, c2}
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
