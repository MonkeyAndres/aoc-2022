package day1

import (
	"strconv"
	"strings"
)

func toNumericCalories(calories []string) []int {
	var numericCalories []int

	for _, calorie := range calories {
		numericCalorie, parsingError := strconv.Atoi(calorie)

		if parsingError != nil {
			panic(parsingError)
		}

		numericCalories = append(numericCalories, numericCalorie)
	}

	return numericCalories
}

func sum(items []int) (totalCalories int) {
	total := 0

	for _, value := range items {
		total = total + value
	}

	return total
}

func part1(input string) (result int) {
	mealsPerElf := strings.Split(input, "\n\n")

	maxCalorieCount := 0

	for _, counter := range mealsPerElf {
		calories := toNumericCalories(strings.Split(counter, "\n"))
		totalCalories := sum(calories)

		if maxCalorieCount < totalCalories {
			maxCalorieCount = totalCalories
		}
	}

	return maxCalorieCount
}
