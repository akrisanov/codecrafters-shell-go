package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// REPL loop
	for {
		// print prompt character
		fmt.Print("$ ")
		// read command and its arguments from stdin
		input, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		// parse command and arguments
		command, args := parseCommand(input)
		// handle command
		handleCommand(command, args...)
	}
}

// parseCommand splits the input into command and its arguments
func parseCommand(input string) (string, []string) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", nil
	}
	command := fields[0]
	args := make([]string, 0, len(fields)-1)
	for _, arg := range fields[1:] {
		args = append(args, strings.ReplaceAll(arg, "\"", ""))
	}
	return command, args
}

// handleCommand executes the given command with its arguments
func handleCommand(command string, args ...string) {
	if command == "" {
		return
	}
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
