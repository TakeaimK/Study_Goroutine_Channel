package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan int, strings <-chan string) <-chan struct{} {
	completed := make(chan struct{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(completed)
		for {
			select {
			case s := <-strings:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return completed
}

func main() {
	done := make(chan int)
	completed := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-completed

	fmt.Println("Done.")
}
