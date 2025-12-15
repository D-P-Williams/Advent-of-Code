package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

func main() {
	lines := aoc.ReadLines("./input.txt")
	// lines := aoc.ReadLines("./example.txt")

	countPt1 := 0
	countPt2 := 0

	countPt1, countPt2 = checkStock(lines)

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func checkStock(lines []string) (int, int) {
	breakIdx := slices.Index(lines, "")

	idRangeStrings := lines[:breakIdx]

	idRanges := make([][]int, 0, len(idRangeStrings))

	for _, idRange := range idRangeStrings {
		idRanges = append(idRanges, parseRange(idRange))
	}

	items := lines[breakIdx+1:]

	fresh := 0

	for _, item := range items {
		if checkFreshness(idRanges, item) {
			fresh++
		}
	}

	changed := true
	mergedRanges := idRanges

	for changed {
		mergedRanges, changed = mergeRanges(mergedRanges)
		for _, merged := range mergedRanges {
			fmt.Println(merged)
		}
	}

	freshIds := countFreshIds(mergedRanges)

	return fresh, freshIds
}

func checkFreshness(idRanges [][]int, line string) bool {
	id := aoc.ParseInt(line)

	for _, idRange := range idRanges {
		if idRange[0] <= id && id <= idRange[1] {
			return true
		}
	}

	return false
}

func countFreshIds(idRanges [][]int) int {
	freshIds := 0

	for _, idRange := range idRanges {
		freshIds += idRange[1] - idRange[0] + 1
	}

	return freshIds
}

func parseRange(line string) []int {
	lineParts := strings.Split(line, "-")

	return []int{aoc.ParseInt(lineParts[0]), aoc.ParseInt(lineParts[1])}
}

// merged range slice [][]int

// if current range fits wholly within any existing range, skip it

// if current range overlaps with any existing range, merge them

// else add current range as new range

func mergeRanges(ranges [][]int) ([][]int, bool) {
	merged := make([][]int, 0, len(ranges))
	changed := false

	for _, current := range ranges {
		overlapped := false

		for i, existing := range merged {
			// current wholly within existing
			if current[0] >= existing[0] && current[1] <= existing[1] {
				overlapped = true
				break
			}

			if current[1] >= existing[0] && current[0] <= existing[1] {
				newStart := math.Min(float64(current[0]), float64(existing[0]))
				newEnd := math.Max(float64(current[1]), float64(existing[1]))
				merged[i] = []int{int(newStart), int(newEnd)}
				overlapped = true
				changed = true
				break
			}
		}

		if !overlapped {
			merged = append(merged, current)
		}
	}

	return merged, changed
}
