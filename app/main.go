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
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		switch {
		case command == "exit 0":
			os.Exit(0)
		case strings.HasPrefix(command, "echo"):
			var output string = strings.TrimPrefix(command, "echo ")
			fmt.Println(output)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
