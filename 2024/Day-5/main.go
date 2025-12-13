package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines := aoc.ReadLines("./input.txt")
	dividerIdx := slices.Index(lines, "")

	rules := lines[:dividerIdx]
	updates := lines[dividerIdx+1:]

	invalidUpdates := []string{}

	countPt1 := 0
	countPt2 := 0

	for _, update := range updates {
		updateRow := strings.Split(update, ",")

		valid := checkUpdateValidity(rules, updateRow)

		if valid {
			middleValue, err := strconv.ParseInt(updateRow[len(updateRow)/2], 10, 0)
			errorCheck(err)

			countPt1 += int(middleValue)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fmt.Println("part 1", countPt1)

	countPt2 = correctUpdate(rules, invalidUpdates)

	fmt.Println("part 2", countPt2)
}

func checkUpdateValidity(rules []string, update []string) bool {
	for _, rule := range rules {
		ruleValues := strings.Split(rule, "|")

		idxA := slices.Index(update, ruleValues[0])
		idxB := slices.Index(update, ruleValues[1])

		if idxA == -1 || idxB == -1 {
			continue
		}

		if idxA > idxB {
			return false
		}
	}

	return true
}

func correctUpdate(rules []string, updates []string) (r int) {
	comparison := func(a, b string) int {
		for _, rule := range rules {
			if rule := strings.Split(rule, "|"); rule[0] == a && rule[1] == b {
				return -1
			}
		}

		return 0
	}

	for _, update := range updates {
		if update := strings.Split(update, ","); !slices.IsSortedFunc(update, comparison) {
			slices.SortFunc(update, comparison)
			n, _ := strconv.Atoi(update[len(update)/2])
			r += n
		}

	}

	return r
}
