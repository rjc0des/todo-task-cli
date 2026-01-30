package cmd

import (
	"bytes"
	"os"
	"strings"
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
