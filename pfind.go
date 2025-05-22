package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/alitto/pond/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <start-directory>\n", os.Args[0])
		os.Exit(1)
	}
	root := os.Args[1]

	const numWorkers = 100
	pool := pond.NewPool(numWorkers)

	dirCh := make(chan string, 1024) // buffered channel for directories
	var wg sync.WaitGroup

	// Closure to process a directory
	var processDir func(string)
	processDir = func(dir string) {
		defer wg.Done()
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", dir, err)
			return
		}
		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				wg.Add(1)
				dirCh <- fullPath
			}
			fmt.Println(fullPath)
		}
	}

	// Start goroutine to process directories from channel
	go func() {
		for dir := range dirCh {
			pool.Submit(func() { processDir(dir) })
		}
	}()

	wg.Add(1)
  dirCh <- root

	wg.Wait()
	pool.StopAndWait()
	close(dirCh)
}
