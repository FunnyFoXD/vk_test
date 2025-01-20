package models

type Point struct {
	X, Y int
}

// IsValid checks if the point is valid in the given labyrinth.
// It returns true if the point is valid and false otherwise.
// The point is valid if it is within the bounds of the labyrinth and
// the value of the cell at the point is not zero (i.e. it is not a wall).
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