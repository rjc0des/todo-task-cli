package cmd

import (
	"fmt"
	"strconv"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"time"
)

func update(task *model.Task, status model.TaskStatus) {
	task.Status = status
	task.UpdatedAt = time.Now()
}

func StatusUpdate(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: <mark-todo|mark-in-progress|mark-done> <id>\n")
	}

	id, err := strconv.Atoi(args[1])

	if err != nil {
		return fmt.Errorf("Invalid ID\n")
	}

	db, err := store.Load()

	if err != nil {
		return fmt.Errorf("File not loaded\n")
	}

	for i := range db.Tasks {
		if db.Tasks[i].ID == id {

			switch args[0] {
			case "mark-todo":
				update(&db.Tasks[i], model.TaskTodo)
			case "mark-in-progress":
				update(&db.Tasks[i], model.TaskInprogress)
			case "mark-done":
				update(&db.Tasks[i], model.TaskDone)
			default:
				return fmt.Errorf("Argument not found\n")
			}

			err := store.Save(db)
			if err != nil {
				return fmt.Errorf("Error saving tasks: %v\n", err)
			}

			fmt.Printf("Task Updated: %d\n", id)
			return nil
		}
	}
	return fmt.Errorf("Task not found")
}
