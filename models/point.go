package models

type Point struct {
	X, Y int
}

func (p Point) IsValid(labyrinth [][]int) bool {
	return p.X >= 0 && p.X < len(labyrinth) &&
		p.Y >= 0 && p.Y < len(labyrinth[0]) &&
		labyrinth[p.X][p.Y] != 0
}

type Node struct {
	Point Point
	GCost int
	HCost int
	FCost int
	Parent *Node
}