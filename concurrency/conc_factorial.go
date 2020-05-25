package main

import (
	"fmt"
	"time"
	"math/rand"
)

func rFact(num int) int {
	if num==0 || num==1 {
		return num
	}
	return rFact(num-1) * num
}

func printFact(num int) {
	fmt.Println(num, "! = ", rFact(num))
}

func main() {
	var N int = 10
	for i:=0; i<N; i++ {
		go printFact(i)
		// If this sleep is removed the above function will not conclude
		// immediately, since main will not wait for printFact() to be finished
		// but instead will continue and this printFact() def will be concluded
		// after main has finished, so its output will never be printed-out
		rt := time.Duration(rand.Intn(200))
		time.Sleep(time.Millisecond * rt) 
	}
	// Also move the above sleep statements here outside our loop and watch
	// what is going to happen. The printed factorials will be printed-out with
	// no corect order from 0-10 but randomly, since they are concurrently calculated
	// and their values are printed-out concurrently
	//rt := time.Duration(rand.Intn(200))
	//time.Sleep(time.Millisecond * rt) 
}
