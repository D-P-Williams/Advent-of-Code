package main

import (
	"fmt"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type coord struct {
	row int
	col int
}

var nodesWithAntiNode = map[coord]bool{}

func main() {
	grid := aoc.ReadGrid("./input.txt")

	countPt1 := 0
	countPt2 := 0

	grid, _ = processGrid(grid, false)

	for _, row := range grid {
		fmt.Println(row)
		for _, cell := range row {
			if cell == "#" {
				countPt1 += 1
			}
		}
	}

	for range nodesWithAntiNode {
		countPt1 += 1
	}

	fmt.Println("part 1", countPt1)

	// Reset map
	nodesWithAntiNode = map[coord]bool{}

	grid, _ = processGrid(grid, true)

	for _, row := range grid {
		fmt.Println(row)
		for _, cell := range row {
			if cell == "#" {
				countPt2 += 1
			}
		}
	}

	for range nodesWithAntiNode {
		countPt2 += 1
	}

	fmt.Println("part 2", countPt2)
}

func processGrid(grid [][]string, part2 bool) ([][]string, int) {
	antiUnderNode := 0

	for rowIdx, row := range grid {
		for colIdx := range row {
			antennaValue := grid[rowIdx][colIdx]

			// Skip empty cells
			if antennaValue == "." || antennaValue == "#" {
				continue
			}

			grid, antiUnderNode = findAntiNodes(grid, antiUnderNode, antennaValue, rowIdx, colIdx, part2)
		}
	}

	return grid, antiUnderNode
}

func findAntiNodes(grid [][]string, antiUnderNode int, target string, currRowIdx, currColIdx int, part2 bool) ([][]string, int) {
	for rowIdx, row := range grid {
		for colIdx := range row {
			// Skip self
			if rowIdx == currRowIdx && colIdx == currColIdx {
				continue
			}

			if grid[rowIdx][colIdx] == target {
				grid, antiUnderNode = addAntiNode(grid, antiUnderNode, currRowIdx, currColIdx, rowIdx, colIdx, part2)
			}
		}
	}

	return grid, antiUnderNode
}

func addAntiNode(grid [][]string, antiUnderNode, currRowIdx, currColIdx, rowIdx, colIdx int, part2 bool) ([][]string, int) {
	diffRow := currRowIdx - rowIdx
	diffCol := currColIdx - colIdx

	var multipliers []int

	if part2 {
		for i := -1000; i < 1000; i++ {
			multipliers = append(multipliers, i)
		}
	} else {
		multipliers = []int{-1, 2}
	}

	for _, multiplier := range multipliers {
		antiNodeRowIdx := currRowIdx + multiplier*diffRow
		antiNodeColIdx := currColIdx + multiplier*diffCol

		grid, antiUnderNode = setAntiNode(grid, antiUnderNode, antiNodeRowIdx, antiNodeColIdx)
	}

	return grid, antiUnderNode
}

func setAntiNode(grid [][]string, antiUnderNode, row, col int) ([][]string, int) {
	// Early return if out of bounds
	if 0 > row || row >= len(grid) {
		return grid, antiUnderNode
	}

	if 0 > col || col >= len(grid[0]) {
		return grid, antiUnderNode
	}

	// If already antiNode return
	if grid[row][col] == "#" {
		return grid, antiUnderNode
	}

	// If non-empty, non-antiNode, cell (i.e. a node), increment antiUnderNode
	if grid[row][col] != "." {
		nodesWithAntiNode[coord{row: row, col: col}] = true
		return grid, antiUnderNode
	}

	grid[row][col] = "#"

	return grid, antiUnderNode
}
