//package ex1_combination_generator
package main

import (
	"fmt"
)

func printCode(arr []int, len int, ch chan []int) {

	snd := make([]int, len)
	for i := 0; i < len; i++ {
		snd[i] = arr[i]
	}
	ch <- snd
}

func printGray(arr []int, n, index, reverse int, ch chan []int) {
	if index == n {
		printCode(arr, n, ch)
		return
	}

	arr[index] = reverse
	printGray(arr, n, index+1, 0, ch)
	arr[index] = 1 - reverse
	printGray(arr, n, index+1, 1, ch)
}

func GrayBinaryGenerator(n int) <-chan []int {
	GrayCodeStream := make(chan []int)

	go func() {
		defer close(GrayCodeStream)

		arr := make([]int, n)
		printGray(arr, n, 0, 0, GrayCodeStream)

	}()
	return GrayCodeStream
}

func main() {
	for i := range GrayBinaryGenerator(3) {
		fmt.Println(i)
	}
}
