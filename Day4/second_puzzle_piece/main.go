package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const H = 140
	var grid [H]string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < H; i++ {
		scanner.Scan()
		grid[i] = scanner.Text()
	}

	W := len(grid[0])
	count := 0

	dirs := [4][2]int{
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for r := 1; r < H-1; r++ {
		for c := 1; c < W-1; c++ {
			if grid[r][c] == 'A' {
				s := ""
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					s += string(grid[nr][nc])
				}

				if s == "SMMS" || s == "MMSS" || s == "SSMM" || s == "MSSM" {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
