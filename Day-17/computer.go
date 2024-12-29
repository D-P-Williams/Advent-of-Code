package main

import (
	"fmt"
	"math"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type Computer struct {
	A                  int
	B                  int
	C                  int
	Program            []string
	isHalted           bool
	instructionPointer int
	output             []int
}

func (c *Computer) getComboOperand() int {
	operand := c.Program[c.instructionPointer+1]

	switch operand {
	case "0", "1", "2", "3":
		return aoc.ParseInt(operand)
	case "4":
		return c.A
	case "5":
		return c.B
	case "6":
		return c.C
	default:
		panic("invalid combo operand, contains reserved operand '7'")

	}
}

func (c *Computer) readOutput() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.output)), ","), "[]")
}

func (c *Computer) runProgram() {
	for !c.isHalted {
		if c.instructionPointer >= len(c.Program) {
			c.isHalted = true
			continue
		}

		opcode := aoc.ParseInt(c.Program[c.instructionPointer])
		operand := aoc.ParseInt(c.Program[c.instructionPointer+1])

		switch opcode {
		case 0: // adv
			c.adv()

		case 1: // bxl
			c.bxl(operand)

		case 2: //bst
			c.bst()

		case 3: // jzn
			if c.jzn(operand) {
				// After a jump, the counter should not be incremented
				continue
			}

		case 4: // bxc
			c.bxc()

		case 5:
			c.out()

		case 6:
			c.bdv()

		case 7:
			c.cdv()

		default:
			panic(fmt.Sprintf("unrecognised opcode: '%d'", opcode))
		}

		c.instructionPointer += 2
	}
}

// A / (2 ** combo operand) and store in A
func (c *Computer) adv() {
	numerator := float64(c.A)
	denominator := math.Exp2(float64(c.getComboOperand())) // 2 ^ comboOperand

	result := numerator / denominator

	c.A = int(result)
}

// Bitwise XOR B with literal operand and store in B
func (c *Computer) bxl(operand int) {
	c.B = c.B ^ operand
}

// Combo operand % 8 and store in B
func (c *Computer) bst() {
	combo := c.getComboOperand()

	c.B = combo % 8
}

// If A non zero, jump instruction pointer to literal operand. Returns true to signify a jump
func (c *Computer) jzn(operand int) bool {
	if c.A == 0 {
		return false
	}

	c.instructionPointer = operand

	return true
}

// Bitwise XOR B with C and store in B
func (c *Computer) bxc() {
	c.B = c.B ^ c.C
}

// Output combo operand % 8
func (c *Computer) out() {
	combo := c.getComboOperand()

	c.output = append(c.output, combo%8)
}

// Same as adv but store in B
func (c *Computer) bdv() {
	numerator := float64(c.A)
	denominator := math.Exp2(float64(c.getComboOperand())) // 2 ^ comboOperand

	result := numerator / denominator

	c.B = int(result)
}

// Same as adv but store in C
func (c *Computer) cdv() {
	numerator := float64(c.A)
	denominator := math.Exp2(float64(c.getComboOperand())) // 2 ^ comboOperand

	result := numerator / denominator

	c.C = int(result)
}
