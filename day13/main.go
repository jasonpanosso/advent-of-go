package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	fmt.Println("== Day 13 ==")

	part1Output, err := part1(input)
	if err != nil {
		fmt.Println("Error solving part 1:", err)
	} else {
		fmt.Println("Part1 output:", part1Output)
	}

	part2Output, err := part2(input)
	if err != nil {
		fmt.Println("Error solving part 2:", err)
	} else {
		fmt.Println("Part2 output:", part2Output)
	}
}

func part1(str string) (int, error) {
	sum := 0

	patterns := parseInput(str)
	fmt.Println(patterns)

	for _, rows := range patterns {
		// setup cols
		var cols []string
		for i := 0; i < len(rows[0]); i++ {
			var col string
			for j := 0; j < len(rows); j++ {
				col = col + string(rows[j][i])
			}
			cols = append(cols, col)
		}

		sum += 100*findMirrorIndex(rows) + findMirrorIndex(cols)
		fmt.Println(cols)
	}

	return sum, nil
}

func part2(str string) (int, error) {
	sum := 0

	return sum, nil
}

func parseInput(str string) [][]string {
	lines := strings.Split(str, "\n")

	var patterns [][]string
	lastPatternSepIndex := 0
	for i, line := range lines {
		if line == "" && i != len(lines)-1 {
			patterns = append(patterns, lines[lastPatternSepIndex:i])
			lastPatternSepIndex = i + 1
		}
	}

	return patterns
}

func findMirrorIndex(strs []string) int {
	for i := 1; i < len(strs); i++ {
		mirrorLen := min(i, len(strs)-i)

		a := slices.Clone(strs[i-mirrorLen : i])
		slices.Reverse(a)

		b := strs[i : i+mirrorLen]
		if slices.Equal(a, b) {
			return i
		}
	}

	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
