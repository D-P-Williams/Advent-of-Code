package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

var seenStones = map[string][]string{}

func main() {
	grid := aoc.ReadLines("./input.txt")

	stonesInitial := strings.Split(grid[0], " ")

	stones := map[string]int{}

	for _, stone := range stonesInitial {
		stones[stone] += 1
	}

	for range 25 {
		stones = processBlink(stones)
	}

	countPt1 := 0

	for _, count := range stones {
		countPt1 += count
	}

	fmt.Println("part 1", countPt1)

	// 75 blinks total, process another 50
	for range 50 {
		stones = processBlink(stones)
	}

	countPt2 := 0
	for _, count := range stones {
		countPt2 += count
	}

	fmt.Println("part 2", countPt2)
}

func processBlink(stones map[string]int) map[string]int {
	output := map[string]int{}

	for stone, count := range stones {
		// If already processed value, re-use it
		if result, found := seenStones[stone]; found {
			for _, val := range result {
				output[val] += count
			}

			continue
		}

		// Otherwise, calculate value and add to the map

		if stone == "0" {
			output["1"] += count
			seenStones[stone] = []string{"1"}

			continue
		}

		if len(stone)%2 == 0 {
			leftHalf, rightHalf := stone[:len(stone)/2], stone[len(stone)/2:]

			// Trim any leading zeros
			rightHalf = strings.TrimLeft(rightHalf, "0")

			// If rightHalf was only 0's, all characters are trimmed out
			if rightHalf == "" {
				rightHalf = "0"
			}

			output[leftHalf] += count
			output[rightHalf] += count
			seenStones[stone] = []string{leftHalf, rightHalf}

			continue
		}

		value, err := strconv.ParseInt(stone, 10, 64)
		if err != nil {
			panic(err)
		}

		output[strconv.FormatInt(value*2024, 10)] += count
		seenStones[stone] = []string{strconv.FormatInt(value*2024, 10)}
	}

	return output
}
