package main

import (
	//"fmt"
	"testing"
)


func fact(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fact(num-1) * num
}

func TestFact(t *testing.T) {
	if fact(3) != 6 {
		t.Error("ERROR, Expected 6, GOT ", fact(3))
	}
}

// go test [-v]

/* OUTPUT

=== RUN   TestFact
--- PASS: TestFact (0.00s)
PASS

*/