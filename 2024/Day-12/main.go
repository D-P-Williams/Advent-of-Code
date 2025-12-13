package main

import (
	"fmt"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type coord struct {
	row int
	col int
}

var gardens = []map[string][]coord{}

var corners = map[coord]int{}

func main() {
	grid := aoc.ReadGrid("./input.txt")

	countPt1 := 0
	countPt2 := 0

	processGrid(grid)

	for _, garden := range gardens {
		area := 0
		perimeter := 0
		// corners := 0
		sides := 4

		// Gardens is a map[string][]coord, so for each to avoid needing to know index value.
		// Should only ever contain a single slice per map entry.
		for _, gardenCells := range garden {
			area = len(gardenCells)
			perimeter = getPerimeter(gardenCells, len(grid), len(grid[0]))

			break
		}

		countPt1 += area * perimeter

		// sides += corners * 2

		for _, corner := range corners {
			switch corner {
			case 4: // Single cell square
				sides += 4
			case 3:
				sides += 3
			case 2: // U shape. +1 as broke existing side
				sides += 3 + 1
			case 1: // Normal L corner
				sides += 2
			case 0: // No corner
				sides += 0
			}
		}

		countPt2 += area * sides

		corners = map[coord]int{}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)

	part2(grid)
}

func processGrid(grid [][]string) {
	for rowIdx, row := range grid {
		for colIdx, cell := range row {
			coords := []coord{{rowIdx, colIdx}}

			if gardensContains(coord{rowIdx, colIdx}, cell) {
				continue
			}

			coords = findGarden(grid, coords, rowIdx, colIdx, cell)

			gardens = append(gardens, map[string][]coord{cell: coords})
		}
	}
}

func findGarden(grid [][]string, coords []coord, rowIdx, colIdx int, value string) []coord {
	// Up
	if rowIdx > 0 && grid[rowIdx-1][colIdx] == value {
		if !slices.Contains(coords, coord{rowIdx - 1, colIdx}) {
			coords = append(coords, coord{rowIdx - 1, colIdx})
			coords = findGarden(grid, coords, rowIdx-1, colIdx, value)
		}
	}

	// Down
	if rowIdx < len(grid)-1 && grid[rowIdx+1][colIdx] == value {
		if !slices.Contains(coords, coord{rowIdx + 1, colIdx}) {
			coords = append(coords, coord{rowIdx + 1, colIdx})
			coords = findGarden(grid, coords, rowIdx+1, colIdx, value)
		}
	}

	// Left
	if colIdx > 0 && grid[rowIdx][colIdx-1] == value {
		if !slices.Contains(coords, coord{rowIdx, colIdx - 1}) {
			coords = append(coords, coord{rowIdx, colIdx - 1})
			coords = findGarden(grid, coords, rowIdx, colIdx-1, value)
		}
	}

	// Right
	if colIdx < len(grid[0])-1 && grid[rowIdx][colIdx+1] == value {
		if !slices.Contains(coords, coord{rowIdx, colIdx + 1}) {
			coords = append(coords, coord{rowIdx, colIdx + 1})
			coords = findGarden(grid, coords, rowIdx, colIdx+1, value)
		}
	}

	return coords
}

func gardensContains(cell coord, value string) bool {
	for _, garden := range gardens {
		if coords, ok := garden[value]; ok {
			for _, coord := range coords {
				if cell.row == coord.row && cell.col == coord.col {
					return true
				}
			}
		}
	}

	return false
}

func getPerimeter(garden []coord, gridX, gridY int) int {
	perimeter := 0
	// corners := 0

	for _, cell := range garden {
		adjacents := findAdjacent(garden, cell, gridX, gridY)
		checkCorner(garden, cell, gridX, gridY)

		switch adjacents {
		case 4:
			perimeter += 0
		case 3:
			perimeter += 1
		case 2:
			perimeter += 2
		case 1:
			perimeter += 3
		case 0:
			perimeter += 4
		}
	}

	// side perim = 4 + corners * 2 ?

	return perimeter
}

func findAdjacent(garden []coord, cell coord, gridX, gridY int) int {
	adjacents := 0

	// Up
	if cell.row > 0 {
		if slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row-1 && gardenCell.col == cell.col
		}) {
			adjacents += 1
		}
	}

	// Down
	if cell.row < gridX {
		if slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row+1 && gardenCell.col == cell.col
		}) {
			adjacents += 1
		}
	}

	// Left
	if cell.col > 0 {
		if slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col-1
		}) {
			adjacents += 1
		}
	}

	// Right
	if cell.col < gridY {
		if slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col+1
		}) {
			adjacents += 1
		}
	}

	return adjacents
}

