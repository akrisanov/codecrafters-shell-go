package main

import (
	"fmt"
)

func main() {
	// print prompt character
	fmt.Print("$ ")
	// read command from stdin
	var command string
	fmt.Scanln(&command)
	// handle command
	// for now, we treat all commands as invalid
	fmt.Printf("%s: command not found\n", command)
}
