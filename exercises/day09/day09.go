package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	steps     int
}

func parseInstructions(input string) (instructions []Instruction) {
	rawInstructions := strings.Split(input, "\n")

	for _, rawInstruction := range rawInstructions {
		instructionParts := strings.Split(rawInstruction, " ")

		numericSteps, err := strconv.Atoi(instructionParts[1])

		if err != nil {
			panic("Error converting steps to int")
		}

		instructions = append(instructions, Instruction{instructionParts[0], numericSteps})
	}

	return
}

func getDistance(head Knot, tail Knot) int {
	dx := head.x - tail.x
	dy := head.y - tail.y

	distance := math.Sqrt(math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2))

	return int(math.Round(distance))
}

type Knot struct {
	x int
	y int
}

func (knot *Knot) moveOne(direction string) {
	switch direction {
	case "U":
		knot.y++

	case "D":
		knot.y--

	case "L":
		knot.x--

	case "R":
		knot.x++
	}
}

func (tail *Knot) followHead(head *Knot) {
	if getDistance(*head, *tail) < 2 {
		return
	}

	if head.y > tail.y {
		tail.moveOne("U")
	} else if head.y < tail.y {
		tail.moveOne("D")
	}

	if head.x > tail.x {
		tail.moveOne("R")
	} else if head.x < tail.x {
		tail.moveOne("L")
	}

}

func part1(input string) int {
	instructions := parseInstructions(input)

	head := Knot{}
	tail := Knot{}

	tailPositions := make(map[string]bool, 0)

	for _, instruction := range instructions {
		for s := 0; s < instruction.steps; s++ {
			head.moveOne(instruction.direction)
			tail.followHead(&head)

			tailPosition := fmt.Sprintf("%d-%d", tail.x, tail.y)
			tailPositions[tailPosition] = true
		}
	}

	return len(tailPositions)
}

const NUMBER_OF_TAILS = 9

func part2(input string) int {
	instructions := parseInstructions(input)

	head := Knot{}

	var tails [NUMBER_OF_TAILS]Knot

	tailPositions := make(map[string]bool, 0)

	for _, instruction := range instructions {
		for s := 0; s < instruction.steps; s++ {
			head.moveOne(instruction.direction)

			for tailIndex := range tails {
				var tailHead *Knot

				if tailIndex == 0 {
					tailHead = &head
				} else {
					tailHead = &tails[tailIndex-1]
				}

				tails[tailIndex].followHead(tailHead)

				if tailIndex+1 == NUMBER_OF_TAILS {
					tailPosition := fmt.Sprintf("%d-%d", tails[tailIndex].x, tails[tailIndex].y)
					tailPositions[tailPosition] = true
				}
			}
		}
	}

	return len(tailPositions)
}
