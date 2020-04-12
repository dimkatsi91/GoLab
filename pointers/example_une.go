package main

import (
	"fmt"
)

type alien struct {
	species string
}

func (a *alien) info() {
	fmt.Println("Alien of Species: " + a.species)
}

func main() {
	fmt.Println("******************************************")
	narcalien := alien{
		species: "Narco Alien",
	}
	narcalien.info()
	fmt.Println("******************************************")
	var alien_ptr *alien
	fmt.Println("Printing pointer's address: ", &alien_ptr)
	fmt.Println("Printing pointer's value: ", alien_ptr)
	fmt.Println("******************************************")
	alien_ptr = &narcalien
	fmt.Println("Printing pointer's address: ", &alien_ptr)
	fmt.Println("Printing pointer's value: ", *&alien_ptr)
	alien_ptr.info()
	fmt.Println("******************************************")
}
