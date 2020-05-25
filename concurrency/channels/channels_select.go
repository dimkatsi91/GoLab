package main

import (
	"fmt"
	"time"
)

func rFact(num int) int {
	if num==0 || num==1 {
		return num
	}
	return rFact(num-1) * num
}

func sFact(num int) string {
	return fmt.Sprintf("%d! = %d", num, rFact(num))
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			for i:=1; i<5; i++ {
				c1 <- sFact(i)
				// every half second
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	go func() {
		for {
			for i:=5; i<11; i++ {
				c2 <- sFact(i)
				// every two seconds
				time.Sleep(time.Second * 2)
			}
		}
	}()

	go func() {
		for {
			select {
			case fact1 := <- c1:
				fmt.Println("Pulled a factorial from channel-1: ", fact1)
			case fact2 := <- c2:
				fmt.Println("Pulled a factorial from channel-2: ", fact2)
			case <- time.After(time.Millisecond * 500):
				fmt.Println("timeout")
			}
		}
	}()

	time.Sleep(time.Second * 5)
}

