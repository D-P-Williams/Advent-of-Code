package main

import (
	"fmt"
	"math"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

func main() {
	lines := aoc.ReadLines("./input.txt")
	// lines := aoc.ReadLines("./example.txt")

	countPt1 := 0
	countPt2 := 0

	startPos := 50

	countPt1 = spinDial(lines, startPos)

	fmt.Println("part 1", countPt1)

	countPt2 = spinDialPt2(lines, startPos)

	fmt.Println("part 2", countPt2)
}

func spinDial(lines []string, startPos int) int {
	dialPos := startPos
	zeroCounts := 0

	for _, line := range lines {
		dirIdent := line[0]
		qty := aoc.ParseInt(line[1:])

		if dirIdent == 'R' {
			dialPos = mod(dialPos + qty)
		} else {
			dialPos = mod(dialPos - qty)
		}

		if dialPos == 0 {
			zeroCounts++
		}
	}

	return zeroCounts
}

func spinDialPt2(lines []string, startPos int) int {
	dialPos := startPos
	zeroCounts := 0

	for _, line := range lines {
		dirIdent := line[0]
		qty := aoc.ParseInt(line[1:])

		var newDialPos int
		if dirIdent == 'R' {
			// Guaranteed crosses every 100 qty spun, i.e. 501 -> 5 crosses and a potential extra one
			zeroCounts += int(math.Abs(float64(qty)) / 100)

			// See what's left and if it crosses zero again. e.g. 1 in example above
			remaining := mod(qty)

			newDialPos = dialPos + remaining

			// If new pos is outside dial range, we crossed zero again
			if 0 > newDialPos || newDialPos >= 100 {
				zeroCounts++
				newDialPos = mod(newDialPos)
			}
		} else {
			// Guaranteed crosses every 100 qty spun, i.e. 501 -> 5 crosses and a potential extra one
			zeroCounts += int(math.Abs(float64(qty)) / 100)

			// See what's left and if it crosses zero again. e.g. 1 in example above
			remaining := mod(qty)

			newDialPos = dialPos - remaining

			// If new pos is outside dial range, we crossed zero again
			if 0 > newDialPos || newDialPos >= 100 {
				// If the dial started at zero, we already counted this crossing
				if dialPos != 0 {
					zeroCounts++
				}
				newDialPos = mod(newDialPos)
			}

			if newDialPos == 0 {
				zeroCounts++
			}
		}

		dialPos = newDialPos

		// fmt.Println("line:", line, "newDialPos:", newDialPos, "dialPos:", dialPos, zeroCounts)
	}

	return zeroCounts
}

func mod(a int) int {
	return (a%100 + 100) % 100
}
