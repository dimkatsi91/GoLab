package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func echo() {
	fmt.Println("Function No.1 runs ...")
	wait.Done()
}

func newEcho(name string) {
	fmt.Println("Function ", name, " runs ...")
	wait.Done()
}

func main() {

	wait.Add(1)
	go echo()
	wait.Wait()
	////////////////////
	wait.Add(1)
	go newEcho("No.2")
	wait.Wait()
}