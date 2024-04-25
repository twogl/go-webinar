package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println("Start", id)
			<-time.After(time.Duration(id) * time.Second)
			fmt.Println("Finished", id)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Program finished!")
}
