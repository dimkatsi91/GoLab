package main

import "fmt"

type Linux struct {
	distroName, codeName string
	distroVersion string
	// Linux is spinf of another Linux in most of cases
	//
	SpinOf *Linux
}

func main() {
	// Manjaro Linux inherits from Arch Linux
	//
	manjaro := &Linux {
		distroName : "Manjaro Linux",
		codeName : "Illyria",
		distroVersion : "19.0.2",
		SpinOf : &Linux {
			distroName : "Arch Linux",
			codeName : "[ NONE ]",
			distroVersion : "2020.04.01",
			SpinOf : nil,
		},
	}
	// Type Manjaro Linux name , version and its predecessor too
	//
	fmt.Println(manjaro.distroName, manjaro.distroVersion, manjaro.SpinOf.distroName)
}