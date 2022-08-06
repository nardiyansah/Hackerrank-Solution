package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
