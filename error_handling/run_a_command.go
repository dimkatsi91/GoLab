package main

import (
	"fmt"
	"os"

	// "log"
	"bufio"
	"os/exec"
)

//
// Run | Execute a command & return its output
//
func runCmd(cmd string) string {
	output, err := exec.Command(cmd).Output()
	if err != nil {
		//log.Fatal(err)
		return "ERROR"
	}
	return string(output)
}

func main() {
	// Static command use
	//
	var whoami string = runCmd("whoami")
	fmt.Printf("System user: %s\n", whoami)
	// Get a command from stdIn and execute it
	//
	var stdInInput, command_result string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter a command %s > ", whoami)
	// start scanning ...
	scanner.Scan()
	// grab typed text ...
	stdInInput = scanner.Text()
	fmt.Printf("Scanned ... %s\n", stdInInput)
	command_result = runCmd(stdInInput)
	if command_result == "ERROR" {
		fmt.Printf("ERROR while trying to execute command: %s\n", stdInInput)
	}
	fmt.Printf("Output >> %s\n", command_result)
}
