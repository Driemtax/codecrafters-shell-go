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

		// The next line is only for debugging
		//input := "cat '/tmp/bar/f   40' '/tmp/bar/f   45' '/tmp/bar/f   85'"

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

// Formats the input given to command and args. Extracts the command and args with a regex
func formatInput(input string) (string, []string) {
	input = strings.TrimSpace(input)
	var command string
	var args []string

	// This regex matches the following criteria:
	// 1. Everything between ""
	// 2. Everything between ''
	// 3. Everything that is not a space
	// 4. Multiple, but at least one, spaces
	re := regexp.MustCompile(`"[^"]*"|'[^']*'|([^\s'])+|(?:\s)*`) //[^\s]+|\s
	matches := re.FindAllString(input, -1)

	// Replacing all multiple occasions of spaces with a single space
	for i, match := range matches {
		if !(strings.HasPrefix(match, "'") || strings.HasPrefix(match, "\"")) {
			re := regexp.MustCompile(`\s+`)
			matches[i] = re.ReplaceAllString(match, " ")
		}
	}

	// Extracting command and possible args
	command = strings.TrimSpace(matches[0])
	if len(matches) > 1 {
		args = matches[2:] // At index 1 is always the space between the command and the args
	}

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
// Something still doesnt work right. Whitespace between to single quoted words should be printed if in input
func executeEcho(args []string) {
	for _, arg := range args {
		// Codecrafters wanted no Whitespace between to args in Single Quotes
		if strings.HasPrefix(arg, "'") {
			arg = strings.TrimPrefix(arg, "'")
			arg = strings.TrimSuffix(arg, "'")
			fmt.Print(arg)
		} else {
			fmt.Print(arg)
		}
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
	// Need to remove the args containing only a single space, because they would fuck up programs like cat
	for i, arg := range args {
		if arg == " " {
			args = append(args[:i], args[i+1:]...)
		}
	}
	// TODO: Think about if this needs to be done in every function, then abstract it to formatInput
	for i, arg := range args {
		if strings.HasPrefix(arg, "'") {
			arg = strings.TrimPrefix(arg, "'")
			args[i] = strings.TrimSuffix(arg, "'")
		}
	}
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

// Changes the working directory with an absolute, relative Path and Home directory (using ~)
func changeDirectory(path string) error {
	if strings.HasPrefix(path, "~") {
		path, _ = os.UserHomeDir()
	}
	err := os.Chdir(path)

	return err
}
