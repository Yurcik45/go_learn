package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	splitted := strings.Fields(s)
	counted := make(map[string]int)
	
	for a := 0; a < len(splitted); a++ {
		count := 0
		for b := 0; b < len(splitted); b++ {
			if splitted[a] == splitted[b] { count++ }
		}
		counted[splitted[a]] = count
	}
	return counted
}

func main() {
	wc.Test(WordCount)
}

