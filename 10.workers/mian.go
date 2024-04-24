package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	maxWorkers := 3

	inChan := make(chan string)
	outChan := make(chan string)

	// starting our workers
	for i := 0; i < maxWorkers; i++ {
		go worker(i, inChan, outChan, &wg)
	}

	// get work done
	go func() {
		for w := range outChan {
			fmt.Println(w, "is ready!")
		}
	}()

	// work
	lowStr := []string{"cheese", "milk", "water", "pineapple", "grapes", "orange", "cake", "soda", "tomato", "cucumber"}

	// sending work
	for i := 0; i < len(lowStr); i++ {
		inChan <- lowStr[i]
	}

	// waiting for all workers to finish
	wg.Wait()
	fmt.Println("All workers finished!")
}

func worker(id int, in <-chan string, out chan<- string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case work := <-in:
			fmt.Printf("worker %d started working\n", id)
			time.Sleep(1 * time.Second)
			work = strings.ToUpper(work)
			out <- work
			fmt.Printf("worker %d is done\n", id)
		case <-time.After(4 * time.Second):
			// if no work for 2 consec seconds
			fmt.Printf("worker %d is now closing\n", id)
			wg.Done()
			return
		}
	}
}
