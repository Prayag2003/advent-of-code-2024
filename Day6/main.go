package main

import (
	"fmt"
	"strings"
)

const H = 130

var (
	grid []string
	curr struct {
		row, col int
	}
	directions = [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
)

type visitedSet map[string]bool

func (vs visitedSet) add(row, col int) {
	key := fmt.Sprintf("%d,%d", row, col)
	vs[key] = true
}

func (vs visitedSet) contains(row, col int) bool {
	key := fmt.Sprintf("%d,%d", row, col)
	return vs[key]
}

func input() {
	grid = make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scanln(&grid[i])
	}
}

func first() {
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '^' {
				curr.row = row
				curr.col = col
				fmt.Printf("^ found at {%d, %d}\n", curr.row, curr.col)
				grid[row] = strings.Replace(grid[row], "^", ".", 1)
				break
			}
		}
	}

	start := 0
	visited := make(visitedSet)

	for {
		r2 := curr.row + directions[start][0]
		c2 := curr.col + directions[start][1]
		visited.add(curr.row, curr.col)

		if r2 < 0 || r2 >= rows || c2 < 0 || c2 >= cols {
			break
		}

		if grid[r2][c2] != '.' {
			start = (start + 1) % 4
		} else {
			curr.row = r2
			curr.col = c2
		}
	}

	fmt.Printf("Size of the visited map is: %d\n", len(visited))
}

func second() {
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
