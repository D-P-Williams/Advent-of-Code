package main

import (
	"fmt"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

func main() {
	lines := aoc.ReadGrid("./input.txt")
	// lines := aoc.ReadGrid("./example.txt")

	countPt2 := 0

	_, countPt1 := searchGrid(lines)

	fmt.Println("part 1", countPt1)

	var count int
	countPt2 += countPt1

	// Loop until no more can be removed
	for {
		lines, count = searchGrid(lines)
		if count == 0 {
			break
		}
		countPt2 += count
	}

	fmt.Println("part 2", countPt2)
}

func searchGrid(grid [][]string) ([][]string, int) {
	validSpots := 0

	matches := [][]int{}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if checkPosition(grid, row, col) {
				validSpots++
				matches = append(matches, []int{row, col})
			}
		}
	}

	for _, match := range matches {
		grid[match[0]][match[1]] = "."
	}

	return grid, validSpots
}

func checkPosition(grid [][]string, row, col int) bool {
	paperCount := 0

	if grid[row][col] != "@" {
		return false
	}

	if row > 0 && col > 0 && grid[row-1][col-1] == "@" {
		paperCount++
	}

	if row > 0 && grid[row-1][col] == "@" {
		paperCount++
	}

	if row > 0 && col < len(grid[0])-1 && grid[row-1][col+1] == "@" {
		paperCount++
	}

	if col > 0 && grid[row][col-1] == "@" {
		paperCount++
	}

	if col < len(grid[0])-1 && grid[row][col+1] == "@" {
		paperCount++
	}

	if row < len(grid)-1 && col > 0 && grid[row+1][col-1] == "@" {
		paperCount++
	}

	if row < len(grid)-1 && grid[row+1][col] == "@" {
		paperCount++
	}

	if row < len(grid)-1 && col < len(grid[0])-1 && grid[row+1][col+1] == "@" {
		paperCount++
	}

	return paperCount < 4
}
