package main

import (
	"fmt"
	"math"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type Point struct {
	x     int
	y     int
	value string
}

var (
	height = 0
	width  = 0
)

func main() {
	gridString := aoc.ReadGrid("./input.txt")

	height = len(gridString)
	width = len(gridString[0])

	countPt1 := 0
	countPt2 := 0

	startY := slices.IndexFunc(gridString, func(s []string) bool { return slices.Contains(s, "S") })
	startX := slices.IndexFunc(gridString[startY], func(s string) bool { return s == "S" })
	startPoint := Point{x: startX, y: startY, value: gridString[startY][startX]}

	endY := slices.IndexFunc(gridString, func(s []string) bool { return slices.Contains(s, "E") })
	endX := slices.IndexFunc(gridString[endY], func(s string) bool { return s == "E" })
	endPoint := Point{x: endX, y: endY, value: gridString[endY][endX]}

	grid := parseGrid(gridString)

	initialPath := bfs(grid, startPoint, endPoint)

	timeSavings := checkCheats(grid, append([]Point{startPoint}, initialPath...), 0)
	fmt.Println(timeSavings)

	for timeSave, count := range timeSavings {
		if timeSave >= 100 {
			countPt1 += count
		}
	}

	// This function works for both parts. Kept part one as original solution.
	timeSavings2 := checkCheats2(append([]Point{startPoint}, initialPath...), 100, 20)
	fmt.Println(timeSavings2)

	for timeSave, count := range timeSavings2 {
		if timeSave >= 100 {
			countPt2 += count
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func parseGrid(gridString [][]string) [][]Point {
	grid := make([][]Point, len(gridString))

	for y, row := range gridString {
		grid[y] = make([]Point, len(row))

		for x, value := range row {
			grid[y][x] = Point{x: x, y: y, value: value}
		}
	}

	return grid
}

func getValidPoints(grid [][]Point, x, y int) []Point {
	points := make([]Point, 0, 4)

	if y > 0 && grid[y-1][x].value != "#" {
		points = append(points, Point{x: x, y: y - 1, value: grid[y-1][x].value})
	}

	if y < height-1 && grid[y+1][x].value != "#" {
		points = append(points, Point{x: x, y: y + 1, value: grid[y+1][x].value})
	}

	if x > 0 && grid[y][x-1].value != "#" {
		points = append(points, Point{x: x - 1, y: y, value: grid[y][x-1].value})
	}

	if x < width-1 && grid[y][x+1].value != "#" {
		points = append(points, Point{x: x + 1, y: y, value: grid[y][x+1].value})
	}

	return points
}

func bfs(grid [][]Point, start Point, end Point) []Point {
	queue := []Point{start}
	visited := make(map[Point]bool)
	paths := make(map[Point][]Point)

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

		for _, neighbour := range getValidPoints(grid, p.x, p.y) {
			if !visited[neighbour] {
				queue = append(queue, neighbour)

				paths[neighbour] = append(paths[p], neighbour)
			}
		}
	}

	return paths[end]
}

func checkCheats(grid [][]Point, path []Point, skip int) map[int]int {
	timeSaves := map[int]int{}

	directions := []aoc.Point{{X: -2, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: -2}, {X: 0, Y: 2}}

	for travelledIdx, point := range path[:len(path)-skip] {
		for _, direction := range directions {
			potentialX := point.x + direction.X
			potentialY := point.y + direction.Y

			// Outside grid bounds
			if potentialX < 0 || potentialX >= width || potentialY < 0 || potentialY >= height {
				continue
			}

			if grid[potentialY][potentialX].value == "#" {
				continue
			}

			// Find the cheat point in the path
			cheatIdx := slices.IndexFunc(path, func(p Point) bool {
				return p.x == potentialX && p.y == potentialY
			})

			if cheatIdx == -1 || cheatIdx < travelledIdx {
				continue
			}

			// Cut section = total minus either end
			timeSave := cheatIdx - travelledIdx - 2

			timeSaves[timeSave] += 1
		}

	}

	return timeSaves
}

// Check all points on path against all future points, any with a manhattan distance less than cheatMax are valid cheats
func checkCheats2(path []Point, skip, cheat int) map[int]int {
	timeSaves := map[int]int{}

	for i, pathPoint := range path[:len(path)-skip] {
		for j, cheatPoint := range path[i+skip:] {
			dx := math.Abs(float64(pathPoint.x - cheatPoint.x))
			dy := math.Abs(float64(pathPoint.y - cheatPoint.y))

			dist := int(dx + dy)

			// Valid cheat must be less than the regular distance to the target point and less than the cheat limit
			if dist <= cheat && dist <= j {
				// Cut section = total minus either end
				timeSave := len(path) - (j + i)

				timeSaves[timeSave] += 1
			}
		}
	}

	return timeSaves
}
