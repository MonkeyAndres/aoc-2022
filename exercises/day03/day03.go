package day03

import (
	"strings"
)

func parseRucksackCompartments(input string) (parsedRucksack [][]string) {
	for _, rucksack := range strings.Split(input, "\n") {
		compartmentDelimiter := len(rucksack) / 2

		firstCompartment := rucksack[0:compartmentDelimiter]
		secondCompartment := rucksack[compartmentDelimiter:]

		parsedRucksack = append(parsedRucksack, []string{firstCompartment, secondCompartment})
	}

	return
}

func getItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	} else {
		return int(item - 'A' + 27)
	}
}

func part1(input string) (prioritySum int) {
	rucksacks := parseRucksackCompartments(input)

	for _, rucksack := range rucksacks {
		firstCompartmentItemsMap := make(map[rune]bool)

		for _, item := range rucksack[0] {
			firstCompartmentItemsMap[item] = true
		}

		secondCompartmentItemsMap := make(map[rune]bool)

		for _, item := range rucksack[1] {
			if firstCompartmentItemsMap[item] && !secondCompartmentItemsMap[item] {
				prioritySum += getItemPriority(item)
			}

			secondCompartmentItemsMap[item] = true
		}
	}

	return
}

func part2(input string) (prioritySum int) {
	rucksacks := strings.Split(input, "\n")

	for cursor := 0; cursor < len(rucksacks); cursor += 3 {
		elvesGroup := rucksacks[cursor : cursor+3]

		firstElfItems := make(map[rune]bool)

		for _, item := range elvesGroup[0] {
			firstElfItems[item] = true
		}

		secondAndFirstElfItems := make(map[rune]bool)

		for _, item := range elvesGroup[1] {
			if firstElfItems[item] && !secondAndFirstElfItems[item] {
				secondAndFirstElfItems[item] = true
			}
		}

		commonToAllItems := make(map[rune]bool)

		for _, item := range elvesGroup[2] {
			if secondAndFirstElfItems[item] && !commonToAllItems[item] {
				prioritySum += getItemPriority(item)
			}

			commonToAllItems[item] = true
		}
	}

	return
}
