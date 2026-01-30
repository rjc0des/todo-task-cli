package cmd

import (
	"os"
	"task-cli/internal/store"
	"testing"
)

func TestUpdate(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	AddCommand("Test Task 1")

	err := UpdateCommand([]string{"1", "Testing task 1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	db, err := store.Load()

	if err != nil {
		t.Fatalf("Error loading file %v", err)
	}

	if db.Tasks[0].Description != "Testing task 1" {
		t.Fatalf("expected updated description, got %s", db.Tasks[0].Description)
	}
}

func TestUpdate_TaskNotFound(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	AddCommand("Test Task 1")

	_, err := store.Load()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = UpdateCommand([]string{"99", "Doesn't matter"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() == "Task not found" {
		t.Errorf("expected 'Task not found', got %v", err)
	}
}

func TestUpdate_InvalidArgs(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	AddCommand("Test Task 1")

	_, err := store.Load()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = UpdateCommand([]string{"1"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
