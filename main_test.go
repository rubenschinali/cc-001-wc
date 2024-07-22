package main

import (
	"os"
	"testing"
)

func TestRunWC(t *testing.T) {
	filename := "test.txt"
	content := "line one\nline two\nline three"
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	if err := runWC(filename); err != nil {
		t.Fatalf("runWC returned an error: %v", err)
	}
}
