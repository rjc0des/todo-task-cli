package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/store"
	"time"
)

func UpdateCommand(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: task-cli edit <id> \"new text\"")
		return
	}

	id, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	newText := args[1]

	db, err := store.Load()

	if err != nil {
		fmt.Println("File not loaded")
		return
	}

	for i := range db.Tasks {
		if db.Tasks[i].ID == id {
			db.Tasks[i].Description = newText
			db.Tasks[i].UpdatedAt = time.Now()

			err := store.Save(db)

			if err != nil {
				fmt.Println("Error saving tasks:", err)
				return
			}

			fmt.Printf("Task Updated: %d\n", id)
			return
		}
	}

	fmt.Println("Task not found")

}
