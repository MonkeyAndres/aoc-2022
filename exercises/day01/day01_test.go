package day1

import (
	"os"
	"testing"
)

const exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

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
		{input: exampleInput, expected: 24000},
		{input: string(finalInput), expected: 74711},
	}

	for _, v := range cases {
		if part1(v.input) != v.expected {
			t.Errorf("Expected value %d", v.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	finalInput, err := os.ReadFile("input.txt")

	if err != nil {
		t.Error("Missing final input file.")
	}

	cases := []testCase{
		{input: exampleInput, expected: 45000},
		{input: string(finalInput), expected: 209481},
	}

	for _, v := range cases {
		result := part2(v.input)

		if result != v.expected {
			t.Errorf("Expected value %d but received %d", v.expected, result)
		}
	}
}
