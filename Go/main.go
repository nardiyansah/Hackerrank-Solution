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
