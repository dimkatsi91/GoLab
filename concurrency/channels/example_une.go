package main

import "fmt"


func calcFact(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	fact := 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	return fact
}


func main() {

	// Channels are created in Go via make, like slices and maps
	//
	channel := make(chan int)

	go func(N int) {
		for i := 3; i < N; i++ {
			channel <- calcFact(i)
		}
		// close the channel
		//
		close(channel)
	}(11)

	for fact := range channel {
		fmt.Println(fact)
	}

}