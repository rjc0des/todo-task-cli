package cmd

import (
	"fmt"
	"os"
)

func Run() {
	args := os.Args

	if len(args) < 1 {
		// panic("Hello")
		return
	}

	switch args[1] {
	case "add":
		AddCommand(args[2])
	default:
		fmt.Println("Command not found")
	}
}
