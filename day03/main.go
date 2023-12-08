package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string
var lines []string

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	lines = strings.Split(input, "\n")
}

func main() {
	fmt.Println("== Day 03 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	sum := 0
	deltas := []struct{ i, j int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*skip middle*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	rows := len(lines)
	cols := len(lines[0])

	for i, line := range lines {
		for j, char := range line {

			if char == '.' || unicode.IsDigit(char) {
				continue
			}

			set := make(map[string]bool)

			for _, delta := range deltas {
				di, dj := i+delta.i, j+delta.j
				// validate index is in bounds
				if di < 0 || di >= rows || dj < 0 || dj >= cols {
					continue
				}

				neighbor := lines[di][dj]

				if !unicode.IsDigit(rune(neighbor)) {
					continue
				}

				// find neighbor's neighbors
				left := 0
				right := 0
				for {
					leftChanged := false
					rightChanged := false

					if dj-left-1 >= 0 && unicode.IsDigit(rune(lines[di][dj-left-1])) {
						left += 1
						leftChanged = true
					}

					if dj+right+1 < cols && unicode.IsDigit(rune(lines[di][dj+right+1])) {
						right += 1
						rightChanged = true
					}

					if !leftChanged && !rightChanged {
						break
					}
				}

				set[lines[di][dj-left:dj+right+1]] = true
			}

			for strNum := range set {
				num, err := strconv.Atoi(strNum)

				if err != nil {
					fmt.Println("Error during str to int conversion:", err)
				}

				sum += num
			}
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0
	deltas := []struct{ i, j int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0} /*skip middle*/, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	rows := len(lines)
	cols := len(lines[0])

	for i, line := range lines {
		for j, char := range line {

			if char != '*' {
				continue
			}

			set := make(map[string]bool)

			for _, delta := range deltas {
				di, dj := i+delta.i, j+delta.j
				// validate index is in bounds
				if di < 0 || di >= rows || dj < 0 || dj >= cols {
					continue
				}

				neighbor := lines[di][dj]

				if !unicode.IsDigit(rune(neighbor)) {
					continue
				}

				// find neighbor's neighbors
				left := 0
				right := 0
				for {
					leftChanged := false
					rightChanged := false

					if dj-left-1 >= 0 && unicode.IsDigit(rune(lines[di][dj-left-1])) {
						left += 1
						leftChanged = true
					}

					if dj+right+1 < cols && unicode.IsDigit(rune(lines[di][dj+right+1])) {
						right += 1
						rightChanged = true
					}

					if !leftChanged && !rightChanged {
						break
					}
				}

				set[lines[di][dj-left:dj+right+1]] = true
			}

			if len(set) == 2 {
				nums := make([]int, 2)
				i := 0

				for strNum := range set {
					num, err := strconv.Atoi(strNum)

					if err != nil {
						fmt.Println("Error during str to int conversion:", err)
					}

					nums[i] = num
					i += 1
				}

				sum += nums[0] * nums[1]
			}
		}
	}

	return sum
}
