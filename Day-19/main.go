package main

import (
	"fmt"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	towels, patterns := parseInput(lines)

	slices.SortFunc(towels, func(a, b string) int { return len(a) - len(b) })

	longestTowel := len(towels[len(towels)-1])

	possibilities := map[string]bool{"": true}

	for _, pattern := range patterns {
		for {
			if len(possibilities) == 0 {
				break
			}

			possibilities = tryPossibilities(pattern, towels, longestTowel, possibilities)

			if _, found := possibilities[pattern]; found {
				countPt1 += 1
				break
			}
		}

		possibilities = map[string]bool{"": true}
	}

	fmt.Println("part 1", countPt1)

	// This solution is simpler than the one above, but the above is a non-recursion solution
	for _, pattern := range patterns {
		countPt2 += getCombinations(pattern, towels)
	}

	fmt.Println("part 2", countPt2)
}

func parseInput(lines []string) (towels []string, patterns []string) {
	lineBreak := slices.Index(lines, "")

	towelLines := lines[:lineBreak]
	patterns = lines[lineBreak+1:]

	for _, line := range towelLines {
		towelParts := strings.Split(line, ", ")

		towels = append(towels, towelParts...)
	}

	return towels, patterns
}

func tryPossibilities(pattern string, towels []string, maxTowelLength int, possibilities map[string]bool) map[string]bool {
	localPossibilities := map[string]bool{}

	for possibility := range possibilities {
		newPossibilities := findMatchingTowels(pattern, towels, maxTowelLength, len(possibility))

		for newPossibility := range newPossibilities {
			localPossibilities[possibility+newPossibility] = true
		}
	}

	return localPossibilities
}

func findMatchingTowels(pattern string, towels []string, maxTowelLength int, partialIndex int) map[string]bool {
	offset := maxTowelLength

	if partialIndex+offset > len(pattern) {
		offset = len(pattern) - partialIndex
	}

	partialPattern := pattern[partialIndex : partialIndex+offset]

	possibilities := map[string]bool{}

	for _, towel := range towels {
		if strings.HasPrefix(partialPattern, towel) {
			possibilities[towel] = true
		}
	}

	return possibilities
}

var cache = map[string]int{}

func getCombinations(pattern string, towels []string) (combinations int) {
	if n, ok := cache[pattern]; ok {
		return n
	}
	defer func() { cache[pattern] = combinations }()

	// If pattern is empty, i.e. complete match, return a successful combination
	if pattern == "" {
		return 1
	}

	// If pattern has a matching towel prefix, recurse with the remaining pattern
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			combinations += getCombinations(pattern[len(towel):], towels)
		}
	}

	return combinations
}
