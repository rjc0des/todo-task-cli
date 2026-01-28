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

func StatusUpdate(args []string) {
	if len(args) < 1 {
		fmt.Println("ID not provided")
		return
	}

	id, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	db, err := store.Load()

	if err != nil {
		fmt.Println("File not loaded")
		return
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
				fmt.Println("Argument not found")
				return
			}

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
