package day02

import (
	"strings"
)

func parseInput(input string) [][]string {
	rounds := strings.Split(input, "\n")

	var parsedRounds [][]string

	for _, round := range rounds {
		shapes := strings.Split(round, " ")

		parsedRounds = append(parsedRounds, shapes)
	}

	return parsedRounds
}

/**

OPPONENT
	ROCK (A)
	PAPER (B)
	SCISSORS (C)

YOU
	ROCK (X)
	PAPER (Y)
	SCISSORS (Z)

SCORE
	WIN (+6)
	DRAW (+3)
	LOSE (+0)

*/

var SHAPE_VALUE = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func part1(input string) int {
	rounds := parseInput(input)
	score := 0

	for _, round := range rounds {
		opponentsMove := round[0]
		yourMove := round[1]

		if opponentsMove == "A" && yourMove == "Y" ||
			opponentsMove == "B" && yourMove == "Z" ||
			opponentsMove == "C" && yourMove == "X" {
			score += 6
		}

		if opponentsMove == "A" && yourMove == "X" ||
			opponentsMove == "B" && yourMove == "Y" ||
			opponentsMove == "C" && yourMove == "Z" {
			score += 3
		}

		score += SHAPE_VALUE[yourMove]
	}

	return score
}
