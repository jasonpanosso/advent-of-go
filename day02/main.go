package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

var re = regexp.MustCompile(`\d+`)

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	lines = strings.Split(input, "\n")
}

func main() {
	fmt.Println("== Day 02 ==")

	games := splitStringsByGame(lines)
	gamesWithTurns := splitGamesByTurns(games)
	gamesWithTurnsWithMoves := splitTurnsByMove(gamesWithTurns)

	part1Output := part1(gamesWithTurnsWithMoves)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(gamesWithTurnsWithMoves)
	fmt.Println("Part2 output:", part2Output)
}

func part1(games [][][]string) int {
	sum := 0

	maxReds := 12
	maxGreens := 13
	maxBlues := 14

	for gameNumber, game := range games {
		validGame := true
		for _, turn := range game {
			var colorCounts = map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}

			for _, move := range turn {
				numberString := re.FindString(move)

				if len(numberString) == 0 {
					panic("invalid move parsed")
				}

				number, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}

				for color := range colorCounts {
					if strings.HasSuffix(move, color) {
						colorCounts[color] += number
					}
				}
			}

			if colorCounts["red"] > maxReds || colorCounts["blue"] > maxBlues || colorCounts["green"] > maxGreens {
				validGame = false
				break
			}
		}

		if validGame {
			sum += gameNumber + 1
		}
	}

	return sum
}

func part2(games [][][]string) int {
	sum := 0

	for _, game := range games {
		var curMaxes = map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, turn := range game {
			var colorCounts = map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}

			for _, move := range turn {
				numberString := re.FindString(move)

				if len(numberString) == 0 {
					panic("invalid move parsed")
				}

				number, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}

				for color := range colorCounts {
					if strings.HasSuffix(move, color) {
						colorCounts[color] += number
					}
				}
			}

			for color, count := range colorCounts {
				if count > curMaxes[color] {
					curMaxes[color] = count
				}
			}
		}

		power := 0

		for _, count := range curMaxes {
			if power == 0 {
				power = count
			} else {
				power *= count
			}
		}
		sum += power
	}

	return sum
}

func splitStringsByGame(lines []string) []string {
	output := make([]string, 0)
	for _, line := range lines {
		split := strings.Split(line, ":")

		if len(split) != 2 {
			continue
		}

		output = append(output, split[1])
	}
	return output
}

func splitGamesByTurns(games []string) [][]string {
	turns := make([][]string, 0)

	for _, game := range games {
		splitByTurn := strings.Split(game, ";")
		turns = append(turns, splitByTurn)
	}
	return turns
}

func splitTurnsByMove(games [][]string) [][][]string {
	output := make([][][]string, 0)
	for _, game := range games {
		splitByMove := make([][]string, 0)
		for _, turn := range game {
			moves := strings.Split(turn, ",")

			for i, move := range moves {
				moves[i] = strings.TrimSpace(move)
			}

			splitByMove = append(splitByMove, moves)
		}
		output = append(output, splitByMove)
	}
	return output
}
