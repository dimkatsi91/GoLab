/*
 * Checked from 'new' godoc.org :	https://pkg.go.dev/sort?tab=doc
 *
*/

package main

import (
	"fmt"
	"sort"
)

type people []string

// Because I am implementing these three methods for people 'type' => 
// people implements sort package , so I can sort.Sort() people's 
//
func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}


func main() {
	
	/////////////////////////////////////////////////////////////////
	studyGroup := people{"Zero", "John", "Al", "Jenny"}
	fmt.Printf("\n*** Before sorting ***\n")
	fmt.Println(studyGroup)

	sort.Sort(studyGroup)
	fmt.Printf("*** After sorting ***\n")
	fmt.Println(studyGroup)

	/////////////////////////////////////////////////////////////////
	brands := []string{"Xiaomi", "Samsung", "Apple", "DELL", "Lenovo"}

	fmt.Printf("\n*** Before sorting ***\n")
	fmt.Println(brands)

	// FALSE :: sort.Sort(s)
	sort.Sort(people(brands))
	
	fmt.Printf("*** After sorting ***\n")
	fmt.Println(brands)
	/////////////////////////////////////////////////////////////////
}

/*

*** Before sorting ***
[Zero John Al Jenny]
*** After sorting ***
[Al Jenny John Zero]

*** Before sorting ***
[Xiaomi Samsung Apple DELL Lenovo]
*** After sorting ***
[Apple DELL Lenovo Samsung Xiaomi]


*/