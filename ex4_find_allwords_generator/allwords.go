//package ex1_combination_generator
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func AllwordsGenerator(filename string) <-chan string {
	WordStream := make(chan string)

	go func() {
		defer close(WordStream)

		data, _ := ioutil.ReadFile(filename)

		temp := string(data)
		ans := strings.Split(temp, " ")

		for _, word := range ans {
			WordStream <- word
		}

	}()
	return WordStream
}

func main() {
	for w := range AllwordsGenerator("./test.txt") {
		fmt.Println(w)
	}
}
