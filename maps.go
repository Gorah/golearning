package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	arr := strings.Fields(s)
	countMap := make(map[string]int)
	for _, el := range arr {
		countMap[el] += 1
	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
