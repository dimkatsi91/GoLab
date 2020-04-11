package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wait sync.WaitGroup

func info() {
	fmt.Println("*****************************************")
	fmt.Println("Architecture:\t", runtime.GOARCH)
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Nums of CPUs:\t", runtime.NumCPU())
	fmt.Println("GoRoutines:\t", runtime.NumGoroutine())
	fmt.Println("*****************************************")
}

func calcFact(num int) int {
	if(num==0 || num==1) {
		return 1;
	}
	f := 1
	for i:=1; i<=num; i++ {
		f *= i
	}
	return f;
}

func fact(num int) {
	fmt.Println("Fact{", num, "} = ", calcFact(num))
	wait.Done()
}


func main() {
	info()
	for i:=5; i<7; i++ {
		wait.Add(1)
		go fact(i)
		wait.Wait()
		info()
	}
	fmt.Println("Factorials calculation was just executed,m right ?")
}