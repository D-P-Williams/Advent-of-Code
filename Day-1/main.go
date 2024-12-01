package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines := readLines("./input.txt")

	listA := make([]int, 0, len(lines))
	listB := make([]int, 0, len(lines))

	for _, line := range lines {
		ints := strings.Split(line, " ")

		intA, err := strconv.ParseInt(ints[0], 10, 0)
		errorCheck(err)

		intB, err := strconv.ParseInt(ints[3], 10, 0)
		errorCheck(err)

		listA = append(listA, int(intA))
		listB = append(listB, int(intB))
	}

	slices.Sort(listA)
	slices.Sort(listB)

	total := 0

	for i := 0; i < len(listA); i++ {
		total += int(math.Abs(float64(listA[i]) - float64(listB[i])))
	}

	fmt.Println("part 1", total)

	similarity := 0

	for _, listAItem := range listA {
		multiplier := 0

		for _, listBItem := range listB {
			if listBItem == listAItem {
				multiplier += 1
			}
		}

		similarity += listAItem * multiplier
	}

	fmt.Println("part 2", similarity)
}
