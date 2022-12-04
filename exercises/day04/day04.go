package day04

import (
	"strconv"
	"strings"
)

func stringSliceToInt(slice []string) (numbers []int) {
	for _, item := range slice {
		num, err := strconv.Atoi(item)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	return
}

func parseInput(input string) (parsedAssignmentPairs [][][]int) {
	rawAssignmentPairs := strings.Split(input, "\n")

	for _, pair := range rawAssignmentPairs {
		parsedPair := strings.Split(pair, ",")

		left := stringSliceToInt(strings.Split(parsedPair[0], "-"))
		right := stringSliceToInt(strings.Split(parsedPair[1], "-"))

		// NOTE: Swap intervals to have the lowest start at index 0
		if left[0] > right[0] {
			parsedAssignmentPairs = append(parsedAssignmentPairs, [][]int{right, left})
		} else if left[0] == right[0] && left[1] < right[1] {
			parsedAssignmentPairs = append(parsedAssignmentPairs, [][]int{right, left})
		} else {
			parsedAssignmentPairs = append(parsedAssignmentPairs, [][]int{left, right})
		}
	}

	return
}

func part1(input string) (overlappingJobsCount int) {
	assignmentPairs := parseInput(input)

	for _, pair := range assignmentPairs {
		leftPair := pair[0]
		rightPair := pair[1]

		if leftPair[0] <= rightPair[0] && leftPair[1] >= rightPair[1] {
			overlappingJobsCount++
		}
	}

	return
}

func part2(input string) (overlappingJobsCount int) {
	assignmentPairs := parseInput(input)

	for _, pair := range assignmentPairs {
		leftPair := pair[0]
		rightPair := pair[1]

		if leftPair[0] <= rightPair[1] && leftPair[1] >= rightPair[0] {
			overlappingJobsCount++
		}
	}

	return
}
