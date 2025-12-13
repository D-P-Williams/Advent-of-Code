package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	for _, line := range lines {
		countPt1 += checkLine(line, false)
		countPt2 += checkLine(line, true)
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func checkLine(line string, includeCombine bool) int {
	stringParts := strings.Split(line, ": ")

	targetString := stringParts[0]
	valuesString := stringParts[1]

	splitValues := strings.Split(valuesString, " ")

	target := aoc.ParseInt(targetString)

	var values = make([]int, 0, len(valuesString))

	for _, valueString := range splitValues {
		values = append(values, aoc.ParseInt(string(valueString)))
	}

	matches := traverse(target, values, includeCombine)

	return matches
}

func traverse(target int, values []int, includeCombine bool) int {
	successful := addTreeLevel(values[0], values[1:], target, includeCombine)

	if len(successful) > 0 {
		return target
	}

	return 0
}

func addTreeLevel(value int, modifiers []int, target int, includeCombine bool) (successful []int) {
	if len(modifiers) == 0 {
		if value == target {
			successful = append(successful, target)
		}

		return successful
	}

	modifier := modifiers[0]

	successful = append(successful, addTreeLevel(value+modifier, modifiers[1:], target, includeCombine)...)
	successful = append(successful, addTreeLevel(value*modifier, modifiers[1:], target, includeCombine)...)

	if includeCombine {
		combinedInt := aoc.ParseInt(strconv.Itoa(value) + strconv.Itoa(modifiers[0]))
		successful = append(successful, addTreeLevel(combinedInt, modifiers[1:], target, includeCombine)...)
	}

	return successful
}
