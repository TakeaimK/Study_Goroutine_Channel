//package ex1_combination_generator
package main

import (
	"fmt"
)

func Move(ch chan [2]string, n int, from, to, by string) {
	if n <= 0 {
		return
	}
	Move(ch, n-1, from, by, to)
	ans := [2]string{from, to}
	ch <- ans
	Move(ch, n-1, by, to, from)
}

func HanoiGenerator(n int, from, to, by string) <-chan [2]string {
	HanoiStream := make(chan [2]string)

	go func() {
		defer close(HanoiStream)

		Move(HanoiStream, n, from, to, by)

	}()
	return HanoiStream
}

func main() {
	for move := range HanoiGenerator(3, "A", "B", "C") {
		fmt.Println(move[0], " ->", move[1])
	}
}
