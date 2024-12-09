package main

import (
	"fmt"
	"slices"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

type cell struct {
	// -1 represents empty cell
	value int
}

func main() {
	lines := aoc.ReadLines("./input.txt")

	startingData := lines[0]

	expandedData := expandData(startingData)

	compactedData := compactData(expandedData)

	checksumPt1 := calculateChecksum(compactedData)

	fmt.Println("part 1", checksumPt1)

	startingData = lines[0]

	expandedData = expandData(startingData)

	// defraggedData := defragData(expandedData)
	defraggedData := defragData(expandedData)

	// fmt.Println(defraggedData)

	checksumPt2 := calculateChecksum(defraggedData)

	fmt.Println("part 2", checksumPt2)
}

func expandData(compactedData string) []cell {
	output := []cell{}

	isFile := true
	fileId := 0

	for _, char := range compactedData {
		value := aoc.ParseInt(string(char))

		cellValue := -1

		if isFile {
			cellValue = fileId
			fileId += 1
		}

		for range value {
			output = append(output, cell{
				value: cellValue,
			})
		}

		isFile = !isFile
	}

	return output
}

func compactData(data []cell) []cell {
	for i := len(data) - 1; i > 0; i -= 1 {
		if data[i].value == -1 {
			continue
		}

		for j := 0; j < len(data)-1; j += 1 {
			if j > i {
				return data
			}

			if data[j].value == -1 {
				data[i], data[j] = data[j], data[i]
				break
			}
		}
	}

	return data
}

func calculateChecksum(data []cell) int {
	checksum := 0

	for cellIdx, cell := range data {
		if cell.value != -1 {
			checksum += cellIdx * cell.value
		}
	}

	return checksum
}

func defragData(data []cell) []cell {
	fileId := 0

	for i := len(data) - 1; i > 0; i -= 1 {
		// Early exit if empty
		if data[i].value == -1 {
			continue
		}

		fileId = data[i].value

		// Find first index, looping from right to left so this gives us the range
		fileEndIdx := findFirstIndexOf(data, fileId)

		fileLength := i - fileEndIdx + 1

		// Find continuous space of length fileLength, or -1
		spaceStart, spaceLength := findSpaceIndex(data, i, fileLength)

		// If a sufficient space exists, swap the data
		if spaceStart != -1 && spaceLength >= fileLength {
			data = swapRanges(data, fileEndIdx, spaceStart, fileLength, fileId)
		}

		// Decrement loop variable by file length to not double shift
		for range fileLength - 1 {
			i -= 1
		}
	}

	return data
}

func findFirstIndexOf(slice []cell, target int) int {
	return slices.IndexFunc(slice, func(a cell) bool {
		return a.value == target
	})
}

func findSpaceIndex(slice []cell, fileStart, fileLength int) (int, int) {
	spaceStart, spaceLength := -1, 0
	for i := 0; i < fileStart; i++ {
		if slice[i].value == -1 {
			if spaceStart == -1 {
				spaceStart = i
			}
			spaceLength++

			if spaceLength >= fileLength {
				break
			}
		} else {
			spaceStart, spaceLength = -1, 0
		}
	}

	return spaceStart, spaceLength
}

func swapRanges(data []cell, fileStart, spaceStart, fileLength, target int) []cell {
	for i := 0; i < fileLength; i++ {
		data[spaceStart+i].value = target
	}

	for i := fileStart; i <= fileStart+fileLength-1; i++ {
		data[i].value = -1
	}

	return data
}
