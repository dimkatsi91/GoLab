package main

import (
    "fmt"
)

func main() {
	// Sprintf example
	var eight int = 8
	line := fmt.Sprintf("\nEight Int: %d\nEight Hex: %#x\nEight Binary: %b\n", eight, eight, eight)
	// Print it out
	fmt.Println(line)
}