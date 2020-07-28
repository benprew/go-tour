package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, n := range words {
		counts[n]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
