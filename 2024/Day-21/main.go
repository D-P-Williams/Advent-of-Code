package main

import (
	"fmt"
	"math"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

/*
Numeric Keypad:

+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
|	| 0 | A |
+---+---+---+
*/
var numericKeypadGrid = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"", "0", "A"},
}

type NumericKeypad struct {
	currPos aoc.Point
}

func (NumericKeypad) getCoords(value string) aoc.Point {
	colIdx := slices.IndexFunc(numericKeypadGrid, func(col []string) bool {
		return slices.Contains(col, value)
	})

	rowIdx := slices.Index(numericKeypadGrid[colIdx], value)

	return aoc.Point{
		X: rowIdx,
		Y: colIdx,
	}
}

func (n *NumericKeypad) getMovesTo(value string) aoc.Point {
	targetPos := n.getCoords(value)

	return aoc.Point{
		X: targetPos.X - n.currPos.X,
		Y: targetPos.Y - n.currPos.Y,
	}
}

func (n *NumericKeypad) moveTo(value string) (aoc.Point, int) {
	toTravel := n.getMovesTo(value)

	// +1 to account for pressing A after input
	dist := manhattanDistance(toTravel) + 1

	n.currPos = n.getCoords(value)

	return toTravel, dist
}

/*
Directional Keypad

+---+---+---+
|	| ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
var directionalKeypadGrid = [][]string{
	{"", "^", "A"},
	{"<", "v", ">"},
}

type DirectionalKeypad struct {
	currPos    aoc.Point
	pathLength int
}

func (DirectionalKeypad) getCoords(value string) aoc.Point {
	colIdx := slices.IndexFunc(directionalKeypadGrid, func(col []string) bool {
		return slices.Contains(col, value)
	})

	rowIdx := slices.Index(directionalKeypadGrid[colIdx], value)

	return aoc.Point{
		X: rowIdx,
		Y: colIdx,
	}
}

func (d *DirectionalKeypad) getMovesTo(value string) aoc.Point {
	targetPos := d.getCoords(value)

	return aoc.Point{
		X: targetPos.X - d.currPos.X,
		Y: targetPos.Y - d.currPos.Y,
	}
}

func (d *DirectionalKeypad) moveTo(value string) (aoc.Point, int) {
	toTravel := d.getMovesTo(value)

	dist := manhattanDistance(toTravel)

	d.currPos = d.getCoords(value)

	return toTravel, dist
}

func (d *DirectionalKeypad) moveNumericBy(travel aoc.Point) (aoc.Point, int) {
	distTravelled := 0

	// Assume positive direction movement
	horizontalDir := ">"
	verticalDir := "v"

	// If negative dir, flip target symbol
	if travel.X < 0 {
		horizontalDir = "<"
	}

	if travel.Y < 0 {
		verticalDir = "^"
	}

	horizontalDist := 0
	verticalDist := 0

	if travel.X != 0 {
		horizontalDist = manhattanDistance(d.getMovesTo(horizontalDir))
	}

	if travel.Y != 0 {
		verticalDist = manhattanDistance(d.getMovesTo(verticalDir))
	}

	// If horizontal distance shorter, do that move first
	if verticalDist >= horizontalDist {
		dist := 0

		if travel.X != 0 {
			_, dist = d.moveTo(horizontalDir)
			distTravelled += dist + int(math.Abs(float64(travel.X)))
		}

		if travel.Y != 0 {
			_, dist = d.moveTo(verticalDir)
			distTravelled += dist + int(math.Abs(float64(travel.Y)))
		}
	} else {
		dist := 0

		if travel.Y != 0 {
			_, dist = d.moveTo(verticalDir)
			distTravelled += dist
		}

		if travel.X != 0 {
			_, dist = d.moveTo(horizontalDir)
			distTravelled += dist
		}
	}

	// // if travel.Y != 0 {
	// _, dist := d.moveTo(verticalDir)

	// // Each travel is a single press of A, add distance as number of A's
	// distTravelled += dist
	// // + int(math.Abs(float64(travel.Y)))
	// // }

	// // if travel.X != 0 {
	// _, dist = d.moveTo(horizontalDir)

	// // Each travel is a single press of A, add distance as number of A's
	// distTravelled += dist
	// // + int(math.Abs(float64(travel.X)))
	// // }

	// distTravelled += manhattanDistance(travel)

	return aoc.Point{}, distTravelled
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	numericKeypad := NumericKeypad{}
	numericKeypad.currPos = numericKeypad.getCoords("A")

	directionalKeypad := DirectionalKeypad{}
	directionalKeypad.currPos = directionalKeypad.getCoords("A")

	for _, line := range lines {
		for _, char := range line {
			toTravel, _ := numericKeypad.moveTo(string(char))

			_, dist := directionalKeypad.moveNumericBy(toTravel)

			fmt.Println(dist)

			// Must input A to confirm input
			_, dist = directionalKeypad.moveTo("A")

			fmt.Println(dist + 1)
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func manhattanDistance(diff aoc.Point) int {
	return int(math.Abs(float64(diff.X)) + math.Abs(float64(diff.Y)))
}
