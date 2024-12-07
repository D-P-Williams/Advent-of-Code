package main

import (
	"fmt"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

const GUARD = "^"
const LoopThreshold = 10000 // Brute force threshold to assume infinite loop

func main() {
	lines := aoc.ReadLines("./input.txt")

	var grid = make([][]string, 0, len(lines))

	for _, line := range lines {
		cells := make([]string, 0, len(line))

		for _, cell := range line {
			cells = append(cells, string(cell))
		}

		grid = append(grid, cells)
	}

	guardRow := slices.IndexFunc(lines, func(line string) bool {
		return strings.Contains(line, GUARD)
	})

	guardColumn := strings.Index(lines[guardRow], GUARD)

	grid, _ = guardLoop(grid, guardRow, guardColumn)

	countPt1 := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell == "X" {
				countPt1 += 1
			}
		}
	}

	fmt.Println("part 1", countPt1)

	countPt2 := 0

	for rowIdx, row := range grid {
		for colIdx, cell := range row {
			// Skip if already a blocker or guard starting point
			if cell == "#" || cell == "^" {
				continue
			}

			localGrid := make([][]string, len(grid))

			for rowIdx, row := range grid {
				localGrid[rowIdx] = make([]string, len(grid[0]))

				for colIdx := range row {
					localGrid[rowIdx][colIdx] = grid[rowIdx][colIdx]
				}
			}

			localGrid[rowIdx][colIdx] = "#"

			_, steps := guardLoop(localGrid, guardRow, guardColumn)

			if steps >= LoopThreshold {
				countPt2 += 1
			}
		}
	}

	fmt.Println("part 2", countPt2)
}

func guardLoop(grid [][]string, guardRow, guardCol int) ([][]string, int) {
	facing := 0 // UP, 90 deg per

	steps := 0

	grid[guardCol][guardRow] = "X"

	// While within the grid
	for 0 <= guardRow && guardRow < len(grid[0]) && 0 <= guardCol && guardCol < len(grid) {
		// Divide by 90 and modulo 4 to get normalised facing
		switch (facing / 90) % 4 {
		case 0: // UP
			// fmt.Println("up")

			// If move would push off grid
			if guardRow-1 < 0 {
				return grid, steps
			}

			if grid[guardRow-1][guardCol] == "#" {
				facing += 90
				continue
			}

			guardRow = guardRow - 1
			grid[guardRow][guardCol] = "X"

		case 1: // LEFT
			// fmt.Println("left")

			// If move would push off grid
			if guardCol+1 >= len(grid[0]) {
				return grid, steps
			}

			if grid[guardRow][guardCol+1] == "#" {
				facing += 90
				continue
			}

			guardCol = guardCol + 1
			grid[guardRow][guardCol] = "X"

		case 2: // DOWN
			// fmt.Println("down")

			// If move would push off grid
			if guardRow+1 >= len(grid) {
				return grid, steps
			}

			if grid[guardRow+1][guardCol] == "#" {
				facing += 90
				continue
			}

			guardRow = guardRow + 1
			grid[guardRow][guardCol] = "X"

		case 3: // RIGHT
			// fmt.Println("right")

			// If move would push off grid
			if guardCol-1 < 0 {
				return grid, steps
			}

			if grid[guardRow][guardCol-1] == "#" {
				facing += 90
				continue
			}

			guardCol = guardCol - 1
			grid[guardRow][guardCol] = "X"
		}

		steps += 1

		if steps >= LoopThreshold {
			return grid, steps
		}
	}

	return grid, steps
}
