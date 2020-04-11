package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func (old * person) changePerson(new *person) {
	//old.name = new.name
	(*old).name = (*new).name
}

func (p * person) info() {
	fmt.Printf("Hello, my name is %s and I am %d y.o.\n", p.name, p.age)
} 

func main() {
	p1 := person{"Dimos", 28}
	p1.info()

	p2 := person{"<< Diman >>", 28}

	p1.changePerson(&p2)
	p2.info()

}