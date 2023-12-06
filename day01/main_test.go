package main

import (
	"testing"
)

var sharedTestCases = []struct {
	name     string
	lines    []string
	expected int
}{
	{
		name:     "parses no digit",
		lines:    []string{"a"},
		expected: 0,
	},
	{
		name:     "parses single digit",
		lines:    []string{"1"},
		expected: 11,
	},
	{
		name:     "parses two digit",
		lines:    []string{"12"},
		expected: 12,
	},
	{
		name:     "correctly sets last digit when more than two digits in string",
		lines:    []string{"135"},
		expected: 15,
	},
	{
		name:     "handles garbage in string",
		lines:    []string{"jk;1ddd35asdf"},
		expected: 15,
	},
	{
		name:     "correctly sums digits across multiple strings",
		lines:    []string{"1asdf5dff", "asdf4a5dd"},
		expected: 60,
	},
}

func TestPart1(t *testing.T) {
	for _, tc := range sharedTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part1(tc.lines); got != tc.expected {
				t.Errorf("part1(%v) = %v, want %v", tc.lines, got, tc.expected)
			}
		})
	}

	var testCases = []struct {
		name     string
		lines    []string
		expected int
	}{
		{
			name:     "gets correct answer",
			lines:    lines,
			expected: 54953,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part1(tc.lines); got != tc.expected {
				t.Errorf("part1(%v) = %v, want %v", tc.lines, got, tc.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range sharedTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part2(tc.lines); got != tc.expected {
				t.Errorf("part2(%v) = %v, want %v", tc.lines, got, tc.expected)
			}
		})
	}

	testCases := []struct {
		name     string
		lines    []string
		expected int
	}{
		{
			name:     "parses single num string",
			lines:    []string{"one"},
			expected: 11,
		},
		{
			name:     "parses two num string",
			lines:    []string{"onetwo"},
			expected: 12,
		},
		{
			name:     "parses mix of num strings and nums",
			lines:    []string{"onetwo3"},
			expected: 13,
		},
		{
			name:     "parses num strings that overlap",
			lines:    []string{"3oneight"},
			expected: 38,
		},
		{
			name:     "parses num strings with garbage",
			lines:    []string{"1dfaeightffd"},
			expected: 18,
		},
		{
			name:     "gets correct answer",
			lines:    lines,
			expected: 53868,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part2(tc.lines); got != tc.expected {
				t.Errorf("part2(%v) = %v, want %v", tc.lines, got, tc.expected)
			}
		})
	}
}
