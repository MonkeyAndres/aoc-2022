package day09

import (
	"os"
	"testing"
)

const exampleInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

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
		{input: exampleInput, expected: 13},
		{input: string(finalInput), expected: 6023},
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
		{input: exampleInput, expected: 1},
		{input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`, expected: 36},
		{input: string(finalInput), expected: 2533},
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
