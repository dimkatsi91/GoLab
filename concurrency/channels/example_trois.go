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
	var num int = 2
	for facts := range c {
		for _, v := range facts {
			fmt.Printf("%d! = %d\n", num, v)
			num++
		}
	}
}


func main() {
	myChannel := factsChannel(10)
	print(myChannel)
}

/* OUTPUT ::
	2! = 2
	3! = 6
	4! = 24
	5! = 120
	6! = 720
	7! = 5040
	8! = 40320
	9! = 362880
	10! = 3628800
*/