package main

import (
	"fmt"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type Wire struct {
	input1    string
	input2    string
	operator  string
	hasResult bool // Handle not yet evaluated set case
	result    bool // Outputs are binary
}

func (w *Wire) canCalculate(wires map[string]*Wire) bool {
	return wires[w.input1].hasResult && wires[w.input2].hasResult
}

func (w *Wire) calculate(wires map[string]*Wire) {
	if !w.canCalculate(wires) {
		panic("Calculate called on wire that can't be calculated")
	}

	switch w.operator {
	case "AND":
		w.result = wires[w.input1].result && wires[w.input2].result
	case "OR":
		w.result = wires[w.input1].result || wires[w.input2].result
	case "XOR":
		w.result = (wires[w.input1].result && !wires[w.input2].result) ||
			(!wires[w.input1].result && wires[w.input2].result)
	default:
		panic("unexpected operator: " + w.operator)
	}

	w.hasResult = true
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	wires := parseInput(lines)

	wires["z10"], wires["z16"] = wires["z16"], wires["z10"]

	wires = evaluate(wires)

	// for key, wire := range wires {
	// 	fmt.Print(key, wire)
	// }
	// fmt.Println()

	countPt1 = generateOutput(wires, "z")

	fmt.Println("part 1", countPt1)

	x := generateOutput(wires, "x")
	y := generateOutput(wires, "y")

	fmt.Println("x", x, "y", y, "z", countPt1, "x + y", x+y, x+y == countPt1)
	fmt.Printf("%046b\n%046b\n%046b\n", countPt1, x+y, x+y^countPt1)

	invalidGates := searchForInvalidGates(wires)

	fmt.Println("Invalid gates", invalidGates)

	fmt.Println("part 2", countPt2)
}

func parseInput(lines []string) map[string]*Wire {
	wires := make(map[string]*Wire)

	inputParts := slices.Index(lines, "")

	// Set up logic gates
	for _, line := range lines[inputParts+1:] {
		lineParts := strings.Split(line, " ")

		// <wire 1> <operator> <wire 2> -> <self>
		wires[lineParts[4]] = &Wire{
			input1:    lineParts[0],
			input2:    lineParts[2],
			operator:  lineParts[1],
			hasResult: false,
			result:    false,
		}
	}

	// Add initial conditions
	for _, line := range lines[:inputParts] {
		lineParts := strings.Split(line, ": ")

		// <wire>: <initial condition>
		wires[lineParts[0]] = &Wire{
			input1:    "",
			input2:    "",
			operator:  "",
			hasResult: true,
			result:    lineParts[1] == "1",
		}
	}

	return wires
}

func evaluate(wires map[string]*Wire) map[string]*Wire {
	keepEvaluating := true
	allEvaluated := true

	for keepEvaluating {
		allEvaluated = true

		for _, wire := range wires {
			if !wire.hasResult {
				allEvaluated = false

				if wire.canCalculate(wires) {
					wire.calculate(wires)
				}
			}
		}

		if allEvaluated {
			break
		}
	}

	return wires
}

func generateOutput(wires map[string]*Wire, prefix string) int {
	var output int

	zWires := make(map[string]*Wire)

	// Filter wires to only outputs
	for key, wire := range wires {
		if strings.HasPrefix(key, prefix) {
			zWires[key] = wire
		}
	}

	for i := 0; i < len(zWires); i += 1 {
		wire := zWires[fmt.Sprintf("%s%02d", prefix, i)]

		if !wire.hasResult {
			panic("Wire has not been evaluated" + fmt.Sprintf("z%02d", i))
		}

		if wire.result {
			output += 1 << i
		}
	}

	return output
}

func searchForInvalidGates(wires map[string]*Wire) []string {
	invalidWires := []string{}

	for key, wire := range wires {
		valid := checkWire(wire, wires[key], wires)

		if !valid {
			invalidWires = append(invalidWires, key)
		}
	}

	return invalidWires
}

/*
AND:

	AND gate can only be input to and OR gate

	AND gate cannot take other AND gate as input

XOR:

	XOR gate can only be input to and AND/XOR gate

	XOR gate cannot take AND gate as input

OR:

	OR gate can only be input of AND/XOR gate

	OR gate can only take AND gate as input

(Xn ⊕ Yn) ⊕ (a + b) should always output a Zxx except for the last carry z45

A gate with Zxx as its output cannot directly use Xn or Yn as inputs.
*/
func checkWire(wire, child *Wire, wires map[string]*Wire) bool {
	if wire.operator == "AND" {
		// // AND gate can only be input to an OR gate
		// if child.operator != "OR" {
		// 	return false
		// }

		// AND gate cannot take other AND gate as input
		if wires[wire.input1].operator == "AND" || wires[wire.input2].operator == "AND" {
			return false
		}
	}

	if wire.operator == "XOR" {
		// // XOR gate can only be input to and AND/XOR gate
		// if child.operator == "OR" {
		// 	return false
		// }

		// XOR gate cannot take AND gate as input
		if wires[wire.input1].operator == "AND" || wires[wire.input2].operator == "AND" {
			return false
		}
	}

	if wire.operator == "OR" {
		// // OR gate can only be input of AND/XOR gate
		// if child.operator == "OR" {
		// 	return false
		// }

		// OR gate can only take AND gate as input
		if wires[wire.input1].operator != "AND" || wires[wire.input2].operator != "AND" {
			return false
		}
	}

	return true
}
