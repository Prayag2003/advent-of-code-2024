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
	const target = "XMAS"
	res := 0

	isSafe := func(r, c int) bool {
		return r >= 0 && r < H && c >= 0 && c < W
	}

	for r := 0; r < H; r++ {
		for c := 0; c < W; c++ {
			if grid[r][c] == 'X' {
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}

						yes := true
						for i := 0; i < 4; i++ {
							nr, nc := r+dr*i, c+dc*i
							if !isSafe(nr, nc) || grid[nr][nc] != target[i] {
								yes = false
								break
							}
						}
						if yes {
							res++
						}
					}
				}
			}
		}
	}

	fmt.Println(res)
}
