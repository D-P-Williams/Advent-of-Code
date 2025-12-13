package main

import (
	"fmt"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type Lock struct {
	pin1 int
	pin2 int
	pin3 int
	pin4 int
	pin5 int
}

func (l *Lock) fits(key Key) bool {
	if l.pin1+key.tumbler1 > 7 {
		return false
	}

	if l.pin2+key.tumbler2 > 7 {
		return false
	}

	if l.pin3+key.tumbler3 > 7 {
		return false
	}

	if l.pin4+key.tumbler4 > 7 {
		return false
	}

	if l.pin5+key.tumbler5 > 7 {
		return false
	}

	return true
}

type Key struct {
	tumbler1 int
	tumbler2 int
	tumbler3 int
	tumbler4 int
	tumbler5 int
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	countPt1 := 0
	countPt2 := 0

	locks, keys := parseInput(lines)

	for _, lock := range locks {
		for _, key := range keys {

			if lock.fits(key) {
				countPt1 += 1
			}
		}
	}

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

func parseInput(lines []string) ([]Lock, []Key) {
	locks := []Lock{}
	keys := []Key{}

	sections := [][]string{}
	section := []string{}

	for _, line := range lines {

		if line == "" {
			sections = append(sections, section)
			section = []string{}
			continue
		}

		section = append(section, line)
	}

	sections = append(sections, section)

	for _, section := range sections {
		if string(section[0][0]) == "#" {
			locks = append(locks, Lock{
				pin1: countCol(section, 0),
				pin2: countCol(section, 1),
				pin3: countCol(section, 2),
				pin4: countCol(section, 3),
				pin5: countCol(section, 4),
			})

			continue
		}

		keys = append(keys, Key{
			tumbler1: countCol(section, 0),
			tumbler2: countCol(section, 1),
			tumbler3: countCol(section, 2),
			tumbler4: countCol(section, 3),
			tumbler5: countCol(section, 4),
		})
	}

	return locks, keys
}

func countCol(section []string, index int) int {
	N := 0

	for _, row := range section {
		if string(row[index]) == "#" {
			N += 1
		}
	}

	return N
}
