package day06

import "strings"

const SEQUENCE_LENGTH = 4

func part1(input string) int {
	parsedBuffer := strings.Split(input, "")

	for index := range parsedBuffer {
		tempPacketStartSequence := ""

		for cursor := index; cursor < index+SEQUENCE_LENGTH; cursor++ {
			nextCharacter := parsedBuffer[cursor]

			if !strings.Contains(tempPacketStartSequence, nextCharacter) {
				tempPacketStartSequence += nextCharacter
			}
		}

		if len(tempPacketStartSequence) == SEQUENCE_LENGTH {
			return index + SEQUENCE_LENGTH
		}
	}

	panic("Unreachable: Buffer contains no start-of-packet marker")
}
