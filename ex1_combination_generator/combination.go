//package ex1_combination_generator
package main

import (
	"fmt"
	"math/bits"
)

//original source : https://github.com/mxschmitt/golang-combinations/blob/master/combinations.go#L32
func Combinations(ch chan []string, set []string, n int) {
	length := uint(len(set))

	if n == 0 {
		ch <- []string{}
		return
	}

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// send subset to channel
		ch <- subset
	}
}

func CombinationGenerator(fruit []string, n int) <-chan []string {
	combStream := make(chan []string)

	go func() {
		defer close(combStream)
		Combinations(combStream, fruit, n)
	}()
	return combStream
}

func main() {
	for i := range CombinationGenerator([]string{"사과", "배", "복숭아", "포도", "귤"}, 0) {
		fmt.Println(i)
	}
}
