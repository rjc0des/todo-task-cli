package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/store"
)

func DeleteCommand(args string) {
	if len(args) <= 0 {
		fmt.Println("Please provide a id")
		return
	}

	id, err := strconv.Atoi(args)

	if err != nil {
		fmt.Println("Please provide a vaild ID")
		return
	}

	db, err := store.Load()

	if err != nil {
		fmt.Println("The file is not loading")
		return
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
		fmt.Println("Task ID not found")
		return
	}

	err = store.Save(db)

	if err != nil {
		fmt.Println("Problem While saving the data")
		return
	}

	fmt.Println("Deleted task:", id)
}
