package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////
type Complex struct {
	real  int
	image int
}

// A Complex Numbers Factory | Desing Pattern #1
//
func ComplexFactory(real int, image int) *Complex {
	return &Complex{
		real:  real,
		image: image,
	}
}

// A restore function for restoring A Complex Number from a Memento Object
//
func (c *Complex) restore(m *Memento) {
	fmt.Println("\nRestoring Complex Number to its previous state ...\n")
	c.real = m.real
	c.image = m.image
}

//A function to update a Complex Number
//
func (c *Complex) update(real, image int) {
	fmt.Printf("\nUpdating Complex Number as follows: real-part %d -> %d AND imaginary-part %d -> %d\n", c.real, real, c.image, image)
	c.real = real
	c.image = image
}

func (c *Complex) echo() {
	fmt.Printf("\nComplex Number: Y = %d + %d * j\n\n", c.real, c.image)
}

////////////////////////////////////////////////////////////

//// MEMENTO For Complex Numbers | Desing Pattern #2
type Memento struct {
	real  int
	image int
}

// Memento object should be created from a Complex Number object
//
func MementoFactory(c *Complex) *Memento {
	return &Memento{
		real:  c.real,
		image: c.image,
	}
}

///////////////////////////
//////// MAIN ////////////
/////////////////////////
func main() {
	cplx := ComplexFactory(3, 4)
	cplx.echo()

	// Keep a backup/snapshot of this Complex number to restore it later
	//
	snap := MementoFactory(cplx)

	// update real & imaginary parts
	//
	cplx.update(11, 33)
	cplx.echo()

	// Restore the Complex Number to its initial state
	//
	cplx.restore(snap)
	cplx.echo()
}
