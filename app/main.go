package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Printf("Command entered: %s", command)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}


		switch command {
			case "exit 0":
				fmt.Printf("Exiting programm..")
				os.Exit(0)
			default:
				fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}
