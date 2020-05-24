package main


import (
	"fmt"
)

func rFact(num int) int {
	if num==0 || num==1 {
		return num
	}
	return rFact(num-1) * num	// Example num=5: fact(4)*5 * fact(3)*4 * fact(2)*3 * fact(1)*2
								//                <--------- this order calculation recursively
}

// This function returns a channel with factorials from 2...N as a slice of ints []{}
//
func factsChannel(N int) <-chan []int {
	var i int
	var factNums []int
	// Bidirectional channel holding a slice of ints, which holds factorials from 2...N
	//
	sliceChannel := make(chan []int, N)
	for i=2; i<=N; i++ {
		factNums = append(factNums, rFact(i))
	}
	sliceChannel <- factNums
	close(sliceChannel)
	return sliceChannel
}

// This function accepts a channel of type slice of ints, 
// pulls values from it and prints them out
//
func print(c <- chan []int) {
	for fact := range c {
		fmt.Println(fact)
	}
}


func main() {
	myChannel := factsChannel(10)
	print(myChannel)
}

/* OUTPUT ::
				[2 6 24 120 720 5040 40320 362880 3628800]
*/