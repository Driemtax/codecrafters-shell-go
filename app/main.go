package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString("\n")

	fmt.Println(command[:len(command)-1] + " command not found")

	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
