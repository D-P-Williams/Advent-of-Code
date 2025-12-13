package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	computer := initialiseComputer(lines)

	computer.runProgram()

	fmt.Println("part 1", computer.readOutput())

	fmt.Println("part 2", reverseEngineer(computer.Program))
}

func initialiseComputer(lines []string) *Computer {
	registerALine := lines[0]
	registerBLine := lines[1]
	registerCLine := lines[2]
	programLine := lines[4]

	registerAString := strings.Split(registerALine, ": ")[1]
	registerBString := strings.Split(registerBLine, ": ")[1]
	registerCString := strings.Split(registerCLine, ": ")[1]

	// Split just the number section, then split on commas
	programString := strings.Split(strings.Split(programLine, ": ")[1], ",")

	return &Computer{
		A:                  aoc.ParseInt(registerAString),
		B:                  aoc.ParseInt(registerBString),
		C:                  aoc.ParseInt(registerCString),
		Program:            programString,
		isHalted:           false,
		instructionPointer: 0,
		output:             []int{},
	}
}

// Reverse the program list, starting at the end check the 8 possibilities (3-bit = 8 values, 0-7)
// for the value which returns the input, shift by 3 bits and repeat.
//
// Gradually building up the single large int which processes down to the program list.
//
// Possible as each program loop ends by shifting A right by 3 (divide 8 = shift 3 for 3 bit),
// Which effectively "moves" all but the final 3 bits out of the area of interest for that loop iteration.
func reverseEngineer(programLine []string) int {
	reversedProgram := programLine
	slices.Reverse(reversedProgram)

	validValuesForA := []int{0}

	for _, requiredInstruction := range reversedProgram {
		nextValuesForA := []int{}

		goodOut := []int{}

		for _, A := range validValuesForA {
			AShifted := A << 3

			for candidateA := AShifted; candidateA < AShifted+8; candidateA += 1 {
				out := runProgram(candidateA)

				if len(out) > 0 && out[0] == aoc.ParseInt(requiredInstruction) {
					goodOut = out
					nextValuesForA = append(nextValuesForA, candidateA)
				}
			}
		}

		fmt.Println("out:", goodOut)
		validValuesForA = nextValuesForA
	}

	return slices.Min(validValuesForA)
}

func runProgram(A int) []int {
	B := 0
	C := 0

	out := []int{}

	for A != 0 { // 3,0
		B = A % 8                          // 2,4
		B = B ^ 6                          // 1,6
		C = A / int(math.Exp2(float64(B))) // 7,5
		B = B ^ C                          // 4,6
		B = B ^ 4                          // 1,4
		out = append(out, B%8)             // 5,5
		A = A / 8                          // 0,3
	}

	return out
}
