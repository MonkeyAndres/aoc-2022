package day1

import (
	"strconv"
	"strings"
)

func sumCalories(calories []string) (totalCalories int) {
	total := 0

	for _, calorie := range calories {
		numericCalorie, parsingError := strconv.Atoi(calorie)

		if parsingError != nil {
			panic(parsingError)
		}

		total = total + numericCalorie
	}

	return total
}

func part1(input string) (result int) {
	mealsPerElf := strings.Split(input, "\n\n")

	maxCalorieCount := 0

	for _, counter := range mealsPerElf {
		calories := strings.Split(counter, "\n")

		totalCalories := sumCalories(calories)

		if maxCalorieCount < totalCalories {
			maxCalorieCount = totalCalories
		}
	}

	return maxCalorieCount
}
