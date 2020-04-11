package main

import (
	"fmt"
)

type Person struct {
	first string
	last string
	age int
}

func (p *Person) echo() string {
	return fmt.Sprintf("Name: %s %s \nAge: %d", p.first, p.last, p.age)
}

func main() {
	// a person
	//
	fmt.Println("***************************************************************")
	diman := Person {
		first : "Dimos",
		last : "Katsimardos",
		age : 28,
	} 
	var diman_info = diman.echo()
	fmt.Println(diman_info)

	// a policeman anonymous struct
	//
	fmt.Println("***************************************************************")
	p2 := struct {
		first string
		last string
		age int
	} {
		first : "Namid",
		last : "Khan",
		age : 33,
	}
	fmt.Println(p2.first, p2.last, p2.age)

	// an anonymous policeman struct
	//
	fmt.Println("***************************************************************")
	pol := struct {
		Person
		isViolent bool
	} { Person : Person {
			first : "James",
			last : "Doe",
			age : 44,
		},
		isViolent : true,
	}
	fmt.Println(pol.first, pol.last, pol.age)
	if pol.isViolent == true {
		fmt.Println("This policeman is a violent one ...")
	} else {
		fmt.Println("This policeman is a lamp oriented one ...")
	}
}