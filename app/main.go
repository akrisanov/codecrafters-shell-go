package main

import (
	"fmt"
)

func main() {
	var command string
	// REPL loop
	for {
		// print prompt character
		fmt.Print("$ ")
		// read command from stdin
		fmt.Scanln(&command)
		// handle command
		// for now, we treat all commands as invalid
		fmt.Printf("%s: command not found\n", command)
	}
}
