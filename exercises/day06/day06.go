package day06

import "strings"

// NOTE: Apparently this also works for part2
// sequenceLength = 4 (part1) and 14 (part2)

func part1(input string, sequenceLength int) int {
	parsedBuffer := strings.Split(input, "")

	for index := range parsedBuffer {
		tempPacketStartSequence := ""

		for cursor := index; cursor < index+sequenceLength; cursor++ {
			nextCharacter := parsedBuffer[cursor]

			if !strings.Contains(tempPacketStartSequence, nextCharacter) {
				tempPacketStartSequence += nextCharacter
			}
		}

		if len(tempPacketStartSequence) == sequenceLength {
			return index + sequenceLength
		}
	}

	panic("Unreachable: Buffer contains no start-of-packet marker")
}
