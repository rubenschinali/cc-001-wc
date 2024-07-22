package main

import (
	"bufio"
	"coding-challenges-001-wc/counter"
	"fmt"
	"os"
	"sync"
)

func runWC(filename string) error {
	var wg sync.WaitGroup
	lineChan := make(chan int)
	wordChan := make(chan int)
	byteChan := make(chan int)
	runeChan := make(chan int)
	errChan := make(chan error, 4)

	wg.Add(4)
	go counter.Count(filename, &wg, bufio.ScanLines, lineChan, errChan)
	go counter.Count(filename, &wg, bufio.ScanWords, wordChan, errChan)
	go counter.Count(filename, &wg, bufio.ScanBytes, byteChan, errChan)
	go counter.Count(filename, &wg, bufio.ScanRunes, runeChan, errChan)

	go func() {
		wg.Wait()
		close(lineChan)
		close(wordChan)
		close(byteChan)
		close(runeChan)
		close(errChan)
	}()

	var lines, words, bytes, runes int
	var err error
	for i := 0; i < 4; i++ {
		select {
		case count := <-lineChan:
			lines = count
		case count := <-wordChan:
			words = count
		case count := <-byteChan:
			bytes = count
		case count := <-runeChan:
			runes = count
		case e := <-errChan:
			err = e
		}
	}

	if err != nil {
		return err
	}

	fmt.Printf("%d\t%d\t%d\t%d\t%s\n", lines, words, runes, bytes, filename)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wc <file>")
		os.Exit(1)
	}

	filename := os.Args[1]

	if err := runWC(filename); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
