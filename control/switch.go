package main

import (
	"fmt"
	"runtime"
)


func check_OS() {
	fmt.Println("This machine runs ...")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println(" on LINUX ")
	case "darwin":
		fmt.Println(" on OS X. Shit! ")
	default:
		fmt.Printf(" on ... %s\n", os)
	}
}

func main() {

	check_OS()
	
}