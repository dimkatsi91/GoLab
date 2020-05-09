package main

import (
	"fmt"
)

// Human Interface
//
// type Human Interface{} 


// Friend implements Human and has a street where lives et a phone-number
//
type Friend struct {
	street, phoneNumber string
}

func create(street, phoneNumber string) Friend {
	return Friend{
		street : street,
		phoneNumber : phoneNumber,
	}
}

// A very simple Dictionary with key values (street) -> (Friend)
//
type Dictio struct {
	entry map[string]Friend
}

func (dictio *Dictio) printEntries() {
	entries := len(dictio.entry)
	fmt.Printf("Detected %d friends in this Dictionary.\n\n", entries)
	fmt.Println("STREET \t\t\t NAME \t\t\t PHONE-NUMBER \n")
	for key, value := range dictio.entry {
		fmt.Printf("%s \t %s \t%s\n", key, value.street, value.phoneNumber)
	}
}


func main() {
	// create two friend
	//
	giovanni := create("Giovanni Silva", "210-102030")
	// fmt.Println(giovanni)
	rivaldo := create("Rivaldo Oliveira", "210-506070")
	// fmt.Println(rivaldo)

	// Creer une notebook avec deux amis braziliens pour moi !!!
	//
	notebook := &Dictio{
		entry : map[string]Friend{
			"Brazilien Avenue 920" : giovanni,
			"Brazilien Avenue 101" : rivaldo,
		},
	}
	// fmt.Println(notebook)
	notebook.printEntries()
}
