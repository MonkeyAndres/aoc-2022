package day1

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	var parsedElfMeals [][]int

	for _, elfMeals := range strings.Split(input, "\n\n") {
		var elfNumericCalories []int

		for _, mealCalories := range strings.Split(elfMeals, "\n") {
			numericCalorie, parsingError := strconv.Atoi(mealCalories)

			if parsingError != nil {
				panic(parsingError)
			}

			elfNumericCalories = append(elfNumericCalories, numericCalorie)
		}

		parsedElfMeals = append(parsedElfMeals, elfNumericCalories)
	}

	return parsedElfMeals
}

func sum(items []int) (totalCalories int) {
	total := 0

	for _, value := range items {
		total = total + value
	}

	return total
}

func part1(input string) (result int) {
	parsedElfMeals := parseInput(input)

	maxCalorieCount := 0

	for _, elfMeals := range parsedElfMeals {
		totalCalories := sum(elfMeals)

		if maxCalorieCount < totalCalories {
			maxCalorieCount = totalCalories
		}
	}

	return maxCalorieCount
}

func part2(input string) (result int) {
	parsedElfMeals := parseInput(input)

	var totalElfCalories []int

	for _, elfMeals := range parsedElfMeals {
		totalCalories := sum(elfMeals)
		totalElfCalories = append(totalElfCalories, totalCalories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalElfCalories)))

	return totalElfCalories[0] + totalElfCalories[1] + totalElfCalories[2]
}
