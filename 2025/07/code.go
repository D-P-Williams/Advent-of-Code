package main

import (
	"fmt"
	"math"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

/*

Initially tried making a tree for part 2. Something was wrong with it, wasn't adding enough nodes.

Looked some solutions and found method to loop grid and keep track for each row instead (gridSum func)

Kept old code for reference

*/

// type split struct {
// 	row   string
// 	col   string
// 	left  *split
// 	right *split
// }

var beamSplits map[string]bool

var splitCount = 0

// var countPt2 = 0

func main() {
	grid := aoc.ReadGrid("./input.txt")
	// grid := aoc.ReadGrid("./example.txt")

	beamSplits = map[string]bool{}
	countPt1 := 0

	startIdx := slices.Index(grid[0], "S")
	grid[1][startIdx] = "|"

	propagateDown(&grid, 1, startIdx)

	countPt1 = countBeams()

	// aoc.PrintSlices(grid)

	fmt.Println("part 1", countPt1)

	// tree := splitsToTree(beamSplits)

	// for k, v := range tree {
	// 	fmt.Printf("%s - %+v\n", k, v)
	// }

	// walkTree(tree, 1, startIdx)

	// fmt.Println("part 2", countPt2)

	fmt.Println("part 2", gridSum(grid, startIdx), "part 1 from grid solve", splitCount)
}

func propagateDown(grid *[][]string, row, col int) {
	cell := (*grid)[row][col]

	if row == len(*grid)-1 {
		return
	}

	if cell == "|" || cell == "S" {

		if (*grid)[row+1][col] == "^" {
			key := fmt.Sprintf("%d-%d", row, col)

			// Skip any splits we've already checked
			if beamSplits[key] {
				return
			}

			beamSplits[key] = true

			if col < len((*grid)[0])-1 {
				(*grid)[row+1][col+1] = "|"
				propagateDown(grid, row+1, col+1)
			}

			if col > 0 {
				(*grid)[row+1][col-1] = "|"
				propagateDown(grid, row+1, col-1)
			}

			return
		}

		if row == len(*grid)-1 {
			(*grid)[row][col] = "|"
			return
		}

		(*grid)[row+1][col] = "|"
		propagateDown(grid, row+1, col)
	}

}

func countBeams() int {
	count := 0

	for range beamSplits {
		count++
	}

	return count
}

// func splitsToTree(beamSplits map[string]bool) map[string]*split {
// 	tree := map[string]*split{}

// 	// Pre-populate all nodes
// 	for key := range beamSplits {
// 		coordParts := strings.Split(key, "-")

// 		splitRow := coordParts[0]
// 		splitCol := coordParts[1]

// 		self := split{
// 			row: splitRow,
// 			col: splitCol,
// 		}

// 		tree[key] = &self
// 	}

// 	// Link nodes together
// 	for key := range beamSplits {
// 		coordParts := strings.Split(key, "-")

// 		splitRow := coordParts[0]
// 		splitCol := coordParts[1]

// 		self := tree[key]

// 		// Check Left
// 		leftOffset := 2
// 		for {
// 			if aoc.ParseInt(splitRow)-leftOffset <= 0 {
// 				break
// 			}

// 			if leftOffset > 2 {
// 				_ = 1
// 			}

// 			potentialParentKey := fmt.Sprintf("%d-%d", aoc.ParseInt(splitRow)-leftOffset, aoc.ParseInt(splitCol)+1)
// 			parent, foundLeft := tree[potentialParentKey]
// 			if foundLeft {
// 				parent.left = self
// 				break
// 			}

// 			leftOffset += 2
// 		}

// 		// Check Right
// 		rightOffset := 2
// 		for {
// 			if aoc.ParseInt(splitRow)-rightOffset <= 0 {
// 				break
// 			}

// 			if rightOffset > 2 {
// 				_ = 1
// 			}

// 			potentialParentKey := fmt.Sprintf("%d-%d", aoc.ParseInt(splitRow)-rightOffset, aoc.ParseInt(splitCol)-1)
// 			parent, foundRight := tree[potentialParentKey]
// 			if foundRight {
// 				parent.right = self
// 				break
// 			}

// 			rightOffset += 2
// 		}

// 	}

// 	return tree
// }

// func walkTree(tree map[string]*split, startRow, startCol int) {
// 	rootKey := fmt.Sprintf("%d-%d", startRow, startCol)

// 	root := tree[rootKey]

// 	traverseTree(root)
// }

// func traverseTree(curr *split) {
// 	if curr.left != nil {
// 		traverseTree(curr.left)
// 	}

// 	if curr.right != nil {
// 		traverseTree(curr.right)
// 	}

// 	// If no left or right nodes, must be at the bottom so increment a potential path
// 	if curr.left == nil && curr.right == nil {
// 		countPt2++
// 	}
// }

func gridSum(grid [][]string, startIdx int) int {
	tachyonRow := make([]int, len(grid[0]))

	tachyonRow[startIdx] = 1

	for rowIdx := 2; rowIdx < len(grid)-1; rowIdx += 2 {
		for colIdx, cell := range grid[rowIdx] {
			// If we're processing a splitter, and there is a tachyon above
			if cell == "^" && tachyonRow[colIdx] >= 1 {
				splitCount++

				tachyonRow[colIdx+1] += int(math.Max(float64(tachyonRow[colIdx]), 1.0))
				tachyonRow[colIdx-1] += int(math.Max(float64(tachyonRow[colIdx]), 1.0))
				tachyonRow[colIdx] = 0
			}
		}
	}

	count := 0

	for _, row := range tachyonRow {
		count += row
	}

	return count
}
