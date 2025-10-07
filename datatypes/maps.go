package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var word string = ""

	for _, char := range s {
		if char != 32 {
			word += string(char)
		} else {
			m[word]++
			word = ""
		}
	}

	m[word]++

	return m
}

func main() {
	wc.Test(WordCount)
}
