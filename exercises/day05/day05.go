package day05

import (
	"reflect"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) PushN(elements []string) {
	*s = append(*s, elements...)
}

func (s *Stack) PopN(amount int) ([]string, bool) {
	if s.IsEmpty() {
		return []string{}, false
	} else {
		index := len(*s) - amount
		element := (*s)[index:]

		*s = (*s)[:index]

		return element, true
	}
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func parseCratesStack(stacksInput string) (stacks []Stack) {
	rows := strings.Split(stacksInput, "\n")
	reverse(rows)

	stacksRow := rows[0]

	for cursor := 1; cursor < len(stacksRow); cursor += 4 {
		stacks = append(stacks, Stack{})
	}

	boxesRows := rows[1:]

	for _, row := range boxesRows {
		cursor := 1

		for stackIndex := range stacks {
			boxValue := string(row[cursor])

			if boxValue != " " {
				stacks[stackIndex].Push(boxValue)
			}

			cursor += 4
		}
	}

	return
}

type Instruction struct {
	amount int
	from   int // NOTE: Must be normalized to 0 index
	to     int // NOTE: Must be normalized to 0 index
}

func parseInstructions(input string) (parsedInstructions []Instruction) {
	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		instructionParts := strings.Split(instruction, " ")

		amount, amountErr := strconv.Atoi(instructionParts[1])
		from, fromErr := strconv.Atoi(instructionParts[3])
		to, toErr := strconv.Atoi(instructionParts[5])

		if amountErr != nil || fromErr != nil || toErr != nil {
			panic("Error parsing instruction parts")
		}

		parsedInstructions = append(
			parsedInstructions,
			Instruction{amount, from - 1, to - 1},
		)
	}

	return
}

func part1(input string) (sequence string) {
	inputParts := strings.Split(input, "\n\n")

	stacks := parseCratesStack(inputParts[0])
	instructions := parseInstructions(inputParts[1])

	// RUNNER

	for _, instruction := range instructions {
		for i := 0; i < instruction.amount; i++ {
			crate, found := stacks[instruction.from].Pop()

			if !found {
				panic("Cannot get more items from stack")
			}

			stacks[instruction.to].Push(crate)
		}
	}

	// BUILD SEQUENCE

	for stackIndex := range stacks {
		crate, found := stacks[stackIndex].Pop()

		if found {
			sequence += crate
		}
	}

	return
}

func part2(input string) (sequence string) {
	inputParts := strings.Split(input, "\n\n")

	stacks := parseCratesStack(inputParts[0])
	instructions := parseInstructions(inputParts[1])

	// RUNNER

	for _, instruction := range instructions {
		crates, found := stacks[instruction.from].PopN(instruction.amount)

		if found {
			stacks[instruction.to].PushN(crates)
		}
	}

	// BUILD SEQUENCE

	for stackIndex := range stacks {
		crate, found := stacks[stackIndex].Pop()

		if found {
			sequence += crate
		}
	}

	return
}
