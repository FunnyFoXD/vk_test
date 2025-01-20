package main

import (
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"os"

	"githab.com/funnyfoxd/vk_test/models"
	"githab.com/funnyfoxd/vk_test/shortestPath"
)

// readInput reads the input from stdin and returns labyrinth and start and finish points.
// labyrinth is represented as [][]int, where 0 represents empty cell and 1 represents wall.
// Start and finish points are represented as two models.Point.
// If there is an error while reading input, it will be returned as error.
// The input should be in the following format:
// First line should contain two numbers: n and m, where n is the height of the labyrinth
// and m is the width.
// Next n lines should contain m numbers each, representing the labyrinth.
// The last line should contain four numbers: x and y of start point and x and y of finish point.
func readInput() ([][]int, models.Point, models.Point, error) {
	var n, m int
	_, err := fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		return nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(labyrinth size): %w", err)
	}

	reader := bufio.NewReader(os.Stdin)
	labyrinth := make([][]int, n)
	for i := 0; i < n; i++ {
		row, err := reader.ReadString('\n')
		if err != nil {
			return nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(row in %d): %w", i ,err)
		}
		
		row = strings.TrimSpace(row)
		cells := strings.Fields(row)
		labyrinth[i] = make([]int, m)
		
		for j := 0; j < m; j++ {
			labyrinth[i][j], err = strconv.Atoi(cells[j])
			if err != nil {
				return nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(cell in %d, %d): %w", i, j, err)
			}
		}
	}

	var startX, startY int
	var finishX, finishY int
	_, err = fmt.Scanf("%d %d %d %d\n", &startX, &startY, &finishX, &finishY)
	if err != nil {
		return nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(start or finish points): %w", err)
	}

	start := models.Point{
		X: startX,
		Y: startY,
	}
	finish := models.Point{
		X: finishX,
		Y: finishY,
	}

	return labyrinth, start, finish, nil
}

// main reads labyrinth and start and finish points from stdin and prints shortest path
// from start to finish using A* algorithm to stdout. If there is an error while reading
// input or finding the path, it will be printed to stderr and program will exit with
// non-zero status code.
func main() {
	labyrinth, start, finish, err := readInput()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	path, err := shortestpath.AStar(labyrinth, start, finish)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.X, p.Y)
	}
	fmt.Println(".") 
}