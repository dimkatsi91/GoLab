package main

import ( 
	"fmt"
)


func main() {
	for i := 101; i < 105; i++ {
		if i % 2 == 0 {
			fmt.Printf("Loop #%d : ", i-100)
			fmt.Println("Hello World from GoLang ...")
		}
	}
}

