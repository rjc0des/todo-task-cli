package cmd

import (
	"fmt"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"time"
)

/* Add Task */
func AddCommand(args string) {
	if len(args) <= 0 {
		fmt.Println("Please add task description")
		return
	}

	db, _ := store.Load()

	db.LastId++

	now := time.Now()

	task := model.Task{
		ID:          db.LastId,
		Description: args,
		Status:      model.TaskTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	db.Tasks = append(db.Tasks, task)
	store.Save(db)

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}
