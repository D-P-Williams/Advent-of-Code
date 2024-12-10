package main

import (
	"fmt"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type coord struct {
	row int
	col int
}

var routes = map[coord]map[coord]int{}

func main() {
	grid := aoc.ReadIntGrid("./input.txt")

	countPt1 := 0
	countPt2 := 0

	processGrid(grid)

	for _, trails := range routes {
		countPt1 += len(trails)

		for _, route := range trails {
			countPt2 += route
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func processGrid(grid [][]int) {
	for rowIdx, row := range grid {
		for colIdx, cell := range row {
			if cell != 0 {
				continue
			}

			findRoutes(grid, rowIdx, colIdx, rowIdx, colIdx, 0)
		}
	}
}

func findRoutes(grid [][]int, startRowIdx, startColIdx, currRowIdx, currColIdx, currVal int) {
	if grid[currRowIdx][currColIdx] == 9 {
		recordRoute(startRowIdx, startColIdx, currRowIdx, currColIdx)
		return
	}

	// Up
	if currRowIdx > 0 && grid[currRowIdx-1][currColIdx] == currVal+1 {
		findRoutes(grid, startRowIdx, startColIdx, currRowIdx-1, currColIdx, currVal+1)
	}

	// Down
	if currRowIdx < len(grid)-1 && grid[currRowIdx+1][currColIdx] == currVal+1 {
		findRoutes(grid, startRowIdx, startColIdx, currRowIdx+1, currColIdx, currVal+1)
	}

	// Left
	if currColIdx > 0 && grid[currRowIdx][currColIdx-1] == currVal+1 {
		findRoutes(grid, startRowIdx, startColIdx, currRowIdx, currColIdx-1, currVal+1)
	}

	// Right
	if currColIdx < len(grid[0])-1 && grid[currRowIdx][currColIdx+1] == currVal+1 {
		findRoutes(grid, startRowIdx, startColIdx, currRowIdx, currColIdx+1, currVal+1)
	}
}

func recordRoute(startRowIdx, startColIdx, endRowIdx, endColIdx int) {
	if routes[coord{row: startRowIdx, col: startColIdx}] == nil {
		routes[coord{row: startRowIdx, col: startColIdx}] = map[coord]int{}
	}

	routes[coord{row: startRowIdx, col: startColIdx}][coord{row: endRowIdx, col: endColIdx}] += 1
}
