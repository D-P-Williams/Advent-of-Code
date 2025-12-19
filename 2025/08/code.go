package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
)

type point struct {
	x, y, z int
	id      int // Line number
}

type pair struct {
	a, b point
	key  string
	dist float64
}

type circuit []pair

var skipList = []pair{}

func main() {
	lines := aoc.ReadLines("./input.txt")
	// lines := aoc.ReadLines("./example.txt")

	countPt1 := 0
	countPt2 := 0

	junctions := linesToJunctions(lines)

	pairs := calculateDistances(junctions)

	circuits := findCircuits(pairs, 1000)

	countPt1 = productLargestCircuits(circuits)

	fmt.Println("part 1", countPt1)

	// Slow, but works.
	finalPair := mergeAllCircuits(pairs, len(junctions))
	fmt.Println(finalPair.a, finalPair.b)
	countPt2 = finalPair.a.x * finalPair.b.x

	fmt.Println("part 2", countPt2)
}

func linesToJunctions(lines []string) []point {
	junctions := make([]point, 0, len(lines))

	for i, line := range lines {
		lineParts := strings.Split(line, ",")

		x, y, z := lineParts[0], lineParts[1], lineParts[2]

		junctions = append(junctions, point{
			x:  aoc.ParseInt(x),
			y:  aoc.ParseInt(y),
			z:  aoc.ParseInt(z),
			id: i,
		})
	}

	return junctions
}

func calculateDistances(points []point) map[string]pair {
	pairs := map[string]pair{}

	for _, source := range points {
		for _, destination := range points {
			// Skip self
			if source == destination {
				continue
			}

			// Skip pairs we've already checked
			keyValues := []int{source.id, destination.id}

			slices.Sort(keyValues)

			key := fmt.Sprintf("%d-%d", keyValues[0], keyValues[1])

			_, found := pairs[key]
			if found {
				continue
			}

			// Otherwise, calc distance and add to map
			distance := euclideanDistance(source, destination)

			pairs[key] = pair{
				a:    source,
				b:    destination,
				key:  key,
				dist: distance,
			}
		}
	}

	return pairs
}

func findClosest(pairs []pair, skip []pair) pair {
	for _, pair := range pairs {
		if slices.Contains(skip, pair) {
			continue
		}

		return pair
	}

	panic("should not happen")
}

// euclidean distance = sqrt(dx^2 + dx^2 + dz^2)
// But an optimisation applied as we don't need the actual distance, but the closest, so skip the sqrt.
func euclideanDistance(a, b point) float64 {
	return math.Pow(float64(a.x)-(float64(b.x)), 2) +
		math.Pow(float64(a.y)-(float64(b.y)), 2) +
		math.Pow(float64(a.z)-(float64(b.z)), 2)
}

func findCircuits(pairs map[string]pair, count int) []circuit {
	circuits := []circuit{}

	// Sort all pairs by distance
	distSortedPairs := []pair{}

	for _, pair := range pairs {
		distSortedPairs = append(distSortedPairs, pair)
	}

	slices.SortFunc(distSortedPairs, func(a, b pair) int {
		return int(a.dist - b.dist)
	})

	for range count {
		closestPair := findClosest(distSortedPairs, skipList)

		// Check if one of the closestPair points already exists in a circuit
		idxA := slices.IndexFunc(circuits, func(c circuit) bool {
			return slices.ContainsFunc(c, func(p pair) bool {
				return closestPair.a == p.a || closestPair.a == p.b
			})
		})

		idxB := slices.IndexFunc(circuits, func(c circuit) bool {
			return slices.ContainsFunc(c, func(p pair) bool {
				return closestPair.b == p.a || closestPair.b == p.b
			})
		})

		// Both found in different circuits, merge them
		if idxA != -1 && idxB != -1 && idxA != idxB {
			// Add B's pairs to A
			circuitB := circuits[idxB]

			circuits[idxA] = append(circuits[idxA], circuitB...)

			// Remove B from circuits list
			circuits = append(circuits[:idxB], circuits[idxB+1:]...)

			// Or, Add to A's circuit
		} else if idxA != -1 {
			circuits[idxA] = append(circuits[idxA], closestPair)

			// Or, Add to B's circuit
		} else if idxB != -1 {
			circuits[idxB] = append(circuits[idxB], closestPair)

			// Or, make a new circuit
		} else {
			circuits = append(circuits, circuit{closestPair})
		}

		// After use, add it to the skip list to stop re-using it
		skipList = append(skipList, closestPair)
	}

	return circuits
}

func mergeAllCircuits(pairs map[string]pair, pointsCount int) pair {
	circuits := []map[pair]bool{}

	// Sort all pairs by distance
	distSortedPairs := []pair{}

	for _, pair := range pairs {
		distSortedPairs = append(distSortedPairs, pair)
	}

	slices.SortFunc(distSortedPairs, func(a, b pair) int {
		return int(a.dist - b.dist)
	})

	for {
		closestPair := findClosest(distSortedPairs, skipList)

		// Check if one of the closestPair points already exists in a circuit

		idxA := slices.IndexFunc(circuits, func(c map[pair]bool) bool {
			for pair := range c {
				if pair.a.id == closestPair.a.id || pair.b.id == closestPair.a.id {
					return true
				}
			}

			return false
		})

		idxB := slices.IndexFunc(circuits, func(c map[pair]bool) bool {
			for pair := range c {
				if pair.a.id == closestPair.b.id || pair.b.id == closestPair.b.id {
					return true
				}
			}

			return false
		})

		// Both found in different circuits, merge them
		if idxA != -1 && idxB != -1 && idxA != idxB {
			// Add B's pairs to A
			circuitB := circuits[idxB]

			for pair := range circuitB {
				circuits[idxA][pair] = true
			}

			// Insert current circuit
			circuits[idxA][closestPair] = true

			// Remove B from circuits list
			circuits = append(circuits[:idxB], circuits[idxB+1:]...)

			// Or, Add to A's circuit
		} else if idxA != -1 {
			circuits[idxA][closestPair] = true

			// Or, Add to B's circuit
		} else if idxB != -1 {
			circuits[idxB][closestPair] = true

			// Or, make a new circuit
		} else {
			circuits = append(circuits, map[pair]bool{closestPair: true})
		}

		// After use, add it to the skip list to stop re-using it
		skipList = append(skipList, closestPair)

		// fmt.Println(longestCircuitLength(circuits), pointsCount, longestCircuitLength(circuits) == pointsCount, closestPair.a, closestPair.b)

		if longestCircuitLength(circuits) == pointsCount {
			return closestPair
		}
	}
}

func productLargestCircuits(circuits []circuit) int {
	slices.SortFunc(circuits, func(a, b circuit) int {
		return len(b) - len(a)
	})

	product := 1

	// +1 as the final node in the chain isn't accounted for using pair chains.
	product *= countPoints(circuits[0])
	product *= countPoints(circuits[1])
	product *= countPoints(circuits[2])

	return product
}

func countPoints(circuit circuit) int {
	count := 0
	points := []point{}

	for _, pair := range circuit {
		if !slices.Contains(points, pair.a) {
			count++
			points = append(points, pair.a)
		}

		if !slices.Contains(points, pair.b) {
			count++
			points = append(points, pair.b)
		}
	}

	return count
}

func longestCircuitLength(circuits []map[pair]bool) int {
	maxLength := 0

	for _, circuit := range circuits {
		points := map[point]bool{}

		for pair := range circuit {
			points[pair.a] = true
			points[pair.b] = true
		}

		circuitLength := len(points)

		if circuitLength > maxLength {
			maxLength = circuitLength
		}
	}

	return maxLength
}
