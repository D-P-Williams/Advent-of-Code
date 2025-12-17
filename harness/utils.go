package aoc

import (
	"bufio"
	"fmt"
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

// readGrid reads a whole file into memory and returns
// a slice of its lines, split into individual strings.
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

// readIntGrid reads a whole file into memory and returns
// a slice of its lines, split into individual ints.
func ReadIntGrid(path string) [][]int {
	lines := ReadLines(path)

	var grid = make([][]int, 0, len(lines))

	for _, line := range lines {
		cells := make([]int, 0, len(line))

		for _, cell := range line {
			cells = append(cells, ParseInt(string(cell)))
		}

		grid = append(grid, cells)
	}
	return grid
}

// ParseInt parses a string as an int and panics on error
func ParseInt(int string) int {
	value, err := strconv.Atoi(int)
	if err != nil {
		panic(err)
	}

	return value
}

func PrintSlices[T any](slices []T) {
	for _, slice := range slices {
		fmt.Println(slice)
	}
}
