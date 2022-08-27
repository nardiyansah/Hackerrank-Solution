package main

import (
	"sort"
	"strconv"
	"strings"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)

	// remove duplicate
	unique := make(map[int]bool)
	newNums := make([]int, 0)
	for _, v := range nums {
		if unique[v] {
			continue
		}
		newNums = append(newNums, v)
		unique[v] = true
	}

	if len(newNums) < 3 {
		temp := 0
		for _, v := range newNums {
			if v > temp {
				temp = v
			}
		}
		return temp
	}

	return newNums[len(newNums)-3]
}

func addStrings(num1 string, num2 string) string {
	p1, p2 := len(num1)-1, len(num2)-1
	res := ""
	carry := ""

	for {
		var carryAscii int
		if len(carry) > 0 {
			carryAscii = int(carry[0])
			carry = ""
		}
		var x1 int
		if p1 >= 0 {
			x1 = int(num1[p1])
			p1 -= 1
		}
		var x2 int
		if p2 >= 0 {
			x2 = int(num2[p2])
			p2 -= 1
		}
		y := (carryAscii + x1 + x2) % 48
		temp := strconv.Itoa(y)
		if len(temp) > 1 {
			carry = string(temp[0])
			temp = string(temp[1])
		}
		res = temp + res
		if p1 < 0 && p2 < 0 && len(carry) <= 0 {
			return res
		}
	}
}

func countSegments(s string) int {
	sliceS := strings.Split(s, " ")
	var counter int

	for _, v := range sliceS {
		if v != "" && v != " " {
			counter += 1
		}
	}
	return counter
}

func arrangeCoins(n int) int {
	var stairs int
	counter := 1

	for n > 0 {
		if n >= counter {
			stairs += 1
			n = n - counter
			counter += 1
		} else {
			n = -1
			break
		}
	}

	return stairs
}

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	idxG, idxS := 0, 0
	var res int

	for {
		if idxG > (len(g)-1) || idxS > (len(s)-1) {
			break
		}

		if g[idxG] <= s[idxS] {
			res += 1
			idxG += 1
			idxS += 1
		} else {
			idxS += 1
		}
	}

	return res
}

func canConstruct(ransomNote string, magazine string) bool {
	charRansom := []rune(ransomNote)
	charMagazine := []rune(magazine)

	mapRansom := make(map[rune]int16)
	mapMagazine := make(map[rune]int16)

	for i := 0; i < len(charRansom); i++ {
		mapRansom[charRansom[i]] += 1
	}

	for i := 0; i < len(charMagazine); i++ {
		mapMagazine[charMagazine[i]] += 1
	}

	for k := range mapRansom {
		if mapRansom[k] > mapMagazine[k] {
			return false
		}
	}

	return true
}

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

func isSubsequence(s string, t string) bool {
	idxs, idxt := 0, 0

	if len(s) == 0 {
		return true
	}

	for {
		// condition to out from loop and string not subsequence
		if idxs >= len(s) || idxt >= len(t) {
			break
		}

		if s[idxs] == t[idxt] {
			if idxs == len(s)-1 {
				return true
			}
			idxs += 1
		}
		idxt += 1
	}

	return false
}

func guess(lower int) int {
	return 0
}

func guessNumber(n int) int {
	lower, upper := 0, n

	if guess(lower) == 0 {
		return lower
	}

	if guess(upper) == 0 {
		return upper
	}

	for {
		guessedNumber := (lower + upper) / 2
		resultGuess := guess(guessedNumber)
		if resultGuess == 0 {
			return guessedNumber
		} else if resultGuess == -1 {
			upper = guessedNumber
		} else {
			lower = guessedNumber
		}

	}
}

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

func fizzBuzz(n int) []string {
	var out []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out = append(out, "FizzBuzz")
		} else if i%3 == 0 {
			out = append(out, "Fizz")
		} else if i%5 == 0 {
			out = append(out, "Buzz")
		} else {
			out = append(out, strconv.FormatInt(int64(i), 10))
		}
	}

	return out
}
