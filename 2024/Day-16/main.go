package main

import (
	"fmt"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type point struct {
	x int
	y int
}

// type Reindeer struct {
// 	x       int
// 	y       int
// 	facing  facing
// 	score   int
// 	visited []point
// }

func turnLeft(dir int) int {
	dir -= 1

	if dir == -1 {
		dir = 3
	}

	return dir
}

func turnRight(dir int) int {
	dir += 1

	if dir == 4 {
		dir = 0
	}

	return dir
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var directions = map[int]point{
	UP:    {0, -1},
	RIGHT: {1, 0},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
}

func main() {
	grid := aoc.ReadGrid("./input.txt")

	countPt2 := 0

	startY := slices.IndexFunc(grid, func(row []string) bool {
		return slices.Contains(row, "S")
	})

	startX := slices.Index(grid[startY], "S")

	endY := slices.IndexFunc(grid, func(row []string) bool {
		return slices.Contains(row, "E")
	})

	endX := slices.Index(grid[endY], "E")

	// Create start node
	mazeHeadNode := node{
		x:      startX,
		y:      startY,
		facing: RIGHT,
		score:  0,
	}

	buildMazeTree(grid, &mazeHeadNode)

	walkTree(&mazeHeadNode, endX, endY)

	// Sort scores to get smallest first
	slices.Sort(scores)

	fmt.Println("part 1", scores[0])

	fmt.Println("part 2", countPt2)
}

// func simulateReindeer(grid [][]string, reindeer Reindeer) {
// 	// If at the end
// 	if grid[reindeer.y][reindeer.x] == "E" {
// 		fmt.Println("found end")
// 		if reindeer.score < lowestScore {
// 			fmt.Println("updating lowest score: ", lowestScore, reindeer.score)
// 			lowestScore = reindeer.score
// 		}

// 		return
// 	}

// 	// Early exit if reindeer score is above the current lowest
// 	if reindeer.score > lowestScore {
// 		return
// 	}

// 	// Map of points and incurred score for action
// 	directionsToTry := map[point]int{
// 		directions[reindeer.facing]:             0,
// 		directions[reindeer.facing.turnLeft()]:  1000,
// 		directions[reindeer.facing.turnRight()]: 1000,
// 	}

// 	for direction, penality := range directionsToTry {
// 		if alreadyVisited := slices.ContainsFunc(reindeer.visited, func(cell point) bool {
// 			return cell.x == reindeer.x && cell.y == reindeer.y
// 		}); alreadyVisited {
// 			return
// 		}

// 		if grid[reindeer.y+direction.y][reindeer.x+direction.x] == "." {
// 			// Go copies non-primitive objects by reference
// 			// Implement a basic deep copy to get a new object
// 			localReindeer := Reindeer{
// 				x:       reindeer.x,
// 				y:       reindeer.y,
// 				facing:  reindeer.facing,
// 				score:   reindeer.score,
// 				visited: reindeer.visited,
// 			}

// 			// Increment score and move reindeer
// 			localReindeer.score += 1 + penality
// 			localReindeer.x += direction.x
// 			localReindeer.y += direction.y

// 			// Record the new tile in the visited list
// 			localReindeer.visited = append(localReindeer.visited, point{localReindeer.y, localReindeer.x})

// 			// fmt.Println(reindeer, localReindeer)
// 			// Recurse movement
// 			simulateReindeer(grid, localReindeer)

// 		}
// 	}

// }

/*

- Build tree of maze routes, linked node type thing
	- node {left, straight, right}

- Walk whole tree, filter to only paths to the end

	- Add scores to secondary slice

- Sort slice to get lowest val

*/

type node struct {
	x       int
	y       int
	facing  int
	score   int
	visited []point

	left     *node
	straight *node
	right    *node
}

func buildMazeTree(grid [][]string, parent *node) {
	if alreadyVisited := slices.ContainsFunc(parent.visited, func(cell point) bool {
		return cell.x == parent.x && cell.y == parent.y
	}); alreadyVisited {
		return
	}

	// Right
	direction := directions[turnRight(parent.facing)]
	if grid[parent.y+direction.y][parent.x+direction.x] == "." || grid[parent.y+direction.y][parent.x+direction.x] == "E" {
		parent.right = &node{
			x:      parent.x + direction.x,
			y:      parent.y + direction.y,
			score:  parent.score + 1001,
			facing: turnRight(parent.facing),
		}

		parent.right.visited = parent.visited

		parent.right.visited = append(
			parent.right.visited,
			point{
				x: parent.x,
				y: parent.y,
			})

		buildMazeTree(grid, parent.right)
	}

	// Straight
	direction = directions[parent.facing]
	if grid[parent.y+direction.y][parent.x+direction.x] == "." || grid[parent.y+direction.y][parent.x+direction.x] == "E" {
		parent.straight = &node{
			x:      parent.x + direction.x,
			y:      parent.y + direction.y,
			score:  parent.score + 1,
			facing: parent.facing,
		}

		parent.straight.visited = parent.visited

		parent.straight.visited = append(
			parent.straight.visited,
			point{
				x: parent.x,
				y: parent.y,
			})

		buildMazeTree(grid, parent.straight)
	}

	// Left
	direction = directions[turnLeft(parent.facing)]
	if grid[parent.y+direction.y][parent.x+direction.x] == "." || grid[parent.y+direction.y][parent.x+direction.x] == "E" {
		parent.left = &node{
			x:      parent.x + direction.x,
			y:      parent.y + direction.y,
			score:  parent.score + 1001,
			facing: turnLeft(parent.facing),
		}

		parent.left.visited = parent.visited

		parent.left.visited = append(
			parent.left.visited,
			point{
				x: parent.x,
				y: parent.y,
			})

		buildMazeTree(grid, parent.left)
	}

}

var scores = []int{}

var lowest = 1_000_000_000

func walkTree(parent *node, endX, endY int) {
	if parent.score > lowest {
		return
	}

	if parent.right != nil {
		walkTree(parent.right, endX, endY)
	} else {
		if parent.x == endX && parent.y == endY {
			scores = append(scores, parent.score)

			if parent.score < lowest {
				lowest = parent.score
			}
		}
	}

	if parent.straight != nil {
		walkTree(parent.straight, endX, endY)
	} else {
		if parent.x == endX && parent.y == endY {
			scores = append(scores, parent.score)

			if parent.score < lowest {
				lowest = parent.score
			}
		}
	}

	if parent.left != nil {
		walkTree(parent.left, endX, endY)
	} else {
		if parent.x == endX && parent.y == endY {
			scores = append(scores, parent.score)

			if parent.score < lowest {
				lowest = parent.score
			}
		}
	}
}
