package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var points []Point
var gridSize = 71

type Point struct {
	x, y int
}

func input() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, Point{x, y})
	}
}

func isValid(x, y int, grid [][]bool) bool {
	return x >= 0 && x < gridSize && y >= 0 && y < gridSize && !grid[y][x]
}

func findShortestPath(grid [][]bool) int {
	if !isValid(0, 0, grid) || !isValid(gridSize-1, gridSize-1, grid) {
		return -1
	}

	queue := []Point{{0, 0}}
	visited := make([][]bool, gridSize)
	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}
	visited[0][0] = true

	distance := make([][]int, gridSize)
	for i := range distance {
		distance[i] = make([]int, gridSize)
	}

	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == gridSize-1 && current.y == gridSize-1 {
			return distance[current.y][current.x]
		}

		for _, dir := range dirs {
			nextX, nextY := current.x+dir.x, current.y+dir.y
			if isValid(nextX, nextY, grid) && !visited[nextY][nextX] {
				queue = append(queue, Point{nextX, nextY})
				visited[nextY][nextX] = true
				distance[nextY][nextX] = distance[current.y][current.x] + 1
			}
		}
	}
	return -1
}

func first() {
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	numPoints := 1024
	if len(points) < numPoints {
		numPoints = len(points)
	}

	for i := 0; i < numPoints; i++ {
		p := points[i]
		grid[p.y][p.x] = true
	}

	result := findShortestPath(grid)
	fmt.Println("Part 1:", result)
}

func second() {
	for i := 0; i < len(points); i++ {

		grid := make([][]bool, gridSize)
		for j := range grid {
			grid[j] = make([]bool, gridSize)
		}

		for j := 0; j <= i; j++ {
			p := points[j]
			grid[p.y][p.x] = true
		}

		if findShortestPath(grid) == -1 {
			p := points[i]
			fmt.Printf("Part 2: %d,%d\n", p.x, p.y)
			break
		}
	}
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
