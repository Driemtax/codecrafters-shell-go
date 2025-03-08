package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
			var envPath string = checkEnvs(arg)
			if checkCommands(commands, arg) {
				fmt.Println(arg + " is a shell builtin")
			} else if envPath != "" {
				fmt.Println(arg + " is " + envPath)
			} else {
				fmt.Println(arg + ": not found")
			}
		case command == "test":
			checkEnvs(command)
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

// checks for the cmd in the PATH env
// cmd is the command that was typed in
// returns the path of the cmd if found, else just an empty string
func checkEnvs(cmd string) string {
	var path string = ""

	var paths []string = strings.Split(os.Getenv("PATH"), ":")

	for _, dir := range paths {
		found, err := findExecutable(dir, cmd)

		if err != nil {
			//fmt.Println("Error: ", err)
			continue // needed to continue instead of return if there is an error opening a single dir
		}

		if found {
			path = dir + "/" + cmd // not the cleanest way but i needed a string instance
			return path
		}

	}

	return path
}

// Finds an executable file in a given dir
// path is the given directory
// fileName is the executable you are looking for
// returns if there was an executable found with the given fileName
func findExecutable(path string, fileName string) (bool, error) {
	// Read all entries in given directory
	entries, err := os.ReadDir(path)

	if err != nil {
		return false, err
	}

	for _, entry := range entries {
		// check if is not a dir and file name is searched one
		if !entry.IsDir() && entry.Name() == fileName {
			info, err := os.Stat(filepath.Join(path, fileName))

			if err != nil {
				return false, err
			}

			// Check if file is an executable
			if (info.Mode() & 0111) != 0 {
				return true, nil
			}
		}
	}

	return false, nil
}
