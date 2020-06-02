package main

import "fmt"

func IntegerGenerator(n int) <-chan int {
	next := 0
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for next < n {
			next++
			intStream <- next
		}
	}()
	return intStream
}

func main() {
	for i := range IntegerGenerator(10) {
		fmt.Println(i)
	}
}
