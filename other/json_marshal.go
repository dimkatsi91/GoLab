package main

import (
	"encoding/json"
	"fmt"
)

type Complex struct {
	Real  int
	Image int
}

func ComplexFactory(real, image int) *Complex {
	return &Complex{
		Real:  real,
		Image: image,
	}
}

func (c *Complex) echo() {
	fmt.Printf("Complex Number: Y = %d + %d *j\n", c.Real, c.Image)
}

func main() {
	complex_une := Complex{Real: 3, Image: 4} //ComplexFactory(3, 4)   // this is not going to run | False
	//complex_une.echo()
	complex_deux := Complex{Real: 11, Image: 33} //ComplexFactory(11, 33)

	// create a slice of two complex numbers
	complex_nums := []Complex{complex_une, complex_deux}

	// use a marshal from json to convert a Complex slice to a json object
	complex_bytes, err := json.Marshal(complex_nums)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(complex_bytes))
}
