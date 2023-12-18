package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type SpringState int32

const (
	operational SpringState = iota
	damaged
	unknown
)

type SpringRow struct {
	states              []SpringState
	damagedSpringGroups []int
}

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	fmt.Println("== Day 12 ==")

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

// brute force.. zzzz
func part1(input string) (int, error) {
	sum := 0

	rows, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, row := range rows {
		sum += generateCombinations(row.states, row.damagedSpringGroups, 0)
	}

	return sum, nil
}

func part2(input string) (int, error) {
	sum := 0

	return sum, nil
}

func parseInput(str string) ([]SpringRow, error) {
	var output []SpringRow

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		spaceIndex := strings.Index(line, " ")
		if spaceIndex == -1 {
			return nil, fmt.Errorf("Invalid line format: %v", line)
		}

		states, err := convStringToSpringStates(line[:spaceIndex])
		if err != nil {
			return nil, err
		}

		groupedDamagedSprings, err := convNumStringsToInts(line[spaceIndex+1:])
		if err != nil {
			return nil, err
		}

		output = append(output,
			SpringRow{
				states:              states,
				damagedSpringGroups: groupedDamagedSprings,
			})
	}

	return output, nil
}

func convStringToSpringStates(str string) ([]SpringState, error) {
	var output []SpringState

	runeToSpringState := map[rune]SpringState{
		'.': operational,
		'?': unknown,
		'#': damaged,
	}

	for _, char := range str {
		state, exists := runeToSpringState[char]

		if !exists {
			return nil, fmt.Errorf("Unexpected input format provided to convStringToSpringStates, %v", char)
		}

		output = append(output, state)
	}

	return output, nil
}

func convNumStringsToInts(nums string) ([]int, error) {
	var output []int

	splitNums := strings.Split(nums, ",")
	for _, numString := range splitNums {
		num, err := strconv.Atoi(numString)
		if err != nil {
			return nil, fmt.Errorf("Error converting string to int during convNumStringsToInts, %v", numString)
		}

		output = append(output, num)
	}

	return output, nil
}

func generateCombinations(states []SpringState, damagedSpringGroups []int, index int) int {
	if index == len(states) {
		if isValidCombination(states, damagedSpringGroups) {
			return 1
		}
		return 0
	}

	if states[index] != unknown {
		return generateCombinations(states, damagedSpringGroups, index+1)
	} else {
		count := 0

		// Test branches where state at index is operational, and damaged
		states[index] = operational
		count += generateCombinations(states, damagedSpringGroups, index+1)

		states[index] = damaged
		count += generateCombinations(states, damagedSpringGroups, index+1)

		// Reset state
		states[index] = unknown

		return count
	}
}

func isValidCombination(states []SpringState, damagedSpringGroups []int) bool {
	groupIndex := 0
	currentGroupSize := 0

	for _, state := range states {
		if state == damaged {
			currentGroupSize++
		} else if currentGroupSize > 0 {
			if groupIndex >= len(damagedSpringGroups) || currentGroupSize != damagedSpringGroups[groupIndex] {
				return false
			}
			currentGroupSize = 0
			groupIndex++
		}
	}

	// Handle when the last group in states matches the last group in damagedSpringGroups
	if currentGroupSize > 0 {
		if groupIndex >= len(damagedSpringGroups) || currentGroupSize != damagedSpringGroups[groupIndex] {
			return false
		}
		groupIndex++
	}

	// Verify all groups were matched
	return groupIndex == len(damagedSpringGroups)
}
