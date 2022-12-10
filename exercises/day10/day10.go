package day10

import (
	"strconv"
	"strings"
)

type Instruction struct {
	command  string
	argument int
}

func parseInstructions(input string) (instructions []Instruction) {
	for _, rawInstruction := range strings.Split(input, "\n") {
		instructionParts := strings.Split(rawInstruction, " ")

		if instructionParts[0] == "noop" {
			instructions = append(instructions, Instruction{instructionParts[0], 0})
			continue
		}

		numericArgument, err := strconv.Atoi(instructionParts[1])

		if err != nil {
			panic("Error converting argument to int")
		}

		instructions = append(instructions, Instruction{instructionParts[0], numericArgument})
	}

	return
}

func runProgram(instructions []Instruction, cyclesToRun int) int {
	x := 1

	instructionPointer := 0

	var pendingInstruction Instruction
	workingUntil := -1

	for cycle := 0; cycle < cyclesToRun; cycle++ {
		if pendingInstruction.command != "" {
			if cycle < workingUntil {
				continue
			}

			x += pendingInstruction.argument

			instructionPointer++
			pendingInstruction = Instruction{}
		}

		instruction := instructions[instructionPointer]

		if instruction.command == "noop" {
			instructionPointer++
			continue
		}

		if instruction.command == "addx" {
			pendingInstruction = instruction
			workingUntil = cycle + 2
		}
	}

	return x
}

func part1(input string) (result int) {
	instructions := parseInstructions(input)

	cycleRuns := []int{20, 60, 100, 140, 180, 220}

	for _, cyclesToRun := range cycleRuns {
		result += runProgram(instructions, cyclesToRun) * cyclesToRun
	}

	return
}
