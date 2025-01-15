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

func readInput() (int, int, [][]int, models.Point, models.Point, error) {
	var n, m int
	_, err := fmt.Scanf("%d %d\n", &n, &m)
	if err != nil {
		return 0, 0, nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(labirinth size): %w", err)
	}

	reader := bufio.NewReader(os.Stdin)
	labyrinth := make([][]int, n)
	for i := 0; i < n; i++ {
		row, err := reader.ReadString('\n')
		if err != nil {
			return 0, 0, nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(row in %d): %w", i ,err)
		}
		
		row = strings.TrimSpace(row)
		cells := strings.Fields(row)
		labyrinth[i] = make([]int, m)
		
		for j := 0; j < m; j++ {
			labyrinth[i][j], err = strconv.Atoi(cells[j])
			if err != nil {
				return 0, 0, nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(cell in %d, %d): %w", i, j, err)
			}
		}
	}

	var startX, startY int
	var finishX, finishY int
	_, err = fmt.Scanf("%d %d %d %d\n", &startX, &startY, &finishX, &finishY)
	if err != nil {
		return 0, 0, nil, models.Point{}, models.Point{}, fmt.Errorf("failed to read input(start or finish points): %w", err)
	}

	start := models.Point{
		X: startX,
		Y: startY,
	}
	finish := models.Point{
		X: finishX,
		Y: finishY,
	}

	return n, m, labyrinth, start, finish, nil
}

func main() {
	n, m, labirinth, start, finish, err := readInput()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	path, err := shortestpath.ShortestPath(n, m, labirinth, start, finish)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.X, p.Y)
	}
	fmt.Println(".") 
}