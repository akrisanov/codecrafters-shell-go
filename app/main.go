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
			panic(err)
		}
		// parse command and arguments
		command, args := parseCommand(input)
		// handle command
		handleCommand(command, args...)
	}
}

func parseCommand(input string) (string, []string) {
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "", nil
	}
	command := fields[0]
	var args []string
	for _, arg := range fields[1:] {
		args = append(args, strings.ReplaceAll(arg, "\"", ""))
	}
	return command, args
}

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
