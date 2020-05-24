package main

import (
	"fmt"
	"encoding/json"
)

type Town struct {
	Name string `json: "name"`
	Population float64 `json: "population"`
	Nation string `json: "nation"`
}


func main() {
	var jsonData = []byte (`[
				{"Name":"Athens","Population":5.15,"Nation":"Greece"},
				{"Name":"Paris","Population":6.5,"Nation":"France"},
				{"Name":"Porto","Population":3.4,"Nation":"Portugal"},
				{"Name":"Madrid","Population":7.34,"Nation":"Spain"} ]`)
	var towns []Town
	// func Unmarshal(data []byte, v interface{}) error
	// Unmarshal takes a slice of []bytes data and an inteface v and returns an error
	//
	err := json.Unmarshal(jsonData, &towns)
	if err != nil {
		fmt.Println("The outcome error is : ", err)
	}
	fmt.Printf("Type of jsondata : %T\n", jsonData)
	fmt.Printf("Type of towns : %T\n", towns)
	//fmt.Println(towns)
	for idx, town := range towns {
		fmt.Printf("Town #%d :: ", idx)
		fmt.Printf("Name: %s  |  Population: %.2f Millions  |  Country: %s \n", town.Name, town.Population, town.Nation)
	}
}
