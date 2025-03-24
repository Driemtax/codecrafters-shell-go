package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		// Commands currently bultin
		commands := []string{"echo", "exit", "type", "pwd", "cd"}

		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command, args := formatInput(input)
		// fmt.Println("Command:", command)
		// fmt.Println("Args:", args, "Count:", len(args))

		_, envErr := exec.LookPath(command)
		var isEnvCommand = false
		if envErr == nil {
			isEnvCommand = true
		}

		// Evaluate the given command and args
		switch {
		// Exit the shell
		case command == "exit":
			os.Exit(0)
		case command == "echo":
			executeEcho(args)
		case command == "type":
			executeType(commands, args)
		case command == "pwd" && len(args) == 0:
			pwd()
		case command == "cd":
			err := changeDirectory(args[0])
			if err != nil {
				fmt.Printf("%s: %s: No such file or directory\n", command, args[0])
			}
		case isEnvCommand:
			executeExternal(command, args)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

// Formats the input given to command and args
func formatInput(input string) (string, []string) {
	input = strings.TrimSpace(input)
	var command string
	var args []string

	re := regexp.MustCompile(`"[^"]*"|[^\s]+`)
	matches := re.FindAllString(input, -1)

	command = matches[0]
	args = matches[1:]

	return command, args

}

// checks if a given command is a built in command in this shell
// commands is the array of builtin commands
// arg is the given command
// returns a bool if the command is builtin or not
func checkCommands(commands []string, arg string) bool {
	found := false

	for _, cmd := range commands {
		if cmd == arg {
			found = true
			break
		}
	}

	return found
}

// echo your input to the console
func executeEcho(args []string) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Print("\n")
}

// type shows how the given command would be interpreted
func executeType(commands []string, args []string) {
	envPath, envErr := exec.LookPath(args[0])
	if checkCommands(commands, args[0]) {
		fmt.Println(args[0] + " is a shell builtin")
	} else if envErr == nil {
		fmt.Println(args[0] + " is " + envPath)
	} else {
		fmt.Println(args[0] + ": not found")
	}
}

// Executes an external Programm, like e.g. git. The Programm has to be in your PATH Variable
// path: The path to the executable you wish to run
func executeExternal(command string, args []string) {
	var cmd *exec.Cmd

	if len(args) < 1 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(command, args...)
	}

	var output, err = cmd.Output()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(string(output))
}

// Prints the current working directory as an absolute path
func pwd() {
	var dir, err = os.Getwd()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dir)
	}
}

// Changes the working directory with an absolute or relative path
func changeDirectory(path string) error {
	err := os.Chdir(path)

	return err
}
