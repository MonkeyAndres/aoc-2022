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

type Shape struct {
	own      string
	opponent string
}

var (
	Rock     = Shape{"X", "A"}
	Paper    = Shape{"Y", "B"}
	Scissors = Shape{"Z", "C"}
)

func getRoundScore(opponentsMove string, yourMove string) int {
	var SHAPE_VALUE = map[string]int{
		Rock.own:     1,
		Paper.own:    2,
		Scissors.own: 3,
	}

	roundScore := 0

	if opponentsMove == Rock.opponent && yourMove == Paper.own ||
		opponentsMove == Paper.opponent && yourMove == Scissors.own ||
		opponentsMove == Scissors.opponent && yourMove == Rock.own {
		roundScore += 6
	}

	if opponentsMove == Rock.opponent && yourMove == Rock.own ||
		opponentsMove == Paper.opponent && yourMove == Paper.own ||
		opponentsMove == Scissors.opponent && yourMove == Scissors.own {
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
		Rock.opponent:     Scissors.own,
		Paper.opponent:    Rock.own,
		Scissors.opponent: Paper.own,
	}

	var winningMoves = map[string]string{
		Rock.opponent:     Paper.own,
		Paper.opponent:    Scissors.own,
		Scissors.opponent: Rock.own,
	}

	var drawMoves = map[string]string{
		Rock.opponent:     Rock.own,
		Paper.opponent:    Paper.own,
		Scissors.opponent: Scissors.own,
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
