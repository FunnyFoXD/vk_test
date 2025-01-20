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

func heuristics(a, b models.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

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