package main

import "fmt"

func IntegerGenerator(done chan struct{}) <-chan int {
	next := 0
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for {
			select {
			case <-done:
				return
			default:
				next++
				intStream <- next
			}
		}

	}()
	return intStream
}

func main() {
	done := make(chan struct{})

	IntegerStream := IntegerGenerator(done)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", <-IntegerStream)
	}
	close(done)
}
