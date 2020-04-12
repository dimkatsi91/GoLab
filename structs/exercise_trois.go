package main

import (
	"fmt"
)

type Vehicle struct {
	doors int
	color string
}

type Sedan struct {
	Vehicle
	luxury bool
}

type Truck struct {
	Vehicle
	fouWheel bool
}

func main() {
	// a truck
	truck := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "red",
		},
		fouWheel: false,
	}
	fmt.Println(truck)

	// a sedan
	sedan := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "beige",
		},
		luxury: true,
	}
	fmt.Println(sedan)
}
