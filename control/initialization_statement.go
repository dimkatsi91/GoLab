package main

import "fmt"

func main() {
	if isTrue := true; isTrue == false {
		fmt.Println("This is not going to be printed cause isTrue is True and NOT False!")
	}
	
	if isTrue := false; !isTrue == true {
		fmt.Println("On the other hand this is going to be printed!")
	}
}
