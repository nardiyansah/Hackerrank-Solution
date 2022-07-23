package main

func findTheDifference(s string, t string) byte {
	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for _, v := range s {
		mapS[v] += 1
	}

	for _, v := range t {
		mapT[v] += 1
	}

	var c byte

	for k := range mapT {
		if mapS[k] != mapT[k] {
			c = byte(k)
			break
		}
	}

	return c
}
