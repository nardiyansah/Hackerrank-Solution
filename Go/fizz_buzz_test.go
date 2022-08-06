package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
