package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rows, cols int
	grid       [][]int
	directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			row[i] = int(char - '0')
		}
		grid = append(grid, row)
	}
	rows = len(grid)
	cols = len(grid[0])
}

func isBound(row, col int) bool {
	return 0 <= row && row < rows && 0 <= col && col < cols
}

func calculateScore(r, c int, countTrails bool) int {
	queue := [][]int{{r, c}}
	visited := make(map[string]int)
	visited[fmt.Sprintf("%d,%d", r, c)] = 1
	score := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currRow, currCol := curr[0], curr[1]
		currKey := fmt.Sprintf("%d,%d", currRow, currCol)

		if grid[currRow][currCol] == 9 && countTrails {
			score += visited[currKey]
		} else if grid[currRow][currCol] == 9 {
			score++
		}

		for _, dir := range directions {
			newR, newCol := currRow+dir[0], currCol+dir[1]

			if isBound(newR, newCol) && grid[newR][newCol] == grid[currRow][currCol]+1 {
				next := fmt.Sprintf("%d,%d", newR, newCol)
				if _, exists := visited[next]; exists {
					visited[next] += visited[currKey]
				} else {
					visited[next] = visited[currKey]
					queue = append(queue, []int{newR, newCol})
				}
			}
		}
	}
	return score
}

func trailHeads(countTrails bool) int {
	totalScore := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				totalScore += calculateScore(r, c, countTrails)
			}
		}
	}
	return totalScore
}

func solve() {
	fmt.Println("First: ", trailHeads(false))
	fmt.Println("Second: ", trailHeads(true))
}

func main() {
	input()
	solve()
}
