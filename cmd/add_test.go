package cmd

import (
	"os"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"testing"
)

func TestAdd(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	AddCommand("Test Task 1")

	db, err := store.Load()

	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	if len(db.Tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(db.Tasks))
	}

	if db.Tasks[0].Description != "Test Task 1" {
		t.Errorf("Expected task description to be 'Learn Go Testing', got '%s'",
			db.Tasks[0].Description)
	}

	if db.Tasks[0].Status != model.TaskTodo {
		t.Errorf("Expected status to be 'todo', got '%s'", db.Tasks[0].Status)
	}
}
