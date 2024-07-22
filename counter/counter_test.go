package counter

import (
	"bufio"
	"os"
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	filename := "test.txt"

	// Create a test file
	content := "line one\nline two\nline three"
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	var wg sync.WaitGroup
	resultChan := make(chan int, 1)
	errChan := make(chan error, 1)

	wg.Add(1)
	go Count(filename, &wg, bufio.ScanLines, resultChan, errChan)

	wg.Wait()
	close(resultChan)
	close(errChan)

	select {
	case err := <-errChan:
		t.Fatalf("Count returned an error: %v", err)
	case count := <-resultChan:
		if count != 3 {
			t.Errorf("expected 3 lines, got %d", count)
		}
	}
}
