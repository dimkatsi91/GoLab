package main

import "fmt"

func main() {

	// make(slice, length, capacity)
	var db [][]string
	fmt.Println("Len before : ", len(db))
	fmt.Println("Capacity before : ", cap(db))
	// Append a few towns to the db
	db = append(db, []string{"Paris", "France", "Europe"})
	db = append(db, []string{"Berlin", "Germany", "Europe"})
	db = append(db, []string{"Chicago", "America", "USA"})

	// fmt.Println(db)
	fmt.Println("-----------------------------------")
	for _, entry := range db {
		for _, town := range entry {
			fmt.Printf("%s | ", town)
		}
		fmt.Println(" ... ")
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Len before : ", len(db))
	fmt.Println("Capacity before : ", cap(db))
}