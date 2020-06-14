package main

import (
	"fmt"
)

func gen(done <-chan struct{}) (<-chan struct{}, <-chan int) {
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
	return done, intStream
}

func square(done <-chan struct{}, in <-chan int) (<-chan struct{}, <-chan int) {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case n := <-in:
				out <- n * n
			case <-done:
				return
			}

		}
	}()
	return done, out
}

func addone(done <-chan struct{}, in <-chan int) (<-chan struct{}, <-chan int) {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n := <-in:
				out <- n + 1
			case <-done:
				return
			}
		}
	}()
	return done, out
}

func finalPipe(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n := <-in:
				out <- n
			case <-done:
				//fmt.Println("done!")
				return
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)

	for n := range finalPipe(addone(square(gen(done)))) {
		fmt.Println(n)
		if n > 100 {
			break
		}
	}
}
