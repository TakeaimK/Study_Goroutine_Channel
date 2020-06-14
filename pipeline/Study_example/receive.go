package main

import "fmt"

func doWork(strings <-chan string) <-chan struct{} {
	completed := make(chan struct{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(completed)
		for s := range strings {
			fmt.Println(s)
		}
	}()
	return completed
}

func main() {
	doWork(nil)
	fmt.Println("Done.")
}
