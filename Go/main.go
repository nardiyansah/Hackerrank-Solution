package main

import (
	"math"
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

// https://leetcode.com/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	var subS string = s[:1]
	lenSubS := 1
	var res bool

	for lenSubS < len(s) {
		i, j := 0, lenSubS
		addJ := j
		for j <= len(s) {
			if subS == s[i:j] {
				if j == len(s) {
					return true
				}
				i = j
				j = j + addJ
			} else {
				break
			}
		}
		lenSubS += 1
		subS = s[:lenSubS]
	}
	return res
}

// https://leetcode.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	var z int = x ^ y
	var res int
	for {
		if z == 0 {
			return res
		}
		if z&1 == 1 {
			res += 1
		}
		z = z >> 1
	}
}

// https://leetcode.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	isLand, neighbour := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				isLand += 1
				up, down, left, right := i-1, i+1, j-1, j+1
				if up >= 0 && up < len(grid) && grid[up][j] == 1 {
					neighbour += 1
				}
				if down >= 0 && down < len(grid) && grid[down][j] == 1 {
					neighbour += 1
				}
				if left >= 0 && left < len(grid[i]) && grid[i][left] == 1 {
					neighbour += 1
				}
				if right >= 0 && right < len(grid[i]) && grid[i][right] == 1 {
					neighbour += 1
				}
			}
		}
	}
	var res int = (isLand * 4) - neighbour
	return res
}

// https://leetcode.com/problems/number-complement/
func findComplement(num int) int {
	pointer := 1
	for pointer <= num {
		num = num ^ pointer
		pointer <<= 1
	}
	return num
}

// https://leetcode.com/problems/license-key-formatting/
func licenseKeyFormatting(s string, k int) string {
	newS := strings.ReplaceAll(s, "-", "")
	newKey := ""
	if len(newS) <= k {
		return strings.ToUpper(newS)
	}
	if len(newS)%k == 0 {
		i := 0
		for i < len(newS) {
			j := i + k
			if j < len(s) {
				newKey += strings.ToUpper(newS[i:j])
			} else {
				newKey += strings.ToUpper(newS[i:])
			}
			i += k
			if i < len(newS) {
				newKey += "-"
			}
		}
	} else {
		i := 0
		x := 0
		for len(newS[x:])%k != 0 {
			x += 1
		}
		newKey += strings.ToUpper(newS[i:x])
		newKey += "-"
		i = x
		for i < len(newS) {
			j := i + k
			if j < len(s) {
				newKey += strings.ToUpper(newS[i:j])
			} else {
				newKey += strings.ToUpper(newS[i:])
			}
			i += k
			if i < len(newS) {
				newKey += "-"
			}
		}
	}
	return newKey
}

// https://leetcode.com/problems/construct-the-rectangle/
func constructRectangle(area int) []int {
	L, W := int(math.Sqrt(float64(area))), int(math.Sqrt(float64(area)))
	tempArea := L * W

	for tempArea != area {
		if tempArea < area {
			L += 1
		} else {
			W -= 1
		}
		tempArea = L * W
	}

	return []int{L, W}
}

// https://leetcode.com/problems/teemo-attacking/
func findPoisonedDuration(timeSeries []int, duration int) int {
	lastEndTimePoison := 0
	totalPoisonedSecond := 0

	for _, attackTime := range timeSeries {
		if attackTime == 0 {
			lastEndTimePoison = attackTime + duration - 1
			totalPoisonedSecond += duration
			continue
		}
		if attackTime <= lastEndTimePoison {
			remainTime := int(math.Abs(float64(lastEndTimePoison-attackTime))) + 1
			totalPoisonedSecond -= remainTime
			totalPoisonedSecond += duration
			lastEndTimePoison = attackTime + duration - 1
			continue
		}
		lastEndTimePoison = attackTime + duration - 1
		totalPoisonedSecond += duration
	}

	return totalPoisonedSecond
}

// https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)

	for _, val1 := range nums1 {
		for i, val2 := range nums2 {
			if val1 == val2 {
				if i+1 == len(nums2) {
					ans = append(ans, -1)
				}
				for j := i + 1; j < len(nums2); j++ {
					if j+1 == len(nums2) && nums2[j] <= val2 {
						ans = append(ans, -1)
						break
					}
					if nums2[j] > val2 {
						ans = append(ans, nums2[j])
						break
					}
				}
			}
		}
	}
	return ans
}

