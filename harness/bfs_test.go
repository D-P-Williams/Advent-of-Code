package aoc_test

import (
	"testing"

	aoc "github.com/D-P-Williams/Advent-of-Code/harness"
	"github.com/stretchr/testify/assert"
)

func Test_BFS(t *testing.T) {
	tree := map[string]aoc.Node{
		"John":     {Value: "John", Neighbors: []string{"George", "Sam", "Edward"}},
		"George":   {Value: "George", Neighbors: []string{"Richard"}},
		"Sam":      {Value: "Sam", Neighbors: []string{"Richard", "Briana"}},
		"Edward":   {Value: "Edward", Neighbors: []string{"Anett", "Shaun"}},
		"Richard":  {Value: "Richard", Neighbors: []string{"Franklin"}},
		"Briana":   {Value: "Briana", Neighbors: []string{"Lynsey", "Karen"}},
		"Anett":    {Value: "Anett", Neighbors: []string{"Wilson"}},
		"Shaun":    {Value: "Shaun", Neighbors: []string{}},
		"Franklin": {Value: "Franklin", Neighbors: []string{}},
		"Lynsey":   {Value: "Lynsey", Neighbors: []string{}},
		"Karen":    {Value: "Karen", Neighbors: []string{}},
		"Wilson":   {Value: "Wilson", Neighbors: []string{}},
	}

	tests := []struct {
		name     string
		input    map[string]aoc.Node
		start    string
		end      string
		expected []string
	}{
		{
			name:     "john_anett_example",
			input:    tree,
			start:    "John",
			end:      "Anett",
			expected: []string{"John", "Edward", "Anett"},
		},
		{
			name:     "roman_root_not_found",
			input:    tree,
			start:    "Roman",
			end:      "Anett",
			expected: []string{"not_found"},
		},
		{
			name:     "emili_target_not_found",
			input:    tree,
			start:    "John",
			end:      "Emili",
			expected: []string{"not_found"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := aoc.BFS(test.input, test.start, test.end)

			if !assert.Equal(t, test.expected, result) {
				t.FailNow()
			}
		})
	}
}
