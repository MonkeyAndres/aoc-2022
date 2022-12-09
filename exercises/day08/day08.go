package day08

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type TreeMap struct {
	grid [][]int

	width  int
	height int
}

func parseTreeMap(input string) TreeMap {
	var grid [][]int

	for _, treeRow := range strings.Split(input, "\n") {
		var parsedRow []int

		for _, tree := range strings.Split(treeRow, "") {
			treeHeight, err := strconv.Atoi(tree)

			if err != nil {
				panic("Error parsing tree map")
			}

			parsedRow = append(parsedRow, treeHeight)
		}

		grid = append(grid, parsedRow)
	}

	height := len(grid)
	width := len(grid[0])

	return TreeMap{grid, width, height}
}

func (treeMap *TreeMap) isVisibleTop(point Point) bool {
	treeHeight := treeMap.grid[point.y][point.x]

	for y := point.y - 1; y >= 0; y-- {
		if treeMap.grid[y][point.x] >= treeHeight {
			return false
		}
	}

	return true
}

func (treeMap *TreeMap) isVisibleBottom(point Point) bool {
	treeHeight := treeMap.grid[point.y][point.x]

	for y := point.y + 1; y < treeMap.height; y++ {
		if treeMap.grid[y][point.x] >= treeHeight {
			return false
		}
	}

	return true
}

func (treeMap *TreeMap) isVisibleLeft(point Point) bool {
	treeHeight := treeMap.grid[point.y][point.x]

	for x := point.x - 1; x >= 0; x-- {
		if treeMap.grid[point.y][x] >= treeHeight {
			return false
		}
	}

	return true
}

func (treeMap *TreeMap) isVisibleRight(point Point) bool {
	treeHeight := treeMap.grid[point.y][point.x]

	for x := point.x + 1; x < treeMap.width; x++ {
		if treeMap.grid[point.y][x] >= treeHeight {
			return false
		}
	}

	return true
}

func part1(input string) (visibleTrees int) {
	treeMap := parseTreeMap(input)

	visibleTrees = (treeMap.width * 2) + (treeMap.height * 2) - 4

	for y := 1; y < treeMap.height-1; y++ {
		for x := 1; x < treeMap.width-1; x++ {
			point := Point{x, y}

			if treeMap.isVisibleTop(point) || treeMap.isVisibleBottom(point) ||
				treeMap.isVisibleLeft(point) || treeMap.isVisibleRight(point) {
				visibleTrees++
			}
		}
	}

	return
}
