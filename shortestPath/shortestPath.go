package shortestpath

import (
	"container/heap"
	"fmt"
	"math"

	"githab.com/funnyfoxd/vk_test/models"
	"githab.com/funnyfoxd/vk_test/prq"
)

var directions = []models.Point{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

// heuristics calculates the heuristic cost between two points using the Manhattan
// distance.
//
// It is used in the A* algorithm to guide the search towards the finish point.
//
// It returns the heuristic cost as an integer.
func heuristics(a, b models.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

// AStar implements the A* algorithm to find the shortest path from start to finish
// points in the given labyrinth.
//
// It returns the shortest path as a slice of models.Point and an error if there is
// no path.
//
// The input labyrinth is represented as a 2D slice of integers, where 0 represents
// an empty cell and 1 represents a wall.
//
// The start and finish points are represented as two models.Point.
//
// The A* algorithm is guided by the Manhattan distance heuristic.
//
// If there is an error while finding the path, it will be returned as error.
func AStar(labyrinth [][]int, start, finish models.Point) ([]models.Point, error) {
	closedSet := make(map[models.Point]bool)
	openSet := &prq.PriorityQueue{}
	heap.Init(openSet)

	startNode := &models.Node{
		Point: start,
		GCost: 0,
		HCost: heuristics(start, finish),
		FCost: heuristics(start, finish),
		Parent: nil,
	}
	
	heap.Push(openSet, startNode)

	for openSet.Len() > 0 {
		currentNode := heap.Pop(openSet).(*models.Node)
		currentPoint := currentNode.Point

		if currentPoint == finish {
			var path []models.Point
			for currentNode != nil {
				path = append([]models.Point{currentNode.Point}, path...)
				currentNode = currentNode.Parent
			}

			return path, nil
		}

		closedSet[currentPoint] = true

		for _, dirrection := range directions {
			neighbor := models.Point{
				X: currentPoint.X + dirrection.X,
				Y: currentPoint.Y + dirrection.Y,
			}

			if !neighbor.IsValid(labyrinth) || closedSet[neighbor] {
				continue
			}

			gCost := currentNode.GCost + labyrinth[neighbor.X][neighbor.Y]
			hCost := heuristics(neighbor, finish)
			fCost := gCost + hCost

			neighborNode := &models.Node{
				Point: neighbor,
				GCost: gCost,
				HCost: hCost,
				FCost: fCost,
				Parent: currentNode,
			}

			heap.Push(openSet, neighborNode)
		}
	}

	return nil, fmt.Errorf("no path found")
}