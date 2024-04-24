package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start 5 goroutines")

	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Println("Goroutine", id)
		}(i)
	}

	fmt.Println("Waiting for 1 second")
	time.Sleep(1 * time.Second)
}
