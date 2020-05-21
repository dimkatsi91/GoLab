package main

import (
	"fmt"
	"sort"
)

type Complex struct {
	real, imag int
}

// Print complex number
//
func (cplx *Complex) print() {
	if cplx.imag < 0 {
		fmt.Printf("Complex: Y = %d %d * j\n", cplx.real, cplx.imag)
	} else {
		fmt.Printf("Complex: Y = %d + %d * j\n", cplx.real, cplx.imag)
	}
}

// Sort Complex numbers By Real part
//
type ByReal []Complex

func (br ByReal) Len() int { return len(br) }
func (br ByReal) Swap(i, j int) { br[i], br[j] = br[j], br[i] }
func (br ByReal) Less(i, j int) bool { return br[i].real < br[j].real }

// Sort Complex Numbers By Imaginary part
//
type ByImag []Complex

func (bi ByImag) Len() int { return len(bi) }
func (bi ByImag) Swap(i, j int) { bi[i], bi[j] = bi[j], bi[i] }
func (bi ByImag) Less(i, j int) bool { return bi[i].imag < bi[j].imag }


func main() {
	// Create #4 Complex Numbers
	//
	c1 := Complex{1, 2}
	c2 := Complex{3, 4}
	c3 := Complex{-1, -2}
	c4 := Complex{7, 8}
	c5 := Complex{5, 2}

	/*
	var complexNumsSlice []Complex

	complexNumsSlice = append(complexNumsSlice, c1)
	complexNumsSlice = append(complexNumsSlice, c2)
	complexNumsSlice = append(complexNumsSlice, c3)
	complexNumsSlice = append(complexNumsSlice, c4)
	complexNumsSlice = append(complexNumsSlice, c5)
	*/
	complexNumsSlice := []Complex{c1, c2, c3, c4, c5} 

	fmt.Println("*** Slice of Complex Numbers Before Sorting ***")
	//fmt.Println(complexNumsSlice)
	for _, cplx_num := range complexNumsSlice {
		cplx_num.print()
	}

	////////////////////////////////// SORT BY REAL PART ///////////////////////////////////////////////
	// Sort Complex Numbers by real part OR by imaginary part | Change it if you wish to !
	//
	sort.Sort(ByReal(complexNumsSlice))
	
	fmt.Println("*** Slice of Complex Numbers After Sorting By Real Part ***")
	//fmt.Println(complexNumsSlice)
	for _, cplx_num := range complexNumsSlice {
		cplx_num.print()
	}


	////////////////////////////////// SORT BY IMAGINARY PART ///////////////////////////////////////////////
	sort.Sort(ByImag(complexNumsSlice))

	fmt.Println("*** Slice of Complex Numbers After Sorting By Imaginary Part ***")
	//fmt.Println(complexNumsSlice)
	for _, cplx_num := range complexNumsSlice {
		cplx_num.print()
	}
}

/* OUTPUT SAMPLE ::

$ go run xrcise_deux.go
*** Slice of Complex Numbers Before Sorting ***
Complex: Y = 1 + 2 * j
Complex: Y = 3 + 4 * j
Complex: Y = -1 -2 * j
Complex: Y = 7 + 8 * j
Complex: Y = 5 + 2 * j
*** Slice of Complex Numbers After Sorting By Real Part ***
Complex: Y = -1 -2 * j
Complex: Y = 1 + 2 * j
Complex: Y = 3 + 4 * j
Complex: Y = 5 + 2 * j
Complex: Y = 7 + 8 * j
*** Slice of Complex Numbers After Sorting By Imaginary Part ***
Complex: Y = -1 -2 * j
Complex: Y = 1 + 2 * j
Complex: Y = 5 + 2 * j
Complex: Y = 3 + 4 * j
Complex: Y = 7 + 8 * j


*/