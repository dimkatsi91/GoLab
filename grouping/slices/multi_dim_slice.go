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
	for _, row := range db {
		for _, row_entry := range row {
			fmt.Printf("%s | ", row_entry)
		}
		fmt.Println(" ... ")
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Len after : ", len(db))
	fmt.Println("Capacity after : ", cap(db))
}

/* SAMPLE OUTPUT : 
	Len before :  0
	Capacity before :  0
	-----------------------------------
	Paris | France | Europe |  ... 
	Berlin | Germany | Europe |  ... 
	Chicago | America | USA |  ... 
	-----------------------------------
	Len after :  3
	Capacity after :  4
*/