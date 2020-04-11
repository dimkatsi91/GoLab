package main

import (
    "fmt"
)

type Person struct {
	first string
	last string
	// slice of favorite ice cream flavors
	flavors []string
}

// create a person
func (person * Person) newPerson(favorite_flavors []string , first, last string) *Person {
	return &Person {
		first : first,
		last : last,
		flavors : favorite_flavors,
	}
}

func (person * Person) echo() {
	fmt.Println(person.first + " " + person.last)
	fmt.Println(" Has also favorite flavors ... :")
	for _,flavor := range person.flavors {
		fmt.Println("\t" + flavor)
	}
}

func main() {
	// one person is :
	p1 := Person {
		first : "Dimos", 
		last : "Kacimardos", 
		flavors : []string{"peanuts", "jack daniels", "olympiakos", "white chocolate"},
	}

	p1.echo()
}

