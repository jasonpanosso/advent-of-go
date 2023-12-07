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
			expected: 2006,
		},
	}

	games := splitStringsByGame(lines)
	gamesWithTurns := splitGamesByTurns(games)
	gamesWithTurnsWithMoves := splitTurnsByMove(gamesWithTurns)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part1(gamesWithTurnsWithMoves); got != tc.expected {
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
			expected: 84911,
		},
	}

	games := splitStringsByGame(lines)
	gamesWithTurns := splitGamesByTurns(games)
	gamesWithTurnsWithMoves := splitTurnsByMove(gamesWithTurns)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := part2(gamesWithTurnsWithMoves); got != tc.expected {
				t.Errorf("part2(%v) = %v, want %v", tc.lines, got, tc.expected)
			}
		})
	}
}
