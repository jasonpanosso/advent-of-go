package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

// Notes: Optimize with two pointers?

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
	fmt.Println("== Day 01 ==")
	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		var firstDigit, lastDigit int
		firstDigitFound := false

		for _, char := range line {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				if !firstDigitFound {
					firstDigit = digit
					firstDigitFound = true
				}
				lastDigit = digit
			}
		}

		if firstDigitFound {
			sum += firstDigit*10 + lastDigit
		}
	}

	return sum
}

var stringToNumMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2(lines []string) int {
	sum := 0

	for _, line := range lines {
		var firstDigit, lastDigit int
		firstDigitFound := false

		for i := 0; i < len(line); {
			digit := -1

			if unicode.IsDigit(rune(line[i])) {
				digit = int(line[i] - '0')
			} else {
				for key, value := range stringToNumMap {
					if strings.HasPrefix(line[i:], key) {
						digit = value
						break
					}
				}
			}

			i++

			if digit == -1 {
				continue
			}

			if !firstDigitFound {
				firstDigit = digit
				firstDigitFound = true
			}
			lastDigit = digit
		}

		if firstDigitFound {
			sum += firstDigit*10 + lastDigit
		}
	}

	return sum
}
