package main

import (
	"log"
)

func firstUniqChar(s string) int {
	mapChar := make(map[rune]int)

	for _, v := range s {
		mapChar[v] += 1
	}

	for i, v := range s {
		if mapChar[v] == 1 {
			return i
		}
	}

	return -1
}

func main() {

	i := firstUniqChar("leetcode")
	if i != 0 {
		log.Fatalf("wrong answer: expected 0 but got %v", i)
	}
}
