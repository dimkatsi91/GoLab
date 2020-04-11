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

	// another person is:
	p2 := Person {
		first : "Dimosthenis", 
		last : "Katsimardos", 
		flavors : []string{"Johny Black", "Panetolikos", "dark chocolate"},
	}

	// A map between a person first name - and its whole entity
	personsMap := map[string]Person{
		p1.first : p1 , 
		p2.first : p2,
	}

	fmt.Println("*********************************")
	for _, person := range personsMap {
		person.echo()
		fmt.Println("*********************************")
	}	
}

