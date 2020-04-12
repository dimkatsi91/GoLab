/*
	Exemple Une : arguments & iotas in GoLang
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

// a few constant exit codes | Style of Go-Lang using iotas
//
const (
	NO_ERROR    = iota     // 0
	INPUT_ERROR = iota + 1 //  1
	OTHER_ERROR = iota + 2 //   2
)

// capitalize a string
//
func capitalize(str string) string {
	return strings.ToUpper(str)
}

func main() {
	// check provided arguments length
	//
	args_len := len(os.Args)
	// sanity check ...
	//
	if args_len < 2 {
		fmt.Println("Please specify at least one argument and try again!")
		os.Exit(INPUT_ERROR)
	}
	// print out all other provided arguments capitalized ...
	//
	for i := 1; i < args_len; i++ {
		fmt.Printf("%s capitalized --> %s\n", os.Args[i], capitalize(os.Args[i]))
	}
}

/* OUTPUT Example ::

user_une@WINDOWS_PC MINGW64 ~/MyDocuments/golab/other (master)
$ go run iotas_example.go iota beta gamma delta ...
iota capitalized --> IOTA
beta capitalized --> BETA
gamma capitalized --> GAMMA
delta capitalized --> DELTA
... capitalized --> ...

*/
