package main

import (
	_ "embed"
	"fmt"
	"math"
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
	fmt.Println("== Day 06 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	sum := 0

	times, distances := part1ParseInput(lines)

	for i := 0; i < len(times); i++ {
		distance := distances[i]
		time := times[i]

		minTime := int((float64(time) - (math.Sqrt(float64((time * time) - (4 * distance))))) / 2)
		maxTime := int((float64(time) + (math.Sqrt(float64((time * time) - (4 * distance))))) / 2)

		if sum == 0 {
			sum = maxTime - minTime
		} else {
			sum *= maxTime - minTime
		}
	}

	return sum
}

func part2(lines []string) int {
	time, distance := part2ParseInput(lines)

	minTime := int((float64(time) - (math.Sqrt(float64((time * time) - (4 * distance))))) / 2)
	maxTime := int((float64(time) + (math.Sqrt(float64((time * time) - (4 * distance))))) / 2)

	return maxTime - minTime
}

func part1ParseInput(lines []string) ([]int, []int) {
	timeFields := strings.Fields(lines[0])[1:]
	distanceFields := strings.Fields(lines[1])[1:]

	var times []int
	var distances []int

	for i := 0; i < len(timeFields); i++ {
		time, timeErr := strconv.Atoi(timeFields[i])

		if timeErr != nil {
			panic("Unexpected input.txt format")
		}

		distance, distanceErr := strconv.Atoi(distanceFields[i])
		if distanceErr != nil {
			panic("Unexpected input.txt format")
		}

		times = append(times, time)
		distances = append(distances, distance)

	}
	return times, distances

}

func part2ParseInput(lines []string) (int, int) {
	timeString := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	distanceString := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")

	time, timeErr := strconv.Atoi(timeString)

	if timeErr != nil {
		panic("Unexpected input.txt format")
	}

	distance, distanceErr := strconv.Atoi(distanceString)
	if distanceErr != nil {
		panic("Unexpected input.txt format")
	}

	return time, distance
}
