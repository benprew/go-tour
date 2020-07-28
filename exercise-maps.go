package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var i int
	start := 0
	words := make(map[string]int)
	for i = 0; i < len(s); i++ {
		if s[i] == ' ' && i > start-1 {
			wrd := s[start:i]
			words[wrd]++
			for i < len(s) && s[i] == ' ' {
				i++
			}
			start = i
		}
	}
	if start < len(s) {
		wrd := s[start:i]
		words[wrd]++
	}
	return words
}

func main() {
	wc.Test(WordCount)
}
