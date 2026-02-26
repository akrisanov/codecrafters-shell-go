package main

import (
	"fmt"
	"os"
)

func main() {
	var command string
	// REPL loop
	for {
		// print prompt character
		fmt.Print("$ ")
		// read command from stdin
		fmt.Scanln(&command)
		handleCommand(command)
	}
}

func handleCommand(command string) {
	switch command {
	case "exit":
		os.Exit(0)
	default:
		// handle command
		// for now, we treat all commands as invalid
		fmt.Printf("%s: command not found\n", command)
	}
}
