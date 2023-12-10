package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

type Hand struct {
	cardStrengths []int
	bet           int
	value         int
}

func init() {
	if len(input) == 0 {
		panic("empty input.txt file")
	}

	lines = strings.Split(input, "\n")
}

func main() {
	fmt.Println("== Day 07 ==")

	part1Output := part1(lines)
	fmt.Println("Part1 output:", part1Output)

	part2Output := part2(lines)
	fmt.Println("Part2 output:", part2Output)
}

func part1(lines []string) int {
	var cardStrengthMap = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	sum := 0

	var hands []Hand
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		bet, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("could not convert bid to int")
		}

		var cardCounts = map[int]int{}
		var cards []int

		handString := fields[0]
		for _, char := range handString {
			card := cardStrengthMap[char]
			cards = append(cards, card)

			_, exists := cardCounts[card]

			if exists {
				cardCounts[card] += 1
			} else {
				cardCounts[card] = 1
			}
		}

		handValue := 0
		if len(cardCounts) == 1 {
			handValue = 6
		} else if len(cardCounts) == 2 {
			for _, count := range cardCounts {
				if count == 1 || count == 4 {
					handValue = 5
				} else {
					handValue = 4
				}
				break
			}
		} else if len(cardCounts) == 3 {
			isThreeOfKind := false
			for _, count := range cardCounts {
				if count == 3 {
					isThreeOfKind = true
					break
				}
			}

			if isThreeOfKind {
				handValue = 3
			} else {
				handValue = 2
			}
		} else if len(cardCounts) == 4 {
			handValue = 1
		} else if len(cardCounts) == 5 {
			handValue = 0
		}
		hands = append(hands, Hand{value: handValue, bet: bet, cardStrengths: cards})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			for n := 0; n < len(hands[i].cardStrengths); n++ {
				if hands[i].cardStrengths[n] != hands[j].cardStrengths[n] {
					return hands[i].cardStrengths[n] < hands[j].cardStrengths[n]
				}
			}

			panic("Unable to sort, hands are exactly the same")
		} else {
			return hands[i].value < hands[j].value
		}
	})

	for i, hand := range hands {
		sum += hand.bet * (i + 1)
	}

	return sum
}

func part2(lines []string) int {
	var cardStrengthMap = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	sum := 0

	var hands []Hand
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		bet, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("could not convert bid to int")
		}

		var cardCounts = map[int]int{}
		var cards []int

		handString := fields[0]
		for _, char := range handString {
			card := cardStrengthMap[char]
			cards = append(cards, card)

			_, exists := cardCounts[card]

			if exists {
				cardCounts[card] += 1
			} else {
				cardCounts[card] = 1
			}
		}

		// this is janky
		if cardCounts[1] > 0 {
			jokers := cardCounts[1]
			delete(cardCounts, 1)
			var highestCount int

			for card, count := range cardCounts {
				if highestCount == 0 {
					highestCount = card
				}
				if count > cardCounts[highestCount] || (count == cardCounts[highestCount] && card > highestCount) {
					highestCount = card
				}
			}

			cardCounts[highestCount] += jokers
		}

		handValue := 0
		if len(cardCounts) == 1 {
			handValue = 6
		} else if len(cardCounts) == 2 {
			for _, count := range cardCounts {
				if count == 1 || count == 4 {
					handValue = 5
				} else {
					handValue = 4
				}
				break
			}
		} else if len(cardCounts) == 3 {
			isThreeOfKind := false
			for _, count := range cardCounts {
				if count == 3 {
					isThreeOfKind = true
					break
				}
			}

			if isThreeOfKind {
				handValue = 3
			} else {
				handValue = 2
			}
		} else if len(cardCounts) == 4 {
			handValue = 1
		} else if len(cardCounts) == 5 {
			handValue = 0
		}
		hands = append(hands, Hand{value: handValue, bet: bet, cardStrengths: cards})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value == hands[j].value {
			for n := 0; n < len(hands[i].cardStrengths); n++ {
				if hands[i].cardStrengths[n] != hands[j].cardStrengths[n] {
					return hands[i].cardStrengths[n] < hands[j].cardStrengths[n]
				}
			}

			panic("Unable to sort, hands are exactly the same")
		} else {
			return hands[i].value < hands[j].value
		}
	})

	for i, hand := range hands {
		sum += hand.bet * (i + 1)
	}

	return sum
}
