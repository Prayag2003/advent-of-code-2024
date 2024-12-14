package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	field      [][]byte
	rows, cols int
)

type Region struct {
	area      int
	perimeter int
	corners   int
	plots     map[[2]int]bool
}

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		field = append(field, row)
	}
	rows = len(field)
	cols = len(field[0])
}

func directions() ([]int, []int) {
	return []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
}

func isValid(row, col int) bool {
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func explore(startRow, startCol int, plantType byte, visited [][]bool) *Region {
	region := &Region{
		area:      0,
		perimeter: 0,
		corners:   0,
		plots:     make(map[[2]int]bool),
	}

	dx, dy := directions()
	queue := [][2]int{{startRow, startCol}}
	region.plots[[2]int{startRow, startCol}] = true
	visited[startRow][startCol] = true
	region.area++

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nRow, nCol := curr[0]+dx[i], curr[1]+dy[i]
			if isValid(nRow, nCol) {
				if field[nRow][nCol] == plantType {
					if !visited[nRow][nCol] {
						queue = append(queue, [2]int{nRow, nCol})
						visited[nRow][nCol] = true
						region.plots[[2]int{nRow, nCol}] = true
						region.area++
					}
				} else {
					region.perimeter++
				}
			} else {
				region.perimeter++
			}
		}
	}
	return region
}

func findAllRegions() []*Region {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var regions []*Region
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if !visited[row][col] && field[row][col] != '.' {
				region := explore(row, col, field[row][col], visited)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func first() int {
	totalCost := 0
	regions := findAllRegions()
	for _, region := range regions {
		totalCost += region.area * region.perimeter
	}
	return totalCost
}

func second() int {
	totalCost := 0
	regions := findAllRegions()

	for _, region := range regions {
		corners := 0
		visitedCorners := make(map[[2]int]bool)

		for plot := range region.plots {
			row, col := plot[0], plot[1]

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					neighbor := [2]int{row + dx, col + dy}

					if !region.plots[neighbor] && !visitedCorners[neighbor] {
						visitedCorners[neighbor] = true
						corners++
					}
				}
			}
		}

		totalCost += corners * region.area
	}

	return totalCost
}

func solve() {
	fmt.Printf("First: %d\n", first())
	fmt.Printf("Second: %d\n", second())
}

func main() {
	input()
	solve()
}
