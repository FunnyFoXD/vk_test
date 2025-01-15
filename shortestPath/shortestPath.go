package shortestpath

import (
	"container/heap"
	"fmt"

	"githab.com/funnyfoxd/vk_test/models"
	"githab.com/funnyfoxd/vk_test/prQ"
)

var directions = []models.Point{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
}

func isValidDir(x, y, n, m int) bool {
	return x >= 0 && x < n &&
		y >= 0 && y < m
}

func ShortestPath(n, m int, labirinth [][]int, start, finish models.Point) ([]models.Point, error) {
	pq := &prq.PriorityQueue{}
	heap.Init(pq)

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = int(^uint(0) >> 1)
		}
	}
	dist[start.X][start.Y] = 0

	previous := make([][]models.Point, n)
	for i := range previous {
		previous[i] = make([]models.Point, m)
	}

	heap.Push(pq, models.Point{X: start.X, Y: start.Y, Cost: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(models.Point)

		if current.X == finish.X && current.Y == finish.Y {
			path := []models.Point{}
			for current.X != start.X || current.Y != start.Y {
				path = append([]models.Point{current}, path...)
				current = previous[current.X][current.Y]	
			}
			path = append([]models.Point{start}, path...)

			return path, nil
		}

		for _, dir := range directions {
			nextX, nextY := current.X + dir.X, current.Y + dir.Y
			if isValidDir(nextX, nextY, n, m) && labirinth[nextX][nextY] != 0 {
				newCost := current.Cost + labirinth[nextX][nextY]
				if newCost < dist[nextX][nextY] {
					dist[nextX][nextY] = newCost
					previous[nextX][nextY] = current
					heap.Push(pq, models.Point{X: nextX, Y: nextY, Cost: newCost})
				}
			}
		}
	}

	return nil, fmt.Errorf("no path in labirinth")
}