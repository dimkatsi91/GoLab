package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func try_open_the_file(filename string) (string, bool) {
	file, err := os.Open(filename)
	if err != nil {
		//fmt.Println("Error: ", err)
		log.Fatal("<<< Error: ", err, " >>>")
		return "ERROR", true
	}

	fmt.Println("Closing the file: ", filename)
	defer file.Close()
	fmt.Println("Closed file : ", filename)

	byteSlice, err := ioutil.ReadAll(file)

	return string(byteSlice), false
}

func main() {
	fileOutput, hasError := try_open_the_file("textFile.txt") // ("unknownFile.txt") to check an error
	if hasError == true {
		fmt.Println("Oops an error occurred ... ")
	} else {
		fmt.Printf("\n%s\n", fileOutput)
	}
}
