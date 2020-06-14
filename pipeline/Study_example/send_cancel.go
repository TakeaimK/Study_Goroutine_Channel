package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newRandStream(done <-chan struct{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return randStream
}

func main() {
	done := make(chan struct{})
	defer close(done)

	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	//

	time.Sleep(1 * time.Second) // 진행중인 작업 시뮬레이션
}
