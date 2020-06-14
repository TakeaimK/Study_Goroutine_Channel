package main

import "fmt"

func IntegerGenerator() <-chan int {
	next := 0
	intStream := make(chan int) // IntegerGenerator 함수의 지역변수
	go func() {                 // 고루틴이자 클로저
		defer close(intStream) // 고루틴이 끝나면 intStream 채널을 닫아라.
		for {
			next++
			intStream <- next
		}
	}()
	return intStream // 생성한 intStream 채널을 리턴
}

func main() {
	for i := range IntegerGenerator() {
		fmt.Println(i)
	}
}
