package day04

import (
	"os"
	"testing"
)

const exampleInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

type testCase struct {
	input    string
	expected int
}

func TestPart1(t *testing.T) {
	finalInput, err := os.ReadFile("input.txt")

	if err != nil {
		t.Error("Missing final input file.")
	}

	cases := []testCase{
		{input: exampleInput, expected: 2},
		{input: string(finalInput), expected: 605},
	}

	for index, v := range cases {
		result := part1(v.input)

		if result != v.expected {
			t.Errorf("Expected value %d but received %d", v.expected, result)
		} else {
			t.Logf("Test #%d passed!", index)
		}
	}
}

func TestPart2(t *testing.T) {
	finalInput, err := os.ReadFile("input.txt")

	if err != nil {
		t.Error("Missing final input file.")
	}

	cases := []testCase{
		{input: exampleInput, expected: 4},
		{input: string(finalInput), expected: 914},
	}

	for index, v := range cases {
		result := part2(v.input)

		if result != v.expected {
			t.Errorf("Expected value %d but received %d", v.expected, result)
		} else {
			t.Logf("Test #%d passed!", index)
		}
	}
}
