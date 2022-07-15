package main

import (
	//"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsCount := make(map[string]int)
	words := strings.Fields(s)

	for _, w := range words {
		wordsCount[w]++
	}

	return wordsCount
}

func main() {
	//wc.Test(WordCount)
}
