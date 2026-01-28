package cmd

import (
	"fmt"
	"os"
)

func Run() {
	args := os.Args

	if len(args) < 1 {
		return
	}

	switch args[1] {
	case "add":
		AddCommand(args[2])
	case "update":
		UpdateCommand(args[2:])
	case "list":
		{
			status := ""
			if len(args) > 2 {
				status = args[2]
			}
			ListCommand(status)
		}
	case "delete":
		DeleteCommand(args[2])
	case "mark-todo", "mark-in-progress", "mark-done":
		StatusUpdate(args[1:])
	default:
		fmt.Println("Command not found")
	}
}
