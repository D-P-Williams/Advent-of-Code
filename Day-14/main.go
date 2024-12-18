package main

import (
	"fmt"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type point struct {
	X int
	Y int
}

type Robot struct {
	position point
	velocity point
}

const (
	Width  int = 101
	Height int = 103
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	robots := parseLines(lines)

	for range 100 {
		robots = tickSecond(robots)
	}

	// renderGrid(buildGrid(robots))

	countPt1 = countQuadrants(buildGrid(robots))

	// Similar process to part 1, but after each iteration check number of collisions,
	// if none, render output to check output
	robots = parseLines(lines)

	for i := range 10000 {
		robots = tickSecond(robots)
		collisions := checkCollisions(robots)

		if collisions == 0 {
			// Render grid as visual check for tree
			renderGrid(buildGrid(robots))

			// i + 1 as loop starts at 1 second but 0 index
			countPt2 = i + 1
			break
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func parseLines(lines []string) []Robot {
	robots := make([]Robot, 0, len(lines))

	for _, line := range lines {
		robots = append(robots, parseLine(line))
	}

	return robots
}

func parseLine(line string) Robot {
	// "p=0,4 v=3,-3"
	// Split on space to separate p and v
	parts := strings.Split(line, " ")

	numbers := []string{}

	// Split on = to isolate numbers
	for _, part := range parts {
		numbers = append(numbers, strings.Split(part, "=")[1])
	}

	coords := []string{}

	// Split on , to isolate each number
	for _, number := range numbers {
		coords = append(coords, strings.Split(number, ",")...)
	}

	return Robot{
		position: point{
			X: aoc.ParseInt(coords[0]),
			Y: aoc.ParseInt(coords[1]),
		},
		velocity: point{
			X: aoc.ParseInt(coords[2]),
			Y: aoc.ParseInt(coords[3]),
		},
	}
}

func buildGrid(robots []Robot) [][][]Robot {
	grid := make([][][]Robot, Height)

	for i := range grid {
		grid[i] = make([][]Robot, Width)

		for j := range grid[i] {
			grid[i][j] = []Robot{}
		}
	}

	for _, robot := range robots {
		grid[robot.position.Y][robot.position.X] = append(grid[robot.position.Y][robot.position.X], robot)
	}

	return grid
}

// Console log the current state of the grid
func renderGrid(grid [][][]Robot) {
	fmt.Println()

	for _, row := range grid {
		for _, col := range row {
			if len(col) == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(len(col))
			}
		}
		fmt.Print("\n")
	}
}

func tickSecond(robots []Robot) []Robot {
	for i, robot := range robots {
		newX := robot.position.X + robot.velocity.X
		newY := robot.position.Y + robot.velocity.Y

		// Wrap x and y if over bounds
		if newX >= Width {
			newX -= Width
		}

		if newX < 0 {
			newX += Width
		}

		if newY >= Height {
			newY -= Height
		}

		if newY < 0 {
			newY += Height
		}

		robots[i].position.X = newX
		robots[i].position.Y = newY
	}

	return robots
}

// [UL, UR, BL, BR]
func countQuadrants(grid [][][]Robot) int {
	counts := make([]int, 0, 4)

	quadrants := make([][][][]Robot, 0, 4)

	// Upper Left
	rowSplit := grid[:Height/2]
	colSplit := [][][]Robot{}

	for _, row := range rowSplit {
		colSplit = append(colSplit, row[:Width/2])
	}

	quadrants = append(quadrants, colSplit)

	// Upper Left
	rowSplit = grid[:Height/2]
	colSplit = [][][]Robot{}

	for _, row := range rowSplit {
		colSplit = append(colSplit, row[(Width/2)+1:])
	}

	quadrants = append(quadrants, colSplit)

	// Bottom Left
	rowSplit = grid[(Height/2)+1:]
	colSplit = [][][]Robot{}

	for _, row := range rowSplit {
		colSplit = append(colSplit, row[:Width/2])
	}

	quadrants = append(quadrants, colSplit)

	// Bottom Right
	rowSplit = grid[(Height/2)+1:]
	colSplit = [][][]Robot{}

	for _, row := range rowSplit {
		colSplit = append(colSplit, row[(Width/2)+1:])
	}

	quadrants = append(quadrants, colSplit)

	for _, quadrant := range quadrants {
		quadrantCount := 0

		for _, row := range quadrant {
			for _, cell := range row {
				quadrantCount += len(cell)
			}
		}

		if quadrantCount != 0 {
			counts = append(counts, quadrantCount)
		}
	}

	safetyFactor := 1

	for _, count := range counts {
		safetyFactor *= count
	}

	return safetyFactor
}

func checkCollisions(robots []Robot) int {
	collisions := 0

	for i := 0; i < len(robots)-1; i += 1 {
		for j := 0; j < len(robots)-1; j += 1 {
			if i == j {
				continue
			}

			if robots[i].position.X == robots[j].position.X &&
				robots[i].position.Y == robots[j].position.Y {
				collisions += 1
			}
		}
	}

	return collisions
}
