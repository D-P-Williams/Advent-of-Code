package aoc

type Point struct {
	X int
	Y int
}

type Node struct {
	Point
	Value     string
	Neighbors []string
}
