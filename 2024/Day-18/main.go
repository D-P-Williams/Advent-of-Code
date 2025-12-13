package main

import (
	"fmt"
	"strings"
	"time"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

const width = 71
const height = 71

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0

	startTime := time.Now()

	grid := populateGrid(lines, 1024)
	renderGrid(grid)

	graph := gridToGraph(grid, func(s string) bool { return s != "#" })

	path := bfs(graph, aoc.Point{X: 0, Y: 0}, aoc.Point{X: width - 1, Y: height - 1})

	countPt1 = len(path)

	part1Time := time.Since(startTime)

	fmt.Printf("part 1: %d in %d μs\n", countPt1, part1Time.Microseconds())

	startTime = time.Now()

	// Know up to 1024 is valid from part 1
	bytes := 1025

	for {
		grid = populateGrid(lines, bytes)

		graph = gridToGraph(grid, func(s string) bool { return s != "#" })

		path = bfs(graph, aoc.Point{X: 0, Y: 0}, aoc.Point{X: width - 1, Y: height - 1})

		if len(path) == 0 {
			break
		}

		bytes++
	}

	part2BFSTime := time.Since(startTime)

	fmt.Printf("part 2 BFS in %d μs\n", part2BFSTime.Microseconds())

	startTime = time.Now()

	binarySearchPart2(lines)

	part2BinarySearchBFSTime := time.Since(startTime)

	fmt.Printf("part 2 binary search in %d μs\n", part2BinarySearchBFSTime.Microseconds())

	fmt.Printf("binary search %dx faster than just BFS\n", part2BFSTime.Nanoseconds()/part2BinarySearchBFSTime.Nanoseconds())

	// Bytes incremented before checking, so need to decrement for answer
	fmt.Printf("part 2: %+v\n", lines[bytes-1])
}

func populateGrid(lines []string, limit int) [][]string {
	grid := initialiseGrid()

	for i, line := range lines {
		if i == limit {
			break
		}

		parts := strings.Split(line, ",")
		x := aoc.ParseInt(parts[0])
		y := aoc.ParseInt(parts[1])

		grid[y][x] = "#"
	}

	return grid
}

func initialiseGrid() [][]string {
	grid := make([][]string, height)

	for i := range grid {
		if grid[i] == nil {
			grid[i] = make([]string, width)
		}

		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	return grid
}

func renderGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func gridToGraph(grid [][]string, isValid func(string) bool) map[aoc.Point][]aoc.Point {
	graph := make(map[aoc.Point][]aoc.Point)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "#" {
				continue
			}

			p := aoc.Point{X: x, Y: y}

			if y > 0 && isValid(grid[y-1][x]) {
				graph[p] = append(graph[p], aoc.Point{X: x, Y: y - 1})
			}

			if y < height-1 && isValid(grid[y+1][x]) {
				graph[p] = append(graph[p], aoc.Point{X: x, Y: y + 1})
			}

			if x > 0 && isValid(grid[y][x-1]) {
				graph[p] = append(graph[p], aoc.Point{X: x - 1, Y: y})
			}

			if x < width-1 && isValid(grid[y][x+1]) {
				graph[p] = append(graph[p], aoc.Point{X: x + 1, Y: y})
			}
		}
	}

	return graph
}

func bfs(graph map[aoc.Point][]aoc.Point, start aoc.Point, end aoc.Point) []aoc.Point {
	queue := []aoc.Point{start}
	visited := make(map[aoc.Point]bool)
	paths := make(map[aoc.Point][]aoc.Point)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == end {
			break
		}

		if visited[p] {
			continue
		}

		visited[p] = true

		for _, neighbour := range graph[p] {
			if !visited[neighbour] {
				queue = append(queue, neighbour)

				paths[neighbour] = append(paths[p], neighbour)
			}
		}
	}

	return paths[end]
}

func binarySearchPart2(lines []string) int {
	low := 0
	high := len(lines)

	for low < high {
		mid := (low + high) / 2

		grid := populateGrid(lines, mid)

		graph := gridToGraph(grid, func(s string) bool { return s != "#" })

		path := bfs(graph, aoc.Point{X: 0, Y: 0}, aoc.Point{X: width - 1, Y: height - 1})

		if len(path) == 0 {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
