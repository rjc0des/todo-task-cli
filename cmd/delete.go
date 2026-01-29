package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/store"
)

func DeleteCommand(args string) error {
	if len(args) <= 0 {
		return fmt.Errorf("Please provide a id\n")
	}

	id, err := strconv.Atoi(args)

	if err != nil {
		return fmt.Errorf("Please provide a vaild ID\n")
	}

	db, err := store.Load()

	if err != nil {
		return fmt.Errorf("The file is not loading\n")
	}

	deleted := false

	updated := db.Tasks[:0]

	for _, task := range db.Tasks {
		if task.ID == id {
			deleted = true
			continue
		}

		updated = append(updated, task)
	}

	db.Tasks = updated

	if !deleted {
		return fmt.Errorf("Task ID not found\n")
	}

	err = store.Save(db)

	if err != nil {
		return fmt.Errorf("Problem While saving the data\n")
	}

	fmt.Println("Deleted task:", id)
	return nil
}