// https://leetcode.com/problems/convert-a-number-to-hexadecimal/
func toHex(num int) string {
	listBiner := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	listStr := []string{"f", "e", "d", "c", "b", "a", "9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}

	hex := ""

	if num == 0 {
		return "0"
	}

	for num != 0 && len(hex) < 8 {
		for i, v := range listBiner {
			if num&v == v {
				hex = listStr[i] + hex
				num >>= 4
				break
			}
		}
	}

	return hex
}

// https://leetcode.com/problems/keyboard-row/
func findWords(words []string) []string {
	keyboards := map[int]int{
		int('q'): 1,
		int('w'): 1,
		int('e'): 1,
		int('r'): 1,
		int('t'): 1,
		int('y'): 1,
		int('u'): 1,
		int('i'): 1,
		int('o'): 1,
		int('p'): 1,
		int('a'): 2,
		int('s'): 2,
		int('d'): 2,
		int('f'): 2,
		int('g'): 2,
		int('h'): 2,
		int('j'): 2,
		int('k'): 2,
		int('l'): 2,
		int('z'): 3,
		int('x'): 3,
		int('c'): 3,
		int('v'): 3,
		int('b'): 3,
		int('n'): 3,
		int('m'): 3,
	}

	wordsLower := make([]string, 0)
	for _, v := range words {
		wordsLower = append(wordsLower, strings.ToLower(v))
	}

	result := make([]string, 0)
	for idxW, w := range wordsLower {
		temp := keyboards[int(w[0])]
		for i, c := range w {
			if temp != keyboards[int(c)] {
				break
			}
			if i == len(w)-1 {
				result = append(result, words[idxW])
			}
		}
	}

	return result
}

// https://leetcode.com/problems/find-mode-in-binary-search-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var result []int
	var arrNode []*TreeNode
	var maxOccur int = 0
	var mapOccur = make(map[int]int, 0)
	arrNode = append(arrNode, root)

	for len(arrNode) > 0 {
		curr := arrNode[0]
		arrNode = arrNode[1:]
		val := curr.Val
		mapOccur[val] += 1
		if mapOccur[val] > maxOccur {
			result = []int{val}
			maxOccur = mapOccur[val]
		} else if mapOccur[val] == maxOccur {
			isContain := false
			for _, v := range result {
				if val == v {
					isContain = true
					break
				}
			}
			if !isContain {
				result = append(result, val)
			}
		}
		if curr.Right != nil {
			arrNode = append(arrNode, curr.Right)
		}
		if curr.Left != nil {
			arrNode = append(arrNode, curr.Left)
		}
	}
	return result
}

// https://leetcode.com/problems/base-7/
func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

// https://leetcode.com/problems/relative-ranks/
func findRelativeRanks(score []int) []string {
	var res []string
	gold := 0
	silver := 0
	bronze := 0

	for _, v := range score {
		if v > gold {
			bronze = silver
			silver = gold
			gold = v
			continue
		}

		if v > silver {
			bronze = silver
			silver = v
			continue
		}

		if v > bronze {
			bronze = v
			continue
		}
	}

	orgScore := make([]int, len(score))

	copy(orgScore, score)

	sort.Slice(score, func(i, j int) bool {
		return score[j] < score[i]
	})

	placement := make(map[int]string)

	for i, v := range score {
		if i == 0 || i == 1 || i == 2 {
			continue
		}
		placement[v] = strconv.Itoa(i + 1)
	}

	for _, v := range orgScore {
		if v == gold {
			res = append(res, "Gold Medal")
			continue
		}

		if v == silver {
			res = append(res, "Silver Medal")
			continue
		}

		if v == bronze {
			res = append(res, "Bronze Medal")
			continue
		}

		res = append(res, placement[v])
	}

	return res
}

// https://leetcode.com/problems/perfect-number/
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	var sum int = 1

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			anotherDivisor := num / i
			sum += i
			sum += anotherDivisor
		}
	}

	return sum == num
}

