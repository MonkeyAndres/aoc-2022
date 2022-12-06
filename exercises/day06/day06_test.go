package day06

import (
	"os"
	"testing"
)

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
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expected: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 11},

		{input: string(finalInput), expected: 1566},
	}

	for index, v := range cases {
		result := part1(v.input, 4)

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
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", expected: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 26},

		{input: string(finalInput), expected: 2265},
	}

	for index, v := range cases {
		result := part1(v.input, 14)

		if result != v.expected {
			t.Errorf("Expected value %d but received %d", v.expected, result)
		} else {
			t.Logf("Test #%d passed!", index)
		}
	}
}
