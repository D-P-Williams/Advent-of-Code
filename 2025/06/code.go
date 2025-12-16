package main

import (
	"fmt"
	"math"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

func main() {
	lines := aoc.ReadLines("./input.txt")
	// lines := aoc.ReadLines("./example.txt")

	countPt1 := 0
	countPt2 := 0

	rowReaders := []*strings.Reader{}

	for _, line := range lines {
		reader := strings.NewReader(line)

		rowReaders = append(rowReaders, reader)
	}

	for {
		values, valid := getNextEquation(rowReaders)
		if !valid {
			break
		}

		result := solveEquation(values)
		// fmt.Println(result)
		countPt1 += result

		result = solveEquationCephalopod(values)
		// fmt.Println(result)
		countPt2 += result

		normaliseReaders(rowReaders)
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func getNextEquation(rows []*strings.Reader) (values []string, valid bool) {
	for _, row := range rows {
		value, valid := getNextEntry(row)
		if !valid {
			return nil, false
		}

		values = append(values, value)
	}

	return values, true
}

func getNextEntry(reader *strings.Reader) (string, bool) {
	entry := ""
	foundStart := false

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			// On an EOF error, if we've started a number, we've just reached the end of it
			// otherwise return invalid
			if err.Error() == "EOF" {
				if !foundStart {
					return "", false
				}

				return entry, true
			}
			panic("string reading error")
		}

		// 32 = space, so skip over. Unless found the start, in which case we are at the end and return
		if char == 32 && foundStart {
			break
		}

		if char == 32 && len(entry) > 0 {

		}

		entry += string(char)

		if char != 32 {
			foundStart = true
		}
	}

	return entry, true
}

func normaliseReaders(readers []*strings.Reader) {
	minLength := math.MaxInt

	// Find the longest number
	for _, reader := range readers {
		minLength = int(math.Min(float64(reader.Len()), float64(minLength)))
	}

	for _, reader := range readers {
		if reader.Len() > minLength {
			for range reader.Len() - minLength {
				reader.ReadByte()
			}
		}
	}
}

func solveEquation(values []string) int {
	total := 0

	operation := strings.ReplaceAll(values[len(values)-1], " ", "")

	switch operation {
	case "+":
		// Loop, skipping the final element
		for i := 0; i <= len(values)-2; i++ {
			total += aoc.ParseInt(strings.ReplaceAll(values[i], " ", ""))
		}

	case "*":
		// Set total to 1 for first multiplication
		total = 1

		for i := 0; i <= len(values)-2; i++ {
			total *= aoc.ParseInt(strings.ReplaceAll(values[i], " ", ""))
		}
	default:
		panic("unexpected operation found: " + operation)
	}

	return total
}

func solveEquationCephalopod(values []string) int {
	total := 0

	operation := strings.ReplaceAll(values[len(values)-1], " ", "")

	transposedValues := convertToCephalopod(values)

	switch operation {
	case "+":
		// Loop, skipping the final element
		for i := 0; i <= len(transposedValues)-1; i++ {
			if values[i] == "" {
				continue
			}

			val := strings.ReplaceAll(transposedValues[i], " ", "")

			if val != "" {
				total += aoc.ParseInt(strings.ReplaceAll(transposedValues[i], " ", ""))
			}
		}

	case "*":
		// Set total to 1 for first multiplication
		total = 1

		for i := 0; i <= len(transposedValues)-1; i++ {
			val := strings.ReplaceAll(transposedValues[i], " ", "")

			if val != "" {
				total *= aoc.ParseInt(strings.ReplaceAll(transposedValues[i], " ", ""))
			}
		}
	default:
		panic("unexpected operation found: " + operation)
	}

	return total
}

func convertToCephalopod(values []string) []string {
	maxLength := 0

	// Find the longest number
	for _, value := range values {
		maxLength = int(math.Max(float64(len(value)), float64(maxLength)))
	}

	transposed := make([]string, len(values)-1)

	for i := range maxLength {
		val := ""

		for j := 0; j < len(values)-1; j++ {
			idxValue := ""

			if len(values[j]) > i {
				idxValue = string(values[j][i])
			}

			val += idxValue
		}

		// Edge case where double space before all values
		if i == maxLength-1 && strings.ReplaceAll(transposed[0], " ", "") == "" {
			transposed[0] = val
			continue
		}

		transposed[i] = val
	}

	return transposed
}