// https://leetcode.com/problems/fibonacci-number/
var memoFibonacci = make(map[int]int)

func fib(n int) int {
	if val, ok := memoFibonacci[n]; ok {
		return val
	}

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	val := fib(n-1) + fib(n-2)
	memoFibonacci[n] = val

	return val
}

// https://leetcode.com/problems/detect-capital/
func detectCapitalUse(word string) bool {
	// A-Z -> 65-90
	// a-z -> 97-122
	firstCapitalLetter := word[0] >= 65 && word[0] <= 90
	isAllCapital := true
	isAllLower := true
	isAllLoweExceptFirst := true

	for i, w := range word {
		isAllCapital = isAllCapital && w >= 65 && w <= 90
		isAllLower = isAllLower && w >= 97 && w <= 122
		if i != 0 {
			isAllLoweExceptFirst = isAllLoweExceptFirst && w >= 97 && w <= 122
		}
	}

	if isAllLower {
		return true
	}

	if isAllCapital {
		return true
	}

	if firstCapitalLetter && isAllLoweExceptFirst {
		return true
	}

	return false
}

// https://leetcode.com/problems/longest-uncommon-subsequence-i/
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return int(math.Max(float64(len(a)), float64(len(b))))
}

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return int(math.Inf(1))
	}

	minNodeLeft := getMinimumDifference(root.Left)
	minNodeRight := getMinimumDifference(root.Right)

	var left int
	if root.Left == nil {
		left = root.Val
	} else {
		left = int(math.Abs(float64(root.Val - root.Left.Val)))
	}

	var right int
	if root.Right == nil {
		right = root.Val
	} else {
		right = int(math.Abs(float64(root.Val - root.Right.Val)))
	}

	var min int
	min = int(math.Min(float64(left), float64(right)))
	min = int(math.Min(float64(min), float64(minNodeLeft)))
	min = int(math.Min(float64(min), float64(minNodeRight)))

	return min
}

// https://leetcode.com/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	r := []rune(s)

	for left := 0; left < len(r); left += 2 * k {
		for i, j := left, int(math.Min(float64(left+k-1), float64(len(r)-1))); i < j; {
			r[i], r[j] = r[j], r[i]
			i++
			j--
		}
	}

	return string(r)
}

// https://leetcode.com/problems/diameter-of-binary-tree/
var max int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDepth_v1(root)
	newMax := max
	max = 0
	return newMax
}

func maxDepth_v1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth_v1(root.Left)
	right := maxDepth_v1(root.Right)
	max = int(math.Max(float64(max), float64(left+right)))

	return int(math.Max(float64(left), float64(right))) + 1
}

// https://leetcode.com/problems/student-attendance-record-i/
func checkRecord(s string) bool {
	absenFewThan2Days := true
	lateFewThan3ConsecutiveDays := true
	countAbsent := 0
	countLate := 0

	criteria := []rune("ALP")
	ch := []rune(s)

	for _, c := range ch {
		if absenFewThan2Days {
			if c == criteria[0] {
				countAbsent += 1
				if countAbsent >= 2 {
					absenFewThan2Days = false
				}
			}
		} else {
			return false
		}

		if lateFewThan3ConsecutiveDays {
			if c == criteria[1] {
				countLate += 1
				if countLate >= 3 {
					lateFewThan3ConsecutiveDays = false
				}
			} else {
				countLate = 0
			}
		} else {
			return false
		}
	}

	return absenFewThan2Days && lateFewThan3ConsecutiveDays
}

// https://leetcode.com/problems/reverse-words-in-a-string-iii/
func reverseWords(s string) string {
	splitS := strings.Split(s, " ")
	var newSplitS []string

	for _, splt := range splitS {
		r := []rune(splt)
		if len(r) < 1 {
			continue
		}
		i, j := 0, len(r)-1
		for i < j {
			r[i], r[j] = r[j], r[i]
			i++
			j--
		}
		newSplitS = append(newSplitS, string(r))
	}
	return strings.Join(newSplitS, " ")
}

// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	max := 0

	for _, v := range root.Children {
		temp := maxDepth(v)
		if temp > max {
			max = temp
		}
	}

	return 1 + max
}

