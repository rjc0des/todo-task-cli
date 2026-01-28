package model

import "time"

type TaskStatus string

const (
	TaskTodo       TaskStatus = "todo"
	TaskInprogress TaskStatus = "in-progress"
	TaskDone       TaskStatus = "done"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (s TaskStatus) IsStatusValid() bool {
	switch s {
	case TaskDone, TaskTodo, TaskInprogress:
		return true
	default:
		return false
	}
}
