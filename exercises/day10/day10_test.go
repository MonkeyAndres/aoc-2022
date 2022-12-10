package day10

import (
	"os"
	"testing"
)

const exampleInput = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

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
		{input: exampleInput, expected: 13140},
		{input: string(finalInput), expected: 12880},
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
// 		{input: exampleInput, expected: 1},
// 		{input: `R 5
// U 8
// L 8
// D 3
// R 17
// D 10
// L 25
// U 20`, expected: 36},
// 		{input: string(finalInput), expected: 2533},
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
