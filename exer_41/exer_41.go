package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	sp := strings.Fields(s)

	for _, v := range sp {
		r[v]++
	}

	return r
}

func main() {
	wc.Test(WordCount)
}
