package cmd

import (
	"os"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	store.FileName = "test-delete.json"
	defer os.Remove(store.FileName)

	now := time.Now()

	db := store.Data{
		LastId: 2,
		Tasks: []model.Task{
			{
				ID:          1,
				Description: "Test Task 1",
				Status:      model.TaskDone,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			{
				ID:          2,
				Description: "Test Task 2",
				Status:      model.TaskTodo,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}

	err := store.Save(db)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = DeleteCommand("1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, err := store.Load()

	if err != nil {
		t.Fatalf("expected to load file %v", err)
	}

	if len(data.Tasks) != 1 {
		t.Fatalf("expected 1 task remaining, got %d", len(data.Tasks))
	}

	if data.Tasks[0].ID != 2 {
		t.Errorf("expected remaining task ID=2, got %d", data.Tasks[0].ID)
	}
}

func TestDelete_TaskNotFound(t *testing.T) {
	store.FileName = "test-notfound.json"
	defer os.Remove(store.FileName)

	now := time.Now()

	db := store.Data{
		LastId: 1,
		Tasks: []model.Task{
			{
				ID:          1,
				Description: "Task Test 1",
				Status:      model.TaskTodo,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}

	err := store.Save(db)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = DeleteCommand("99")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "Task ID not found\n" {
		t.Errorf("expected 'Task ID not found', got %v", err)
	}
}

func TestDelete_InvalidID(t *testing.T) {
	err := DeleteCommand("abc")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "Please provide a vaild ID\n" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDelete_EmptyInput(t *testing.T) {
	err := DeleteCommand("")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "Please provide a id\n" {
		t.Errorf("unexpected error: %v", err)
	}
}
