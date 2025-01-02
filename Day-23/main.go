package main

import (
	"bytes"
	"fmt"
	"maps"
	"slices"
	"strings"

	aoc "github.com/D-P-Williams/Advent-of-Code-24"
)

func main() {
	lines := aoc.ReadLines("./input.txt")

	graph := collateMaps(lines)

	loops := checkMatches(graph)

	graphKeys := []string{}

	for key := range maps.Keys(graph) {
		graphKeys = append(graphKeys, key)
	}

	loopsPt2 := bronKerbosch(graph, []string{}, graphKeys, []string{}, [][]string{})

	password := generatePassword(loopsPt2)

	fmt.Println("part 1", len(loops))

	fmt.Println("part 2", password)
}

func collateMaps(lines []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range lines {
		lineParts := strings.Split(line, "-")

		graph[lineParts[0]] = append(graph[lineParts[0]], lineParts[1])
		graph[lineParts[1]] = append(graph[lineParts[1]], lineParts[0])
	}

	return graph
}

func checkMatches(graph map[string][]string) [][]string {
	loops := make([][]string, 0)

	for key := range maps.Keys(graph) {
		for _, loopA := range graph[key] {
			for _, loopB := range graph[loopA] {
				if slices.Contains(graph[key], loopB) {
					loop := []string{key, loopA, loopB}

					slices.Sort(loop)

					if !strings.HasPrefix(key, "t") && !strings.HasPrefix(loopA, "t") && !strings.HasPrefix(loopB, "t") {
						continue
					}

					if !slices.ContainsFunc(loops, func(existing []string) bool {
						return slices.Equal(existing, loop)
					}) {
						loops = append(loops, loop)
					}

				}
			}
		}
	}

	return loops
}

func bronKerbosch(graph map[string][]string, selected []string, candidates []string, excluded []string, cliques [][]string) [][]string {
	if len(candidates) == 0 && len(excluded) == 0 {
		cliques = append(cliques, append([]string{}, selected...))
		return cliques
	}

	for index := 0; index < len(candidates); {
		node := candidates[index]

		cliques = bronKerbosch(graph, append(selected, node), intersection(candidates, graph[node]), intersection(excluded, graph[node]), cliques)

		candidates = append(candidates[:index], candidates[index+1:]...)
		excluded = append(excluded, node)
	}

	return cliques
}

func intersection(a, b []string) []string {
	intersection := []string{}

	for _, aItem := range a {
		if slices.Contains(b, aItem) {
			intersection = append(intersection, aItem)
		}
	}

	return intersection
}

func generatePassword(loops [][]string) string {
	longestLoop := []string{}

	for _, loop := range loops {
		if len(loop) > len(longestLoop) {
			longestLoop = loop
		}
	}

	slices.Sort(longestLoop)

	var password bytes.Buffer

	for _, element := range longestLoop {
		password.WriteString(element + ",")
	}

	return password.String()[0 : len(password.String())-1]
}
