package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		// Commands currently bultin
		commands := [3]string{"echo", "exit", "type"}

		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		// Evaluate the given command and args
		switch {
		// Exit the shell
		case command == "exit 0":
			os.Exit(0)
		// echo your input to the console
		case strings.HasPrefix(command, "echo"):
			var output string = strings.TrimPrefix(command, "echo ")
			fmt.Println(output)
		// type shows how the given command would be interpreted
		case strings.HasPrefix(command, "type"):
			var arg string = strings.TrimPrefix(command, "type ")
			if checkCommands(commands, arg) {
				fmt.Println(arg + " is a shell builtin")
			} else {
				fmt.Println(arg + ": not found")
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

// checks if a given command is a built in command in this shell
// commands is the array of builtin commands
// arg is the given command
// returns a bool if the command is builtin or not
func checkCommands(commands [3]string, arg string) bool {
	found := false

	for _, cmd := range commands {
		if cmd == arg {
			found = true
			break
		}
	}

	return found
}
