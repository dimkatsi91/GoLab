package complex

import (
	"fmt"
	"testing"
	//"complex"
)

////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////

func TestNorm(t *testing.T) {
	c1 := &Complex{3, 4}

	if(c1.norm() != 5) {
		t.Error("Error, Expected 5, got ", c1.norm())
	}
}


func TestPrint(t *testing.T) {
	c1 := &Complex{3,4}

	if(c1.print() != "Complex: Y = 3 + 4 * j\n") {
		t.Error("Ooops failure ...")
	}
}


func BenchmarkPrint(b *testing.B) {
	c1 := &Complex{3, 4}

	for i:=0 ; i<b.N; i++ {
		c1.print()
	}
}

func BenchmarkNorm(b *testing.B) {
	c1 := Complex{3,4}

	for i:=0; i<b.N; i++ {
		fmt.Println(c1.norm())
	}
}


func ExamplePrint() {
	c1 := &Complex{3, 4}
	fmt.Println(c1.print())
}


func ExampleNorm() {
	c1 := &Complex{3, 4}
	fmt.Println(c1.norm())
}

// go test
// go test -bench .
// go test -cover
// go test -coverprofile a.out
// go tool cover -html=a.out