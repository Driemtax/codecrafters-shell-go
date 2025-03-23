package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		// Commands currently bultin
		commands := [4]string{"echo", "exit", "type", "pwd"}

		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command, args := formatInput(input)

		_, envErr := exec.LookPath(command)
		var isEnvCommand = false
		if envErr == nil {
			isEnvCommand = true
		}

		// Evaluate the given command and args
		switch {
		// Exit the shell
		case command == "exit" && args == "0":
			os.Exit(0)
		// echo your input to the console
		case command == "echo":
			fmt.Println(args)
		// type shows how the given command would be interpreted
		case command == "type":
			envPath, envErr := exec.LookPath(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error in type:", envErr)
			}
			if checkCommands(commands, args) {
				fmt.Println(args + " is a shell builtin")
			} else if envErr == nil {
				fmt.Println(args + " is " + envPath)
			} else {
				fmt.Println(args + ": not found")
			}
		case command == "pwd" && args == "":
			pwd()
		case isEnvCommand:
			executeExternal(command, args)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

// Formats the input given to command and args
func formatInput(input string) (string, string) {
	input = strings.TrimSpace(input)
	var newInput = strings.SplitN(input, " ", 2)
	var command = newInput[0]
	var args = ""
	if len(newInput) > 1 {
		args = newInput[1]
	}

	return command, args

}

// checks if a given command is a built in command in this shell
// commands is the array of builtin commands
// arg is the given command
// returns a bool if the command is builtin or not
func checkCommands(commands [4]string, arg string) bool {
	found := false

	for _, cmd := range commands {
		if cmd == arg {
			found = true
			break
		}
	}

	return found
}

// Executes an external Programm, like e.g. git. The Programm has to be in your PATH Variable
// path: The path to the executable you wish to run
func executeExternal(command string, args string) {
	var seperated = strings.Split(args, " ")

	var cmd = exec.Command(command, seperated...)

	if len(seperated) == 0 {
		cmd = exec.Command(command)
	}

	var output, err = cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(output))
}

// Prints the current working directory
func pwd() {
	var dir, err = os.Getwd()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}
