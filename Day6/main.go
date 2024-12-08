package main

import (
	"fmt"
)

const H = 130

var grid [H]string
var curr struct {
	first, second int
}

var directions = [][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func input() {
	for i := 0; i < H; i++ {
		fmt.Scan(&grid[i])
	}
}

func first() {
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '^' {
				curr = struct{ first, second int }{row, col}
				fmt.Printf("^ found at {%d, %d}\n", curr.first, curr.second)
				grid[row] = grid[row][:col] + "." + grid[row][col+1:]
				break
			}
		}
	}

	start := 0
	visited := make(map[struct{ first, second int }]bool)
	for {
		r2 := curr.first + directions[start][0]
		c2 := curr.second + directions[start][1]
		vis := struct{ first, second int }{r2, c2}
		visited[vis] = true

		// out of bounds
		if !(r2 >= 0 && r2 < rows && c2 >= 0 && c2 < cols) {
			break
		}

		if grid[r2][c2] != '.' {
			start = (start + 1) % 4
		} else {
			curr.first = r2
			curr.second = c2
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
