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
