package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered channel

	go func(c chan int) {
		fmt.Println("Doing something important...")
		time.Sleep(3 * time.Second)
		c <- 1
	}(ch)

	fmt.Println("Channel is preventing the program from closing...")
	<-ch

	// ch := make(chan struct{}, 1) // FIFO
	// ch <- struct{}{}
	// fmt.Println("success pushing to channel")
	// ch <- struct{}{}
	// fmt.Println("success pushing to channel")

}
