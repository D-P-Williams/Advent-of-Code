package main

import (
	"fmt"
	"math"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type coord struct {
	X int
	Y int
}

type game struct {
	// Additions to X and Y, not actual coordinates
	A coord

	// Additions to X and Y, not actual coordinates
	B coord

	// Coordinates of prize
	Prize coord
}

type presses struct {
	A int
	B int
}

const (
	partTwoOffset = 10_000_000_000_000
	ACost         = 3
	BCost         = 1
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	games := parseLines(lines)

	for _, game := range games {
		// Should use linear algebra approach for both parts.
		// But kept brute force approach for posterity.
		possibleMatches := processGameBruteForce(game)

		if len(possibleMatches) == 0 {
			continue
		}

		minCost := 1_000_000 // Very large initial score

		for cost := range possibleMatches {
			if cost < minCost {
				minCost = cost
			}
		}

		countPt1 += minCost
	}

	fmt.Println("part 1", countPt1)

	for _, game := range games {
		offsetGame := game
		offsetGame.Prize.X += partTwoOffset
		offsetGame.Prize.Y += partTwoOffset

		cost := processGameLinearAlgebra(offsetGame)

		if cost == 0 {
			continue
		}

		countPt2 += cost
	}

	fmt.Println("part 2", countPt2)
}

func parseLines(lines []string) []game {
	games := make([]game, 0, (len(lines) / 3))

	for i := 0; i < len(lines)-1; i += 4 {
		a := parseLine(lines[i], "+")
		b := parseLine(lines[i+1], "+")
		prize := parseLine(lines[i+2], "=")

		games = append(games, game{
			A:     a,
			B:     b,
			Prize: prize,
		})
	}

	return games
}

/*
Split the input string based on `initialSeparator`, then split the second element of the result by `,`.

e.g. "Prize: X=8400, Y=5400" + split on "="

Result: ["Prize: X", "8400, Y", "5400"]

split on ","

Result: ["Prize: X", "8400", " Y", "5400"]

Elements 1 and 3 are used for the `coord` object
*/
func parseLine(line, initialSeparator string) coord {
	parts := strings.Split(line, initialSeparator)

	// Split the X coord section again on , to isolate the number
	xString := strings.Split(parts[1], ",")[0]

	// Y coord is already isolated
	yString := parts[2]

	x := aoc.ParseInt(xString)
	y := aoc.ParseInt(yString)

	return coord{
		X: x,
		Y: y,
	}
}

// Initial solution - works for part 1, impractical for part 2 as takes too long
func processGameBruteForce(targetGame game) map[int]presses {
	// Map of costs with associated
	possibleCombinations := map[int]presses{}

	// Find mind of x or y prize divided by incrementor, or 100 as max value
	aLoop := int(math.Min(
		math.Ceil(float64(targetGame.Prize.X)/float64(targetGame.A.X)),
		math.Ceil(float64(targetGame.Prize.Y)/float64(targetGame.A.Y)),
	))

	bLoop := int(math.Min(
		math.Ceil(float64(targetGame.Prize.X)/float64(targetGame.B.X)),
		math.Ceil(float64(targetGame.Prize.Y)/float64(targetGame.B.Y)),
	))

	for aIdx := range aLoop {
		aXValue := targetGame.A.X * aIdx
		aYValue := targetGame.A.Y * aIdx
		for bIdx := range bLoop {
			bXValue := targetGame.B.X * bIdx
			bYValue := targetGame.B.Y * bIdx

			if aXValue+bXValue == targetGame.Prize.X && aYValue+bYValue == targetGame.Prize.Y {
				cost := aIdx*ACost + bIdx*BCost

				possibleCombinations[cost] = presses{aIdx, bIdx}
			}
		}
	}

	return possibleCombinations
}

func processGameLinearAlgebra(game game) int {
	// Doing algebraic rearranging of the equations:
	// a*x_A + b*x_B = prize_X
	// a*y_A + b*y_B = prize_Y
	// Two equations for a and b can be found.
	// Use these equations to check for solution, and thus cost.

	aNumerator := game.Prize.X*game.B.Y - game.Prize.Y*game.B.X
	aDenominator := game.A.X*game.B.Y - game.A.Y*game.B.X

	if aDenominator != 0 && aNumerator%aDenominator == 0 {
		a := aNumerator / aDenominator

		bNumerator := game.Prize.X - a*game.A.X

		if bNumerator%game.B.X == 0 {
			b := bNumerator / game.B.X

			cost := a*ACost + b*BCost

			return cost
		}
	}

	// Catch the case where a single solution exists despite A.x being 0
	if game.A.X == 0 && game.Prize.X%game.B.X == 0 {
		b := game.Prize.X / game.B.X

		aNumerator := game.Prize.Y - b*game.B.Y

		if game.A.Y != 0 && aNumerator%game.A.Y == 0 {
			a := aNumerator / game.A.Y

			cost := a*ACost + b*BCost

			return cost
		}
	}

	return 0
}
