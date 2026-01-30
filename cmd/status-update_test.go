package cmd

import (
	"os"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"testing"
	"time"
)

func TestStatusUpdate(t *testing.T) {
	store.FileName = "test-db.json"
	defer os.Remove(store.FileName)

	now := time.Now()

	db := store.Data{
		LastId: 1,
		Tasks: []model.Task{
			{
				ID:          1,
				Description: "Test Task 1",
				Status:      model.TaskTodo,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}

	_ = store.Save(db)

	err := StatusUpdate([]string{"mark-done", "1"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, _ := store.Load()

	if data.Tasks[0].Status != model.TaskDone {
		t.Errorf("expected status 'done', got %s", data.Tasks[0].Status)
	}
}

func TestStatusUpdate_TaskNotFound(t *testing.T) {
	store.FileName = "test-status-notfound.json"
	defer os.Remove(store.FileName)

	_ = store.Save(store.Data{})

	err := StatusUpdate([]string{"mark-done", "99"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "Task not found\n" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStatusUpdate_InvalidCommand(t *testing.T) {
	store.FileName = "test-status-invalid.json"
	defer os.Remove(store.FileName)

	_ = store.Save(store.Data{
		LastId: 1,
		Tasks:  []model.Task{{ID: 1}},
	})

	err := StatusUpdate([]string{"mark-wrong", "1"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestStatusUpdate_InvalidArgs(t *testing.T) {
	err := StatusUpdate([]string{"mark-done"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
