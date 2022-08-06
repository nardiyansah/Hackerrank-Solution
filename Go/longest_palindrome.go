package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v] += 1
	}

	if len(m) == 1 {
		for _, v := range m {
			return v
		}
	}

	var length int

	for _, v := range m {
		if v%2 == 0 {
			length += v
		} else if v > 2 {
			length = length + v - 1
		}
	}

	for _, v := range m {
		if v == 1 {
			length += 1
			break
		}
	}

	return length
}
