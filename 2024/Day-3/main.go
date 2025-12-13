package main

import (
	"fmt"
	"regexp"
	"strconv"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	regexPt1 := regexp.MustCompile(`mul\((\d*?),(\d*?)\)`)

	countPt1 := 0

	for _, line := range lines {
		matches := regexPt1.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, err := strconv.ParseInt(match[1], 10, 0)
			errorCheck(err)

			num2, err := strconv.ParseInt(match[2], 10, 0)
			errorCheck(err)

			countPt1 += int(num1 * num2)
		}
	}

	regexPt2 := regexp.MustCompile(`mul\((\d*?),(\d*?)\)|do\(\)|don't\(\)`)

	countPt2 := 0

	enabled := true

	for _, line := range lines {
		localLine := line

		matchLocs := regexPt2.FindStringIndex(localLine)

		for matchLocs != nil {
			fmt.Println(matchLocs)
			localMatch := regexPt2.FindStringSubmatch(localLine[matchLocs[0]:matchLocs[1]])
			fmt.Println(localMatch)

			localLine = localLine[matchLocs[1]:]

			if localMatch[0] == "do()" {
				enabled = true

				matchLocs = regexPt2.FindStringIndex(localLine)
				continue
			} else if localMatch[0] == "don't()" {
				enabled = false

				matchLocs = regexPt2.FindStringIndex(localLine)
				continue
			}

			num1, err := strconv.ParseInt(localMatch[1], 10, 0)
			errorCheck(err)

			num2, err := strconv.ParseInt(localMatch[2], 10, 0)
			errorCheck(err)

			if enabled {
				countPt2 += int(num1 * num2)
			}

			matchLocs = regexPt2.FindStringIndex(localLine)
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}
