package day11

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	worryLevel int
}

type WorryOperation struct {
	operation string
	argument  int
}

type MonkeyTest struct {
	divisor     int
	trueBranch  int
	falseBranch int
}

type Monkey struct {
	items          []Item
	itemsInspected int

	worryOperation *WorryOperation
	test           *MonkeyTest
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (monkey *Monkey) inspectNextItem() (int, Item) {
	if len(monkey.items) == 0 {
		return 0, Item{}
	}

	nextItem := monkey.items[0]
	monkey.items = monkey.items[1:]

	monkey.itemsInspected++

	switch monkey.worryOperation.operation {
	case "+":
		nextItem.worryLevel = nextItem.worryLevel + monkey.worryOperation.argument
	case "*":
		nextItem.worryLevel = nextItem.worryLevel * monkey.worryOperation.argument
	case "^":
		nextItem.worryLevel = int(math.Pow(float64(nextItem.worryLevel), 2))
	}

	nextItem.worryLevel /= 3

	passesTest := nextItem.worryLevel%monkey.test.divisor == 0

	if passesTest {
		return monkey.test.trueBranch, nextItem
	} else {
		return monkey.test.falseBranch, nextItem
	}
}

func parseMonkeys(input string) (parsedMonkeys []Monkey) {
	rawMonkeys := strings.Split(input, "\n\n")

	for _, rawMonkey := range rawMonkeys {
		monkeyParts := strings.Split(rawMonkey, "\n")

		items := strings.Split(monkeyParts[1][18:], ", ")
		var parsedItems []Item

		for _, item := range items {
			numberValue, err := strconv.Atoi(item)
			check(err)

			parsedItems = append(parsedItems, Item{numberValue})
		}

		operator := monkeyParts[2][23:24]
		argument := monkeyParts[2][25:]

		if argument == "old" {
			operator = "^"
			argument = "2"
		}

		parsedArgument, err := strconv.Atoi(argument)
		check(err)

		divisor, err := strconv.Atoi(monkeyParts[3][21:])
		check(err)

		trueBranch, err := strconv.Atoi(monkeyParts[4][29:])
		check(err)

		falseBranch, err := strconv.Atoi(monkeyParts[5][30:])
		check(err)

		parsedMonkeys = append(parsedMonkeys,
			Monkey{
				parsedItems,
				0,
				&WorryOperation{
					operator,
					parsedArgument,
				},
				&MonkeyTest{
					divisor,
					trueBranch,
					falseBranch,
				},
			},
		)
	}

	return
}

func part1(input string, roundsToRun int) (monkeyBusiness int) {
	monkeys := parseMonkeys(input)

	for round := 0; round < roundsToRun; round++ {
		for monkeyIndex := range monkeys {
			for {
				toMonkey, item := monkeys[monkeyIndex].inspectNextItem()

				if item.worryLevel == 0 {
					break
				}

				monkeys[toMonkey].items = append(monkeys[toMonkey].items, item)
			}
		}
	}

	var itemsInspectedByMonkey []int

	for _, monkey := range monkeys {
		itemsInspectedByMonkey = append(itemsInspectedByMonkey, monkey.itemsInspected)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(itemsInspectedByMonkey)))

	return itemsInspectedByMonkey[0] * itemsInspectedByMonkey[1]
}
