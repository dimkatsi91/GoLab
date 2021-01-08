/* 
 * Facade Design Pattern
 *
*/

package main

import (
	"fmt"
)

var temperature = 20

// Like a Lights class
//
type Lights struct {}

func (lights * Lights) on() {
	fmt.Println("Turning lights on ...")
}

func (lights * Lights) off() {
	fmt.Println("Turning lights off ...")
}

// Like a Firewall class
//
type Firewall struct {}

func (firewall * Firewall) on() {
	fmt.Println("Firewall is now UP ...")
}

func (firewall * Firewall) off() {
	fmt.Println("Firewall is now DOWN ...")
}

// Like a Heating class
//
type Heat struct {}

func (heat * Heat) adjust(value int) {
	if value >= 0 {
		fmt.Printf("Increasing temperature by %d degrees ... Current temperature is: %d Celsius\n", value, temperature+value)
	}
	fmt.Printf("Decreasing temperature by %d degrees ... Current temperature is: %d Celsius\n", value, temperature+value)
}

// Like a Facade class
//
type SmartHomeFacade struct {
	lights * Lights
	firewall * Firewall
	heat * Heat 
}

func (home * SmartHomeFacade) raise_shields() {
	home.lights.off()
	home.firewall.on()
}

// like an ... int main()
//
func main() {

	sh := new(SmartHomeFacade)
	sh.raise_shields()
	sh.heat.adjust(-20)
}