// https://leetcode.com/problems/array-partition/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var max int

	for i, v := range nums {
		if i%2 == 0 {
			max += v
		}
	}

	return max
}

//https://leetcode.com/problems/binary-tree-tilt/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDiffTilt(root) + findTilt(root.Left) + findTilt(root.Right)
}

func findDiffTilt(root *TreeNode) int {
	left := sumTreeNode(root.Left)
	right := sumTreeNode(root.Right)

	if left > right {
		return left - right
	} else {
		return right - left
	}
}

var mapSumTreeNode = make(map[*TreeNode]int)

func sumTreeNode(root *TreeNode) int {
	_, ok := mapSumTreeNode[root]
	if ok {
		return mapSumTreeNode[root]
	}

	if root == nil {
		return 0
	}
	sum := root.Val + sumTreeNode(root.Left) + sumTreeNode(root.Right)
	mapSumTreeNode[root] = sum
	return sum
}

// https://leetcode.com/problems/reshape-the-matrix/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if (r * c) != (len(mat) * len(mat[0])) {
		return mat
	}

	result := make([][]int, r)
	// from seed matrix
	//i, j := len(mat)-1, len(mat[0])-1
	// for result matrix
	x, y := 0, 0

	for _, a := range mat {
		for _, b := range a {
			if y < c {
				result[x] = append(result[x], b)
				y += 1
			} else {
				x += 1
				y = 0
				result[x] = append(result[x], b)
				y += 1
			}
		}
	}

	return result
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	if root.Val == subRoot.Val && isIdentical(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	return root.Val == subRoot.Val && isIdentical(root.Left, subRoot.Left) && isIdentical(root.Right, subRoot.Right)
}

func distributeCandies(candyType []int) int {
	doctorAdvised := len(candyType) / 2
	manyType := make(map[int]bool, 0)

	for _, v := range candyType {
		manyType[v] = true
	}

	if len(manyType) < doctorAdvised {
		return len(manyType)
	} else {
		return doctorAdvised
	}
}

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/description/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}

//https://leetcode.com/problems/n-ary-tree-postorder-traversal/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	var res []int

	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}

	res = append(res, root.Val)

	return res
}

// https://leetcode.com/problems/range-addition-ii/description/
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	maxR, maxC := 1000000, 1000000

	for _, op := range ops {
		if op[0] < maxR {
			maxR = op[0]
		}
		if op[1] < maxC {
			maxC = op[1]
		}
	}

	return maxR * maxC
}

// https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n == 1 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			break
		}
		if flowerbed[i] == 0 {
			if i == 0 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n -= 1
				}
			} else if i == len(flowerbed)-1 {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n -= 1
				}
			} else {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n -= 1
				}
			}
		}
	}

	return n == 0
}

// https://leetcode.com/problems/construct-string-from-binary-tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := strconv.Itoa(root.Val)

	if root.Left != nil || root.Right != nil {
		res += "("
		res += tree2str(root.Left)
		res += ")"
	}

	if root.Right != nil {
		res += "("
		res += tree2str(root.Right)
		res += ")"
	}

	return res
}

// https://leetcode.com/problems/merge-two-binary-trees/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var mergedTree *TreeNode

	mergedTree = combineTree(root1, root2)

	var l1, l2, r1, r2 *TreeNode
	if root1 != nil {
		l1 = root1.Left
	} else {
		l1 = nil
	}
	if root2 != nil {
		l2 = root2.Left
	} else {
		l2 = nil
	}
	if root1 != nil {
		r1 = root1.Right
	} else {
		r1 = nil
	}
	if root2 != nil {
		r2 = root2.Right
	} else {
		r2 = nil
	}

	mergedTree.Left = mergeTrees(l1, l2)
	mergedTree.Right = mergeTrees(r1, r2)

	return mergedTree
}

func combineTree(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	var mergedNode TreeNode

	if root1 != nil && root2 != nil {
		mergedNode.Val = root1.Val + root2.Val
	} else if root1 != nil {
		mergedNode.Val = root1.Val
	} else {
		mergedNode.Val = root2.Val
	}

	return &mergedNode
}
