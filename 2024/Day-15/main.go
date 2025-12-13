package main

import (
	"bytes"
	"fmt"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type point struct {
	X int
	Y int
}

type Cell struct {
	x       int
	y       int
	isWall  bool
	isBox   bool
	isRobot bool
}

var directions = map[rune]point{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	blankLineIdx := slices.Index(lines, "")

	floorString := lines[:blankLineIdx]
	instructionsString := lines[blankLineIdx+1:]

	grid, robotX, robotY := parseFloor(floorString)
	instructions := parseInstructions(instructionsString)

	for _, instruction := range instructions {
		grid, robotX, robotY = processInstruction(grid, instruction, robotX, robotY)
	}

	renderGrid(grid)

	countPt1 = calculateGPSValue(grid)

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func parseFloor(lines []string) (cells [][]Cell, robotX int, robotY int) {
	cells = make([][]Cell, len(lines))

	for i, line := range lines {
		cells[i] = make([]Cell, len(line))

		for j, char := range line {
			if char == '#' {
				cells[i][j] = Cell{isWall: true, x: j, y: i}
				continue
			}

			if char == 'O' {
				cells[i][j] = Cell{isBox: true, x: j, y: i}
				continue
			}

			if char == '@' {
				robotX = j
				robotY = i
				cells[i][j] = Cell{isRobot: true, x: j, y: i}
				continue
			}

			cells[i][j] = Cell{x: j, y: i}
		}
	}

	return cells, robotX, robotY
}

// Concat together all the instructions
func parseInstructions(lines []string) string {
	var buffer bytes.Buffer

	for _, line := range lines {
		buffer.WriteString(line)
	}

	return buffer.String()
}

func processInstruction(grid [][]Cell, instruction rune, robotX, robotY int) ([][]Cell, int, int) {
	stackLength, validMove := validMove(grid, directions[instruction], robotX, robotY, 0)

	if validMove {
		grid, robotX, robotY = moveRobot(grid, directions[instruction], robotX, robotY, stackLength)
	}

	return grid, robotX, robotY
}

func validMove(grid [][]Cell, direction point, robotX, robotY int, stackLength int) (int, bool) {
	checkPosition := grid[robotY+direction.Y][robotX+direction.X]

	if !checkPosition.isBox && !checkPosition.isRobot && !checkPosition.isWall {
		if stackLength > 0 {
			return stackLength, true
		}

		return 0, true
	}

	if checkPosition.isWall {
		return -1, false
	}

	if checkPosition.isBox {
		count, valid := validMove(grid, direction, robotX+direction.X, robotY+direction.Y, stackLength+1)
		if !valid {
			return -1, false
		}

		return count, valid
	}

	if stackLength > 0 {
		return stackLength, true
	}

	return -1, false
}

func moveRobot(grid [][]Cell, direction point, robotX, robotY int, stackLength int) ([][]Cell, int, int) {
	// Left
	if direction == directions['>'] {
		// Loop backwards to avoid overwriting cells
		for i := stackLength; i >= 0; i -= 1 {
			grid[robotY][robotX+i+1] = grid[robotY][robotX+i]

			// Shift the x coord of the target cell
			grid[robotY][robotX+i+1].x = robotX + i + 1
		}

		// Set original robot position to an empty cell
		grid[robotY][robotX] = Cell{}

		return grid, robotX + 1, robotY
	}

	// Right
	if direction == directions['<'] {
		for i := stackLength; i >= 0; i -= 1 {
			grid[robotY][robotX-i-1] = grid[robotY][robotX-i]

			grid[robotY][robotX-i-1].x = robotX - i - 1
		}

		grid[robotY][robotX] = Cell{}

		return grid, robotX - 1, robotY
	}

	// Up
	if direction == directions['^'] {
		for i := stackLength; i >= 0; i -= 1 {
			grid[robotY-i-1][robotX] = grid[robotY-i][robotX]

			grid[robotY-i-1][robotX].y = robotY - i - 1
		}

		grid[robotY][robotX] = Cell{}

		return grid, robotX, robotY - 1
	}

	// Down
	if direction == directions['v'] {
		for i := stackLength; i >= 0; i -= 1 {
			grid[robotY+i+1][robotX] = grid[robotY+i][robotX]

			grid[robotY+i+1][robotX].y = robotY + i + 1
		}

		grid[robotY][robotX] = Cell{}

		return grid, robotX, robotY + 1
	}

	return grid, robotX, robotY
}

func renderGrid(grid [][]Cell) {
	for _, row := range grid {
		for _, cell := range row {
			if cell.isWall {
				fmt.Print("#")
				continue
			}

			if cell.isBox {
				fmt.Print("O")
				continue
			}

			if cell.isRobot {
				fmt.Print("@")
				continue
			}

			fmt.Print(".")
		}

		fmt.Print("\n")
	}
}

func calculateGPSValue(grid [][]Cell) int {
	gpsTotal := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell.isBox {
				gpsTotal += 100*cell.y + cell.x
			}
		}
	}

	return gpsTotal
}
