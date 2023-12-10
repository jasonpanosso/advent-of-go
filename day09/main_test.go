package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var testCases = []struct {
		name     string
		lines    []string
		expected int
	}{
		{
			name:     "gets correct answer",
			lines:    lines,
			expected: 2038472161,
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
	testCases := []struct {
		name     string
		lines    []string
		expected int
	}{
		{
			name:     "gets correct answer",
			lines:    lines,
			expected: 1091,
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
