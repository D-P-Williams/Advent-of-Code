package main

import (
	"fmt"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	grid := aoc.ReadGrid("./input.txt")

	countPt1 := 0
	countPt2 := 0

	grid, _ = processGrid(grid)

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

	fmt.Println("part 2", countPt2)
}

// func processGrid(grid [][]string) (count int) {
// 	for rowIdx, row := range grid {
// 		for colIdx := range row {
// 			antennaValue := grid[rowIdx][colIdx]

// 			// Skip empty cells
// 			if antennaValue == "." || antennaValue == "#" {
// 				continue
// 			}

// 			count += findAntiNodes(grid, antennaValue, rowIdx, colIdx)
// 		}
// 	}

// 	return count
// }

// func findAntiNodes(grid [][]string, target string, currRowIdx, currColIdx int) (count int) {
// 	for rowIdx, row := range grid {
// 		for colIdx := range row {
// 			// Skip self
// 			if rowIdx == currRowIdx && colIdx == currColIdx {
// 				continue
// 			}

// 			if grid[rowIdx][colIdx] == target {
// 				count += addAntiNode(grid, currRowIdx, currColIdx, rowIdx, colIdx)
// 			}
// 		}
// 	}

// 	return count
// }

// func addAntiNode(grid [][]string, currRowIdx, currColIdx, rowIdx, colIdx int) (count int) {
// 	diffRow := currRowIdx - rowIdx
// 	diffCol := currColIdx - colIdx

// 	antiNode1RowIdx := currRowIdx + diffRow
// 	antiNode1ColIdx := currColIdx + diffCol

// 	antiNode2RowIdx := currRowIdx - 2*diffRow
// 	antiNode2ColIdx := currColIdx - 2*diffCol

// 	count += setAntiNode(grid, antiNode1RowIdx, antiNode1ColIdx)
// 	count += setAntiNode(grid, antiNode2RowIdx, antiNode2ColIdx)

// 	return count
// }

// func setAntiNode(grid [][]string, row, col int) (count int) {
// 	// Early return if out of bounds
// 	if 0 > row || row >= len(grid) {
// 		return count
// 	}

// 	if 0 > col || col >= len(grid[0]) {
// 		return count
// 	}

// 	// Early return if non empty cell
// 	if grid[row][col] != "." {
// 		return count
// 	}

// 	count += 1

// 	return count
// }

var nodesWithAntiNode = map[struct {
	row int
	col int
}]bool{}

func processGrid(grid [][]string) ([][]string, int) {
	antiUnderNode := 0

	for rowIdx, row := range grid {
		for colIdx := range row {
			antennaValue := grid[rowIdx][colIdx]

			// Skip empty cells
			if antennaValue == "." || antennaValue == "#" {
				continue
			}

			grid, antiUnderNode = findAntiNodes(grid, antiUnderNode, antennaValue, rowIdx, colIdx)
		}
	}

	return grid, antiUnderNode
}

func findAntiNodes(grid [][]string, antiUnderNode int, target string, currRowIdx, currColIdx int) ([][]string, int) {
	for rowIdx, row := range grid {
		for colIdx := range row {
			// Skip self
			if rowIdx == currRowIdx && colIdx == currColIdx {
				continue
			}

			if grid[rowIdx][colIdx] == target {
				grid, antiUnderNode = addAntiNode(grid, antiUnderNode, currRowIdx, currColIdx, rowIdx, colIdx)
			}
		}
	}

	return grid, antiUnderNode
}

func addAntiNode(grid [][]string, antiUnderNode, currRowIdx, currColIdx, rowIdx, colIdx int) ([][]string, int) {
	diffRow := currRowIdx - rowIdx
	diffCol := currColIdx - colIdx

	antiNode1RowIdx := currRowIdx + diffRow
	antiNode1ColIdx := currColIdx + diffCol

	antiNode2RowIdx := currRowIdx - 2*diffRow
	antiNode2ColIdx := currColIdx - 2*diffCol

	grid, antiUnderNode = setAntiNode(grid, antiUnderNode, antiNode1RowIdx, antiNode1ColIdx)
	grid, antiUnderNode = setAntiNode(grid, antiUnderNode, antiNode2RowIdx, antiNode2ColIdx)

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
		nodesWithAntiNode[struct {
			row int
			col int
		}{row: row, col: col}] = true
		return grid, antiUnderNode
	}

	grid[row][col] = "#"

	return grid, antiUnderNode
}
