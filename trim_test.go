package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTrim(t *testing.T) {
	foobar := "aaabbbcccaaa"
	var tests = []struct{
		pattern string
		left bool
		right bool
		expected string
		message string
	}{
		{"aaa", true, false, "bbbcccaaa", "trim left failed"},
		{"aaa",false,  true, "aaabbbccc", "trim right failed"},
		{"aaa", true, true, "bbbccc", "trim both failed"},
	}

	for _, test := range tests {
		s := trim(foobar, test.pattern, test.left, test.right)
		require.Equal(t, test.expected, s, test.message)
	}
}