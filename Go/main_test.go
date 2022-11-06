package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2},
			expected: 2,
		},
		{
			input:    []int{3, 2, 1},
			expected: 1,
		},
		{
			input:    []int{2, 2, 3, 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, thirdMax(test.input))
	}
}

func TestAddStrings(t *testing.T) {
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			"11",
			"123",
			"134",
		},
		{
			"456",
			"77",
			"533",
		},
		{
			"0",
			"0",
			"0",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, addStrings(test.num1, test.num2))
	}
}

func TestCountSegments(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"Hello, my name is John",
			5,
		},
		{
			"Hello",
			1,
		},
		{
			"",
			0,
		},
		{
			"                ",
			0,
		},
		{
			", , , ,        a, eaefa",
			6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, countSegments(test.input))
	}
}

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			5,
			2,
		},
		{
			8,
			3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, arrangeCoins(test.input))
	}
}

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		inputG   []int
		inputS   []int
		expected int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2},
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{1, 2, 3},
			[]int{3},
			1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findContentChildren(test.inputG, test.inputS))
	}
}

func TestLongestPalindrome(t *testing.T) {
	test := []struct {
		input    string
		expected int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"ccc", 3},
		// {"ababababa", 9}, TODO
	}

	for _, v := range test {
		assert.Equal(t, v.expected, longestPalindrome(v.input))
	}
}

type TestForIsSubsequence struct {
	s        string
	t        string
	expected bool
}

var isSubsequenceTests = []TestForIsSubsequence{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdx", false},
	{"", "ahbgdc", true},
}

func TestIsSubsequence(t *testing.T) {
	assertions := assert.New(t)

	for _, test := range isSubsequenceTests {
		assertions.Equal(test.expected, isSubsequence(test.s, test.t))
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected []string
	}{
		{
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			input:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			input:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, fizzBuzz(test.input))
	}
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			"abab",
			true,
		},
		{
			"aba",
			false,
		},
		{
			"abcabcabcabc",
			true,
		},
		{
			"ababab",
			true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, repeatedSubstringPattern(test.input))
	}
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 4, 2},
		{3, 1, 1},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, hammingDistance(test.x, test.y))
	}
}

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			16,
		},
		{
			[][]int{
				{1},
			},
			4,
		},
		{
			[][]int{
				{1, 0},
			},
			4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, islandPerimeter(test.input))
	}
}

func TestFindComplement(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 2},
		{1, 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findComplement(test.input))
	}
}

func TestLicenseKeyFormatting(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected string
	}{
		{
			"5F3Z-2e-9-w",
			4,
			"5F3Z-2E9W",
		},
		{
			"2-5g-3-J",
			2,
			"2-5G-3J",
		},
		{
			"2-4A0r7-4k",
			4,
			"24A0-R74K",
		},
		{
			"2-4A0r7-4k",
			3,
			"24-A0R-74K",
		},
		{
			"2",
			2,
			"2",
		},
		{
			"r",
			1,
			"R",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, licenseKeyFormatting(test.s, test.k))
	}
}

func TestConstructRectangle(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{
			4,
			[]int{2, 2},
		},
		{
			37,
			[]int{37, 1},
		},
		{
			122122,
			[]int{427, 286},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, constructRectangle(test.input))
	}
}

func TestFindPoisonedDuration(t *testing.T) {
	tests := []struct {
		timeSeries []int
		duration   int
		expected   int
	}{
		{
			[]int{1, 4},
			2,
			4,
		},
		{
			[]int{1, 2},
			2,
			3,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			9,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500},
			1,
			501,
		},
		{
			[]int{0, 3, 6, 9},
			5,
			14,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findPoisonedDuration(test.timeSeries, test.duration))
	}
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, nextGreaterElement(test.nums1, test.nums2))
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		num         int
		expectedHex string
	}{
		{
			26,
			"1a",
		},
		{
			-1,
			"ffffffff",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedHex, toHex(test.num))
	}
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{"Hello", "Alaska", "Dad", "Peace"},
			[]string{"Alaska", "Dad"},
		},
		{
			[]string{"omk"},
			[]string{},
		},
		{
			[]string{"adsdf", "sfd"},
			[]string{"adsdf", "sfd"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findWords(test.input))
	}
}

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{
			100,
			"202",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.output, convertToBase7(test.input))
	}
}

func TestFindRelativeRanks(t *testing.T) {
	tests := []struct {
		input    []int
		expected []string
	}{
		{
			[]int{5, 4, 3, 2, 1},
			[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			[]int{10, 3, 8, 9, 4},
			[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findRelativeRanks(test.input))
	}
}

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{
			28,
			true,
		},
		{
			7,
			false,
		},
		{
			2016,
			false,
		},
		{
			1,
			false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, checkPerfectNumber(test.input))
	}
}

func TestFib(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			2,
			1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, fib(test.input))
	}
}

func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			"USA",
			true,
		},
		{
			"FlaG",
			false,
		},
		{
			"Leetcode",
			true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, detectCapitalUse(test.input))
	}
}

func TestFindLUSlength(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{
			"aaa",
			"aaa",
			-1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, findLUSlength(test.a, test.b))
	}
}
