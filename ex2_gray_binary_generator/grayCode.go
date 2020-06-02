//package ex1_combination_generator
package main

import (
	"fmt"
	"strconv"
)

func convertToGray(num int) int {
	return num ^ (num >> 1)
}

func GrayBinaryGenerator(n int) <-chan []int {
	GrayCodeStream := make(chan []int)

	rng := 1

	for i := 0; i < n; i++ {
		rng *= 2
	}
	go func() {
		defer close(GrayCodeStream)
		var x string
		arr := make([]int, n)

		for i := 0; i < rng; i++ {
			x = fmt.Sprintf("%0*b", n, convertToGray(i))

			for j := 0; j < n; j++ {
				arr[j], _ = strconv.Atoi(string(x[j]))
			}
			GrayCodeStream <- arr
		}
	}()
	return GrayCodeStream
}

func main() {
	for i := range GrayBinaryGenerator(5) {
		fmt.Println(i)
	}
}
