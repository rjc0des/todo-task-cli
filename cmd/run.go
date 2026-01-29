package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
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
	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func printHelp() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	fmt.Println("Tasker Cli")
	fmt.Print("An cli app for todo task without internet or database\n\n")
	fmt.Fprintln(w, "Command\tUsage\tdescription")
	fmt.Fprintln(w, "-------\t-----\t-----------")
	fmt.Fprintln(w, "add\tadd <task description>\tAdd task")
	fmt.Fprintln(w, "list\tlist [todo|in-progress|done]\tList the task")
	fmt.Fprintln(w, "update\tupdate <task id> <task description>\tUpdate the task description")
	fmt.Fprintln(w, "delete\tdelete <task id>\tDelete a task using task id")
	fmt.Fprintln(w, "mark-todo\tmark-todo <task id>\tUpdate the task status to todo")
	fmt.Fprintln(w, "mark-in-progress\tmark-in-progress <task id>\tUpdate the task status to in-progress")
	fmt.Fprintln(w, "mark-done\tmark-done <task id>\tUpdate the task status to done")

	w.Flush()
}
