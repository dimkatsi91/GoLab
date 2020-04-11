package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup


type Human interface {
	speak()
} 

type Person struct {
	first, last string
	age int
}

func (p *Person) speak() {
	fmt.Println("Hello, my name is ", p.first, p.last, " and I am ", p.age, " y.o.")
}

func saySomething(human Human) {
	human.speak()
	wait.Done()
}


///////////
// main //
/////////
func main() {

	diman := Person {
		first: "Diman",
		last: "Kaci",
		age: 28,
	}
	sumi := Person {
		first: "Michael",
		last: "Schumacher",
		age: 34,
	}
	wait.Add(2)
	go saySomething(&diman)
	go saySomething(&sumi)
	wait.Wait()
}