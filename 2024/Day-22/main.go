package main

import (
	"fmt"
	"strconv"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt2 := 0

	sellCounts := make(map[string]int, 2000)

	for _, line := range lines {
		secretNumbers, prices := generateSecretNumbers(aoc.ParseInt(line))

		sequences := calculateSequencesMap(prices)
		_ = secretNumbers

		for sequence, count := range sequences {
			sellCounts[sequence] += count
		}
	}

	maxSequenceCount := 0

	for _, count := range sellCounts {
		if count > maxSequenceCount {
			maxSequenceCount = count
		}
	}

	fmt.Println("part 1", maxSequenceCount)

	fmt.Println("part 2", countPt2)
}

func generateSecretNumbers(start int) ([]int, []int) {
	secretNumber := start
	numbers := make([]int, 0, 2000)
	prices := make([]int, 0, 2000)

	for range 2000 {
		secretNumber = generateSecretNumber(secretNumber)
		numbers = append(numbers, secretNumber)

		// Extract the final digit of the secret number
		prices = append(prices, aoc.ParseInt(string(strconv.Itoa(secretNumber)[len(strconv.Itoa(secretNumber))-1])))
	}

	return numbers, prices
}

func generateSecretNumber(start int) int {
	step1 := prune(mix(start, start*64))

	step2 := prune(mix(step1, step1/32))

	return prune(mix(step2, step2*2048))
}

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func calculateSequencesMap(secretNumbers []int) map[string]int {
	sequencesMap := make(map[string]int)

	for i := 1; i < len(secretNumbers)-5; i++ {
		sequence := fmt.Sprintf(
			"%d,%d,%d,%d",
			secretNumbers[i]-secretNumbers[i+1],
			secretNumbers[i+1]-secretNumbers[i+2],
			secretNumbers[i+2]-secretNumbers[i+3],
			secretNumbers[i+3]-secretNumbers[i+4],
		)

		// Only care about first instance of a value
		if _, found := sequencesMap[sequence]; found {
			continue
		}

		sequencesMap[sequence] = secretNumbers[i+4]
	}

	return sequencesMap
}
