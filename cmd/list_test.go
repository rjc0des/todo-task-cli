package cmd

import (
	"bytes"
	"os"
	"strings"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"testing"
)

func TestList_AllTask(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	err := AddCommand("Test Task 1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = AddCommand("Test Task 2")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = ListCommand("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)

	output := buf.String()

	if !strings.Contains(output, "Test Task 1") {
		t.Errorf("expected output to contain Task One")
	}

	if !strings.Contains(output, "Test Task 2") {
		t.Errorf("expected output to contain Task Two")
	}

}

func TestListCommand_InvalidStatus(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	err := AddCommand("Test Task 1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ListCommand("wrongstatus")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "status invalid\n" {
		t.Errorf("expected 'status invalid', got %v", err)
	}
}

func TestList_FilteredStatus(t *testing.T) {
	testFile := "test-db.json"
	store.FileName = testFile

	defer os.Remove(testFile)

	err := AddCommand("Test Task 1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err = ListCommand("todo")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)

	output := buf.String()

	if !strings.Contains(output, "todo") {
		t.Errorf("did NOT expect other than todo in output")
	}
}

func TestListCommand_NoTasks(t *testing.T) {
	store.FileName = "test-empty.json"
	defer os.Remove(store.FileName)

	// Save empty DB
	_ = store.Save(store.Data{LastId: 0, Tasks: []model.Task{}})

	err := ListCommand("")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "No tasks found\n" {
		t.Errorf("expected 'No tasks found', got %v", err)
	}
}
