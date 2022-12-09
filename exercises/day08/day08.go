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

func (treeMap *TreeMap) isVisibleTop(point Point) (bool, int) {
	treeHeight := treeMap.grid[point.y][point.x]
	treesInView := 0

	for y := point.y - 1; y >= 0; y-- {
		treesInView++

		if treeMap.grid[y][point.x] >= treeHeight {
			return false, treesInView
		}
	}

	return true, treesInView
}

func (treeMap *TreeMap) isVisibleBottom(point Point) (bool, int) {
	treeHeight := treeMap.grid[point.y][point.x]
	treesInView := 0

	for y := point.y + 1; y < treeMap.height; y++ {
		treesInView++

		if treeMap.grid[y][point.x] >= treeHeight {
			return false, treesInView
		}
	}

	return true, treesInView
}

func (treeMap *TreeMap) isVisibleLeft(point Point) (bool, int) {
	treeHeight := treeMap.grid[point.y][point.x]
	treesInView := 0

	for x := point.x - 1; x >= 0; x-- {
		treesInView++

		if treeMap.grid[point.y][x] >= treeHeight {
			return false, treesInView
		}
	}

	return true, treesInView
}

func (treeMap *TreeMap) isVisibleRight(point Point) (bool, int) {
	treeHeight := treeMap.grid[point.y][point.x]
	treesInView := 0

	for x := point.x + 1; x < treeMap.width; x++ {
		treesInView++

		if treeMap.grid[point.y][x] >= treeHeight {
			return false, treesInView
		}
	}

	return true, treesInView
}

func part1(input string) (visibleTrees int) {
	treeMap := parseTreeMap(input)

	visibleTrees = (treeMap.width * 2) + (treeMap.height * 2) - 4

	for y := 1; y < treeMap.height-1; y++ {
		for x := 1; x < treeMap.width-1; x++ {
			point := Point{x, y}

			topVisible, _ := treeMap.isVisibleTop(point)
			bottomVisible, _ := treeMap.isVisibleBottom(point)
			leftVisible, _ := treeMap.isVisibleLeft(point)
			rightVisible, _ := treeMap.isVisibleRight(point)

			if topVisible || bottomVisible || leftVisible || rightVisible {
				visibleTrees++
			}
		}
	}

	return
}

func part2(input string) (highestScenicScore int) {
	treeMap := parseTreeMap(input)

	for y := 1; y < treeMap.height-1; y++ {
		for x := 1; x < treeMap.width-1; x++ {
			point := Point{x, y}

			_, topTreesInView := treeMap.isVisibleTop(point)
			_, bottomTreesInView := treeMap.isVisibleBottom(point)
			_, leftTreesInView := treeMap.isVisibleLeft(point)
			_, rightTreesInView := treeMap.isVisibleRight(point)

			scenicScore := topTreesInView * bottomTreesInView * leftTreesInView * rightTreesInView

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return
}
