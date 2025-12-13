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

	countPt1, countPt2 = addUpInvalidIDs(lines[0])

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func addUpInvalidIDs(data string) (int, int) {
	invalidSumPt1 := 0
	invalidSumPt2 := 0

	IDPairs := strings.SplitSeq(data, ",")

	for pair := range IDPairs {
		IDs := strings.Split(pair, "-")

		startId := aoc.ParseInt(IDs[0])
		endId := aoc.ParseInt(IDs[1])

		for id := startId; id <= endId; id++ {
			idStr := fmt.Sprintf("%d", id)

			if !checkValidity(idStr) {
				invalidSumPt1 += id
			}

			if !checkValidityPt2(idStr) {
				invalidSumPt2 += id
			}

			// fmt.Println("ID:", idStr, "valid:", checkValidityPt2(idStr))
		}

	}

	return invalidSumPt1, invalidSumPt2
}

func checkValidity(id string) bool {
	// If odd length, can't be evenly split so can't be invalid
	if len(id)%2 != 0 {
		return true
	}

	midpoint := len(id) / 2
	firstHalf := id[:midpoint]
	secondHalf := id[midpoint:]

	return firstHalf != secondHalf
}

func checkValidityPt2(id string) bool {
	midpoint := len(id) / 2

	for i := 1; i < len(id); i++ {
		// If we get over half way, must be valid ID
		if i > midpoint {
			return true
		}

		potentialPrefix := id[:i]
		remaining := id[i:]

		for len(remaining) > 0 {
			localRemaining, hadPrefix := strings.CutPrefix(remaining, potentialPrefix)

			// fmt.Println("Prefix:", potentialPrefix, "Remaining:", localRemaining, hadPrefix)

			if !hadPrefix {
				break
			}

			remaining = localRemaining

			if remaining == "" {
				return false
			}
		}
	}

	return true
}
