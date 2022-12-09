package day08

import (
	"os"
	"testing"
)

const exampleInput = `30373
25512
65332
33549
35390`

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
		{input: exampleInput, expected: 21},
		{input: string(finalInput), expected: 1719},
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

// func TestPart2(t *testing.T) {
// 	finalInput, err := os.ReadFile("input.txt")

// 	if err != nil {
// 		t.Error("Missing final input file.")
// 	}

// 	cases := []testCase{
// 		{input: exampleInput, expected: 0},
// 		{input: string(finalInput), expected: 0},
// 	}

// 	for index, v := range cases {
// 		result := part2(v.input)

// 		if result != v.expected {
// 			t.Errorf("Expected value %d but received %d", v.expected, result)
// 		} else {
// 			t.Logf("Test #%d passed!", index)
// 		}
// 	}
// }
