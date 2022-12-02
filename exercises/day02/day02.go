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

func getRoundScore(opponentsMove string, yourMove string) int {
	var SHAPE_VALUE = map[string]int{
		"X": 1, // ROCK
		"Y": 2, // PAPER
		"Z": 3, // SCISSORS
	}

	roundScore := 0

	if opponentsMove == "A" && yourMove == "Y" ||
		opponentsMove == "B" && yourMove == "Z" ||
		opponentsMove == "C" && yourMove == "X" {
		roundScore += 6
	}

	if opponentsMove == "A" && yourMove == "X" ||
		opponentsMove == "B" && yourMove == "Y" ||
		opponentsMove == "C" && yourMove == "Z" {
		roundScore += 3
	}

	roundScore += SHAPE_VALUE[yourMove]

	return roundScore
}

func part1(input string) int {
	rounds := parseInput(input)
	score := 0

	for _, round := range rounds {
		roundScore := getRoundScore(round[0], round[1])
		score += roundScore
	}

	return score
}

/**

X - LOSE
Y - DRAW
Z - WIN

*/

func part2(input string) int {
	rounds := parseInput(input)
	score := 0

	var losingMoves = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	var winningMoves = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	var drawMoves = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	for _, round := range rounds {
		opponentsMove := round[0]
		yourMove := ""

		roundResult := round[1]

		switch roundResult {
		case "X":
			{
				yourMove = losingMoves[opponentsMove]
			}

		case "Y":
			{
				yourMove = drawMoves[opponentsMove]
			}

		case "Z":
			{
				yourMove = winningMoves[opponentsMove]
			}
		}

		roundScore := getRoundScore(opponentsMove, yourMove)
		score += roundScore
	}

	return score
}
