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
