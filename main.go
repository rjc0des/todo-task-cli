package main

import (
	"fmt"
	"task-cli/cmd"
)

var Version = "dev"

func main() {
	fmt.Println("Task cli Version:", Version)
	cmd.Run()
}
