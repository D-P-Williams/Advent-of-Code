package aoc

import (
	"bufio"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadGrid(path string) [][]string {
	lines := ReadLines(path)

	var grid = make([][]string, 0, len(lines))

	for _, line := range lines {
		cells := make([]string, 0, len(line))

		for _, cell := range line {
			cells = append(cells, string(cell))
		}

		grid = append(grid, cells)
	}
	return grid
}

func ParseInt(int string) int {
	value, err := strconv.Atoi(int)
	if err != nil {
		panic(err)
	}

	return value
}