func checkCorner(garden []coord, cell coord, gridX, gridY int) {
	// # = current, X is adjacents, . = space

	// #X
	// X.
	if cell.row < gridX && cell.col < gridY {
		hasLower := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row+1 && gardenCell.col == cell.col
		})

		hasRight := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col+1
		})

		hasDiagonal := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row+1 && gardenCell.col == cell.col+1
		})

		if hasLower && hasRight && !hasDiagonal {
			corners[coord{cell.row + 1, cell.col + 1}] += 1
			// corners += 1
		}
	}

	// X#
	// .X
	if cell.row < gridX && cell.col > 0 {
		hasLower := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row+1 && gardenCell.col == cell.col
		})

		hasLeft := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col-1
		})

		hasDiagonal := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row+1 && gardenCell.col == cell.col-1
		})

		if hasLower && hasLeft && !hasDiagonal {
			corners[coord{cell.row + 1, cell.col - 1}] += 1
		}
	}

	// X.
	// #X
	if cell.row > 0 && cell.col < gridY {
		hasUpper := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row-1 && gardenCell.col == cell.col
		})

		hasRight := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col+1
		})

		hasDiagonal := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row-1 && gardenCell.col == cell.col+1
		})

		if hasUpper && hasRight && !hasDiagonal {
			corners[coord{cell.row - 1, cell.col + 1}] += 1
		}
	}

	// .X
	// X#
	if cell.row > 0 && cell.col > 0 {
		hasUpper := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row-1 && gardenCell.col == cell.col
		})

		hasRight := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row && gardenCell.col == cell.col-1
		})

		hasDiagonal := slices.ContainsFunc(garden, func(gardenCell coord) bool {
			return gardenCell.row == cell.row-1 && gardenCell.col == cell.col-1
		})

		if hasUpper && hasRight && !hasDiagonal {
			corners[coord{cell.row - 1, cell.col - 1}] += 1
		}
	}
}

// ###########################################################################
// Not solved, above is close but doesn't handle all edge cases.
// Remaining known edge case: larger than 1x1 squares within an area, e.g.:
// 	XXXXXX
// 	XXX..X
// 	XXX..X
// 	X..XXX
// 	X..XXX
// 	XXXXXX
//
// https://github.com/OrenRosen/adventofcode/blob/main/2024/day12/main.go
// ###########################################################################

func part2(matrix [][]string) {
	ResetVisited()

	total := 0
	for x := range matrix {
		for y := range matrix[x] {
			if WasVisited(x, y) {
				continue
			}

			s := service{
				matrix: matrix,
			}

			BFSMatrix(matrix, x, y, s.ShouldTravel, s.HandleVisit, s.HandleTravelBlockedPart2)
			total += s.area * s.parameter
		}
	}

	fmt.Println("------------------ total", total)
}

type service struct {
	matrix [][]string

	area      int
	parameter int

	// only used in part2
	preventedTiles map[Point][]bool
}

func (s *service) AddToPreventedTiles(to Point, direction int) {
	if s.preventedTiles == nil {
		s.preventedTiles = make(map[Point][]bool)
	}

	if s.preventedTiles == nil {
		s.preventedTiles = make(map[Point][]bool)
	}

	if s.preventedTiles[to] == nil {
		s.preventedTiles[to] = make([]bool, 4)
	}

	s.preventedTiles[to][direction] = true
}

func (s *service) ShouldTravel(from, to Point) bool {
	fromVal := s.matrix[from.X][from.Y]
	toVal := s.matrix[to.X][to.Y]

	return fromVal == toVal
}

func (s *service) HandleVisit(p Point) {
	s.area++
}

func (s *service) HandleTravelBlockedPart1(from, to Point, becauseAlreadyVisited bool) {
	if becauseAlreadyVisited {
		// we need to add to perimeter, only if it's not the same value
		fromValue := s.matrix[from.X][from.Y]
		toValue := s.matrix[to.X][to.Y]
		if fromValue == toValue {
			return
		}
	}

	s.parameter++
}

func (s *service) HandleTravelBlockedPart2(from, to Point, becauseAlreadyVisited bool) {
	direction := getDirection(from, to)
	toValue, fromValue := "", ""

	if IsOutOfBoundsMatrix(s.matrix, to) {
		s.AddToPreventedTiles(to, direction)
	} else {
		toValue = s.matrix[to.X][to.Y]
		fromValue = s.matrix[from.X][from.Y]

		// only add to prevented if it's not the same value
		if fromValue != toValue {
			s.AddToPreventedTiles(to, direction)
		}
	}

	// if we are blocked because it was already visited, we need to add to perimeter, only if it's not the same value
	if becauseAlreadyVisited && fromValue == toValue {
		return
	}

	adjacentTiles := []Point{}
	isPreventedFromRightOrLeft := to.X == from.X
	if isPreventedFromRightOrLeft {
		adjacentTiles = []Point{
			{X: to.X + 1, Y: to.Y},
			{X: to.X - 1, Y: to.Y},
		}
	} else {
		adjacentTiles = []Point{
			{X: to.X, Y: to.Y + 1},
			{X: to.X, Y: to.Y - 1},
		}
	}

	for _, t := range adjacentTiles {
		preventedDirections, ok := s.preventedTiles[t]
		if ok && preventedDirections[direction] {
			return
		}
	}

	s.parameter++
}

const (
	UP    = 0
	DOWN  = 1
	LEFT  = 2
	RIGHT = 3
)

func getDirection(from, to Point) int {
	if from.X < to.X {
		return DOWN
	} else if from.X > to.X {
		return UP
	} else if from.Y < to.Y {
		return RIGHT
	} else if from.Y > to.Y {
		return LEFT
	}

	panic("Invalid direction")
}

func IsOutOfBoundsMatrix(matrix [][]string, p Point) bool {
	return p.X < 0 || p.Y < 0 || p.X >= len(matrix) || p.Y >= len(matrix[p.X])
}
