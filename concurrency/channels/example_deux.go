package main

import "fmt"

type Gender string

const (
	MALE    Gender = "Male"
	FEMALE  Gender = "Female"
	UNKNOWN Gender = "Unknown"
)

type Human struct {
	gender Gender
	name string
}

func (person *Human) info() string {
	return fmt.Sprintf("Human Name : %s\nHuman Gender : %s\n", person.gender, person.name)
}

func main() {

	// Check with buffered and unbuffered channels behavior
	personsChannel := make(chan string)
	//personsChannel := make(chan string, 3)
	completedChannel := make(chan bool)

	go func() {
		p1 := Human{MALE, "Person Une"}
		personsChannel <- p1.info()
		completedChannel <- true
	}()

	go func() {
		p2 := Human{FEMALE, "Person Deux"}
		personsChannel <- p2.info()
		completedChannel <- true
	}()

	go func() {
		p3 := Human{UNKNOWN, "Person Trois"}
		personsChannel <- p3.info()
		completedChannel <- true
	}()
	
	// Check also when not launching go routine waiting
	// for the completedChannel to be true | Uncomment to check together with (chan string, 3)
	//
	go func() {
		<-completedChannel
		<-completedChannel
		<-completedChannel
		close(personsChannel)
	}()

	for personInfo := range personsChannel {
		fmt.Println(personInfo)
	}
	
}