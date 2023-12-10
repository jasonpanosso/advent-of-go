package main

import (
	_ "embed"
	"fmt"
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
	fmt.Println("== Day 08 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

var nodeMap = map[string][]string{}

func part1(lines []string) int {
	sum := 0

	var directions []int
	for _, char := range lines[0] {
		if char == 'L' {
			directions = append(directions, 0)
		} else if char == 'R' {
			directions = append(directions, 1)
		}
	}

	for _, line := range lines[1:] {
		split := strings.Split(line, " = ")
		if len(split) != 2 {
			continue
		}

		key := split[0]
		r := strings.NewReplacer(" ", "", "(", "", ")", "")

		nodeMap[key] = strings.Split(r.Replace(split[1]), ",")

	}

	next := "AAA"
	var nodes []string

	for {

		for _, direction := range directions {
			nodes = nodeMap[next]
			next = nodes[direction]
			sum += 1

			if next == "ZZZ" {
				break
			}
		}

		if next == "ZZZ" {
			break
		}
	}

	return sum
}

// TODO: Refactor to LCM
func part2(lines []string) int {
	sum := 0

	var directions []int
	for _, char := range lines[0] {
		if char == 'L' {
			directions = append(directions, 0)
		} else if char == 'R' {
			directions = append(directions, 1)
		}
	}

	for _, line := range lines[1:] {
		split := strings.Split(line, " = ")
		if len(split) != 2 {
			continue
		}

		key := split[0]
		r := strings.NewReplacer(" ", "", "(", "", ")", "")

		nodeMap[key] = strings.Split(r.Replace(split[1]), ",")

	}

	var startingNodes []string

	for key := range nodeMap {
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}

	fmt.Println(startingNodes)

	allEndWithZ := true
	for {
		for _, direction := range directions {
			allEndWithZ = true
			var newNodes []string

			for _, node := range startingNodes {
				nodes := nodeMap[node]
				next := nodes[direction]
				newNodes = append(newNodes, next)

				if next[2] != 'Z' {
					// fmt.Println(next)
					allEndWithZ = false
				}
			}

			startingNodes = newNodes
			sum += 1

			if allEndWithZ {
				break
			}
		}

		if allEndWithZ {
			break
		}
	}

	fmt.Println(startingNodes)
	return sum
}
