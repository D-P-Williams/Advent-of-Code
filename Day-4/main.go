package main

import (
	"fmt"
)

func main() {
	lines := readLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	for j := 0; j < len(lines); j += 1 {
		for i := 0; i < len(lines[j]); i += 1 {
			// Search for XMAS
			countPt1 += getXMASCount(lines, i, j)

			// Search for a cross of two MAS
			countPt2 += getCrossedMASCount(lines, i, j)
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func checkMatch(test, match string) bool {
	if test == match || test == Reverse(match) {
		return true
	}

	return false
}

func getXMASCount(lines []string, i, j int) int {
	count := 0

	// Vertical
	if i < len(lines[0])-3 {
		word := lines[j][i : i+4]
		if checkMatch(word, "XMAS") {
			count++
		}
	}

	// Horizontals
	if j < len(lines)-3 {
		word := string(lines[j][i]) + string(lines[j+1][i]) + string(lines[j+2][i]) + string(lines[j+3][i])
		if checkMatch(word, "XMAS") {
			count++
		}
	}

	if i < len(lines[0])-3 && j < len(lines)-3 {
		// Diagonal Down
		word := string(lines[j][i]) + string(lines[j+1][i+1]) + string(lines[j+2][i+2]) + string(lines[j+3][i+3])
		if checkMatch(word, "XMAS") {
			count++
		}

		// Diagonal Up
		word = string(lines[j][i+3]) + string(lines[j+1][i+2]) + string(lines[j+2][i+1]) + string(lines[j+3][i+0])
		if checkMatch(word, "XMAS") {
			count++
		}
	}

	return count
}

func getCrossedMASCount(lines []string, i, j int) int {
	count := 0

	if i < len(lines[0])-2 && j < len(lines)-2 {
		diagonalA := string(lines[j][i]) + string(lines[j+1][i+1]) + string(lines[j+2][i+2])
		diagonalB := string(lines[j][i+2]) + string(lines[j+1][i+1]) + string(lines[j+2][i])

		if checkMatch(diagonalA, "MAS") {
			if checkMatch(diagonalB, "MAS") {
				count++
			}
		}
	}

	return count
}
