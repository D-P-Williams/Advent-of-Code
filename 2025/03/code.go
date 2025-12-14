package main

import (
	"fmt"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

func main() {
	lines := aoc.ReadLines("./input.txt")
	// lines := aoc.ReadLines("./example.txt")

	countPt1 := 0
	countPt2 := 0

	countPt1 = sumVoltages(lines)

	fmt.Println("part 1", countPt1)

	countPt2 = sumVoltagesWithOverride(lines, 12)

	fmt.Println("part 2", countPt2)

	// Part 2 solution also works for part 1
	countPt2 = sumVoltagesWithOverride(lines, 2)

	fmt.Println("part 1 using 2 solution", countPt2)
}

func sumVoltages(lines []string) int {
	voltageSum := 0

	for _, line := range lines {
		voltageSum += checkBank(line)
	}

	return voltageSum
}

func checkBank(line string) int {
	batteries := strings.Split(line, "")

	maxPower := 0

	for i := range batteries {
		for j := i + 1; j < len(batteries); j++ {
			comboPower := aoc.ParseInt(batteries[i] + batteries[j])

			if comboPower > maxPower {
				maxPower = comboPower
			}
		}
	}

	return maxPower
}

func sumVoltagesWithOverride(lines []string, batteryCount int) int {
	voltageSum := 0

	for _, line := range lines {
		voltageSum += checkBankWithOverride(line, batteryCount)
	}

	return voltageSum
}

func checkBankWithOverride(line string, searchFor int) int {
	batteries := []rune(line)

	maxPowerSlice := make([]rune, searchFor)

	searchOffset := len(batteries) - searchFor
	remainingDigits := batteries[:searchOffset]

	for i := range searchFor {
		// // Append the newly available digit from the right side
		remainingDigits = append(remainingDigits, batteries[i+searchOffset])

		max, maxIdx := findMaxInSlice(remainingDigits)

		// Cut out the max index value from slice
		remainingDigits = remainingDigits[maxIdx+1:]

		// Update the max power slice
		maxPowerSlice[i] = max
	}

	return aoc.ParseInt(string(maxPowerSlice))
}

func findMaxInSlice(slice []rune) (rune, int) {
	maxRune := '0'
	maxIndex := -1

	for i, r := range slice {
		if r > maxRune {
			maxRune = r
			maxIndex = i
		}
	}

	return maxRune, maxIndex
}
