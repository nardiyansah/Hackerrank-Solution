package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert := assert.New(t)

	for _, test := range isSubsequenceTests {
		assert.Equal(test.expected, isSubsequence(test.s, test.t))
	}
}
