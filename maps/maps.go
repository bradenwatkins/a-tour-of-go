package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := map[string]int{}
	for _, word := range(strings.Fields(s)) {
		res[word] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
