package perfectNums

import "testing"

// If a number is perfect return true | else return false 
// If sum(divisors of the number) == number --> number is perfect
//
func isPerfect(num int) bool {
	var sum int // initialized to zero
	for i:=1; i<num; i++ {
		if num%i == 0 {
			sum += i 
		}
	}
	if sum==num {
		return true
	}
	return false
}

// perfectNums function return N first perfect numbers as a slice of ints []{} 
// I.E. perfectNums(5) returns the slice []{6, 28, 496, 8128, 33550336}
//
func perfectNums(N int) []int {
	var sumOfPerfectNums int // initialized to zero
	var i int = 2	// initialized to 2 since we do not consider 0 and 1 as perfect numbers
	var perNums []int
	// like a while loop
	//
	for (sumOfPerfectNums != N) {
		if isPerfect(i) == true {
			perNums = append(perNums, i)
			sumOfPerfectNums +=1
		}
		i += 1
	}
	return perNums
}


func TestPerfectnums(t *testing.T) {

	fourPerfNums := perfectNums(4)
	correctPerfNums := []int{6, 28, 496, 8128}
	falsePerNums := []int{6, 28, 333, 444}

	for i:=0; i<len(fourPerfNums); i++ {
		if correctPerfNums[i] != fourPerfNums[i] {
			t.Errorf("ERROR ... Number: %d is NOT a Perfect Number ... \n", i)
		}
	}


	for i:=0; i<len(fourPerfNums); i++ {
		if falsePerNums[i] != fourPerfNums[i] {
			t.Errorf("ERROR ... Number: %d is NOT a Perfect Number ... \n", i)
		}
	}

}

/* OUTPUT ::

$ go test -v perfectNums_test.go
=== RUN   TestPerfectnums
--- FAIL: TestPerfectnums (0.10s)
    perfectNums_test.go:56: ERROR ... Number: 2 is NOT a Perfect Number ...
    perfectNums_test.go:56: ERROR ... Number: 3 is NOT a Perfect Number ...
FAIL
FAIL    command-line-arguments  0.273s
FAIL

*/