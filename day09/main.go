package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
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
	fmt.Println("== Day 09 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	sum := 0

	parsedInput := parseInput(lines)

	for _, nums := range parsedInput {
		diffs := nums
		var lastNums []int

		for {
			var newDiffs []int

			for i := 0; i < len(diffs); i++ {
				if i+1 == len(diffs) {
					lastNums = append(lastNums, diffs[i])
					diffs = newDiffs
					break
				}

				cur := diffs[i]
				next := diffs[i+1]
				newDiffs = append(newDiffs, next-cur)

			}

			allDiffsZero := true
			for _, diff := range newDiffs {
				if diff != 0 {
					allDiffsZero = false
					break
				}
			}

			if allDiffsZero {
				break
			}
		}

		for i := len(lastNums) - 1; i >= 0; i-- {
			sum += lastNums[i]
		}
	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	parsedInput := parseInput(lines)

	for _, nums := range parsedInput {
		diffs := nums
		var firstNums []int

		for {
			var newDiffs []int

			for i := 0; i < len(diffs); i++ {
				if i+1 == len(diffs) {
					firstNums = append(firstNums, diffs[0])
					diffs = newDiffs
					break
				}

				cur := diffs[i]
				next := diffs[i+1]
				newDiffs = append(newDiffs, next-cur)

			}

			allDiffsZero := true
			for _, diff := range newDiffs {
				if diff != 0 {
					allDiffsZero = false
					break
				}
			}

			if allDiffsZero {
				break
			}
		}

		diffSum := 0
		for i := len(firstNums) - 1; i >= 0; i-- {
			diffSum = firstNums[i] - diffSum
		}

		sum += diffSum
	}

	return sum
}

func parseInput(lines []string) [][]int {
	var output [][]int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var nums []int
		numStrings := strings.Fields(line)

		for _, str := range numStrings {
			num, err := strconv.Atoi(str)

			if err != nil {
				panic("Non number provided in input.txt")
			}

			nums = append(nums, num)
		}

		output = append(output, nums)
	}

	return output
}
