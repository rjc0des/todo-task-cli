package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/store"
	"time"
)

func UpdateCommand(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: task-cli update <id> \"new text\"")
	}

	id, err := strconv.Atoi(args[0])

	if err != nil {
		return fmt.Errorf("Invalid ID\n")
	}

	newText := args[1]

	db, err := store.Load()

	if err != nil {
		return fmt.Errorf("File not loaded\n")
	}

	for i := range db.Tasks {
		if db.Tasks[i].ID == id {
			db.Tasks[i].Description = newText
			db.Tasks[i].UpdatedAt = time.Now()

			err := store.Save(db)

			if err != nil {
				return fmt.Errorf("Error saving tasks: %v\n", err)
			}

			fmt.Printf("Task Updated: %d\n", id)
			return nil
		}
	}

	return fmt.Errorf("Task not found\n")
}
