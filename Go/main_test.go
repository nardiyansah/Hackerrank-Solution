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
