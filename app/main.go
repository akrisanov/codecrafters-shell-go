package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtins = map[string]struct{}{
	"echo": {},
	"exit": {},
	"type": {},
}

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

// parseCommand splits the input into command and its arguments.
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

// handleCommand executes the given command with its arguments.
func handleCommand(command string, args ...string) {
	if command == "" {
		return
	}
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "type":
		if len(args) == 0 {
			fmt.Println("type: missing argument")
			return
		}
		handleTypeCommand(args[0])
	case "pwd":
		handlePwdCommand()
	default:
		runExecutable(command, args...)
	}
}

// handleTypeCommand determines how a command would be interpreted if it were used.
func handleTypeCommand(command string) {
	if _, ok := builtins[command]; ok {
		fmt.Printf("%s is a shell builtin\n", command)
		return
	}
	if fullPath, err := exec.LookPath(command); err == nil {
		fmt.Printf("%s is %s\n", command, fullPath)
		return
	}
	fmt.Printf("%s: not found\n", command)
}

// runExecutable finds and runs the executable corresponding to the given command name and arguments.
func runExecutable(command string, args ...string) {
	fullPath, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s: command not found\n", command)
		return
	}
	cmd := exec.Command(fullPath)
	cmd.Args = append([]string{command}, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing %s: %v\n", command, err)
	}
}

// handlePwdCommand prints the current working directory.
func handlePwdCommand() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		return
	}
	fmt.Println(cwd)
}
