//package ex1_combination_generator
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func AllwordsGenerator(filename string) <-chan string {
	WordStream := make(chan string)

	go func() {
		defer close(WordStream)

		fo, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer fo.Close()

		reader := bufio.NewReader(fo)

		for {
			line, err := reader.ReadString(' ')

			if err == io.EOF {
				WordStream <- line
				break
			}

			if err != nil {
				break
			}
			WordStream <- line
		}
		//data, _ := ioutil.ReadFile(filename)
		//temp := string(data)
		//ans := strings.Split(temp, " ")
		//
		//for _, word := range ans {
		//	WordStream <- word
		//}

	}()
	return WordStream
}

func main() {
	for w := range AllwordsGenerator("./test.txt") {
		fmt.Println(w)
	}
}
