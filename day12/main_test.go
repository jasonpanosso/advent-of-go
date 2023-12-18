package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "gets correct answer",
			input:    input,
			expected: 7173,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := part1(tc.input); err != nil || got != tc.expected {
				t.Errorf("part1(%v) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "gets correct answer",
			input:    input,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := part2(tc.input); err != nil || got != tc.expected {
				t.Errorf("part2(%v) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}
