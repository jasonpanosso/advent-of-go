package main

import (
	_ "embed"
	"fmt"
	"math"
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
	fmt.Println("== Day 04 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {

		set := make(map[string]bool)

		splitOffGamePrefix := strings.Split(line, ":")

		if len(splitOffGamePrefix) != 2 {
			continue
		}

		splitOffWinners := strings.Split(splitOffGamePrefix[1], "|")

		if len(splitOffWinners) != 2 {
			continue
		}

		winners := strings.Split(strings.TrimSpace(splitOffWinners[0]), " ")
		ourNums := strings.Split(strings.TrimSpace(splitOffWinners[1]), " ")

		for _, winner := range winners {
			if winner == "" {
				continue
			}

			set[strings.TrimSpace(winner)] = true
		}

		count := 0.0

		for _, num := range ourNums {
			exists := set[strings.TrimSpace(num)]
			if exists {
				count += 1
			}
		}

		if count == 0.0 {
			continue
		}

		sum += int(math.Pow(2, count-1))

	}

	return sum
}

func part2(lines []string) int {
	sum := 0

	cardCounts := make(map[int]int)
	for i := 0; i < len(lines); i++ {
		cardCounts[i] = 1
	}

	for gameNumber, line := range lines {
		set := make(map[string]bool)

		splitOffGamePrefix := strings.Split(line, ":")

		if len(splitOffGamePrefix) != 2 {
			continue
		}

		splitOffWinners := strings.Split(splitOffGamePrefix[1], "|")

		if len(splitOffWinners) != 2 {
			continue
		}

		winners := strings.Split(strings.TrimSpace(splitOffWinners[0]), " ")
		ourNums := strings.Split(strings.TrimSpace(splitOffWinners[1]), " ")

		for _, winner := range winners {
			if winner == "" {
				continue
			}
			set[strings.TrimSpace(winner)] = true
		}

		repeat := cardCounts[gameNumber]
		count := 0
		for _, num := range ourNums {
			exists := set[strings.TrimSpace(num)]
			if exists {
				count += 1
				if gameNumber+count >= len(lines) {
					break
				}

				cardCounts[gameNumber+count] = cardCounts[gameNumber+count] + repeat
			}
		}

		sum += repeat
	}

	return sum
}
