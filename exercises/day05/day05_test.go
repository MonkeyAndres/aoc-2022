package day05

import (
	"os"
	"testing"
)

const exampleInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

type testCase struct {
	input    string
	expected string
}

func TestPart1(t *testing.T) {
	finalInput, err := os.ReadFile("input.txt")

	if err != nil {
		t.Error("Missing final input file.")
	}

	cases := []testCase{
		{input: exampleInput, expected: "CMZ"},
		{input: string(finalInput), expected: "GFTNRBZPF"},
	}

	for index, v := range cases {
		result := part1(v.input)

		if result != v.expected {
			t.Errorf("Expected value %s but received %s", v.expected, result)
		} else {
			t.Logf("Test #%d passed!", index)
		}
	}
}
