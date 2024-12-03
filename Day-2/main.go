package main

import (
	"fmt"
	"math"
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

	safeCountPt1 := 0
	safeCountPt2 := 0

	for _, line := range lines {
		result := checkLine(line)

		if result {
			safeCountPt1 += 1
		}

		readings := strings.Split(line, " ")
		for i := 0; i < len(readings); i += 1 {
			readings := strings.Split(line, " ")

			result := checkLine(strings.Join(append(readings[:i], readings[i+1:]...), " "))

			if result {
				safeCountPt2 += 1
				break
			}
		}
	}

	fmt.Println("part 1", safeCountPt1)

	fmt.Println("part 2", safeCountPt2)
}

func checkLine(line string) bool {
	readings := strings.Split(line, " ")

	var slopeIncreasing bool

	for i := 0; i < len(readings)-1; i += 1 {
		current := readings[i]
		next := readings[i+1]

		currInt, err := strconv.ParseInt(current, 10, 0)
		errorCheck(err)

		nextInt, err := strconv.ParseInt(next, 10, 0)
		errorCheck(err)

		if i == 0 {
			slopeIncreasing = currInt > nextInt
		}

		if currInt > nextInt != slopeIncreasing {
			return false
		}

		diff := int(math.Abs(float64(nextInt) - float64(currInt)))

		if diff < 1 {
			return false
		}

		if diff > 3 {
			return false
		}
	}

	return true
}
