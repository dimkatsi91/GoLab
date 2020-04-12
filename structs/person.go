package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

// a policeman is a person
//
type Policeman struct {
	Person
	isViolent bool
}

func NewPerson(first string, last string, age int) *Person {
	return &Person{
		first: first,
		last:  last,
		age:   age,
	}
}

func (p *Person) echo() string {
	return fmt.Sprintf("Name: %s %s \nAge: %d", p.first, p.last, p.age)
}

func main() {
	fmt.Println("***************************************************************")
	// a person
	//
	diman := Person{
		first: "Dimos",
		last:  "Katsimardos",
		age:   28,
	}
	var diman_info = diman.echo()
	fmt.Println(diman_info)

	fmt.Println("***************************************************************")

	// a policeman
	//
	pol := Policeman{
		Person: Person{
			first: "Namid",
			last:  "Khan",
			age:   33,
		},
		isViolent: true,
	}
	//fmt.Println(pol)
	var pol_info = pol.echo()
	fmt.Println(pol_info)
	fmt.Println("***************************************************************")
}
