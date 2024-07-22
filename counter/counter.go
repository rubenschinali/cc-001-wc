package counter

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func Count(filename string, wg *sync.WaitGroup, split bufio.SplitFunc, resultChan chan int, errChan chan error) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		errChan <- fmt.Errorf("error opening file: %w", err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(split)
	count := 0
	for s.Scan() {
		count++
	}
	if err := s.Err(); err != nil {
		errChan <- fmt.Errorf("error reading file: %w", err)
		return
	}
	resultChan <- count
}
