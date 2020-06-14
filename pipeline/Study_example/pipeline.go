package main

import "fmt"

func gen() <-chan int {
	next := 0
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for {
			next++
			intStream <- next
		}
	}()
	return intStream
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func addone(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

func main() {
	n := 0
	for s := range square(square(addone(square(gen())))) {
		if n == 5 {
			break
		}
		fmt.Println(s)
		n++
	}
}
