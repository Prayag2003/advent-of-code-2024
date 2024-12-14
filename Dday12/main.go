package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	field      []string
	rows, cols int
)

type Region struct {
	area      int
	peri      int
	plantType byte
	plots     map[[2]int]bool
}

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		field = append(field, scanner.Text())
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

func explore(startRow, startCol int, plantType byte, vis [][]bool) *Region {

	region := &Region{
		area:      0,
		peri:      0,
		plantType: plantType,
		plots:     make(map[[2]int]bool),
	}

	dx, dy := directions()

	queue := [][2]int{{startRow, startCol}}
	region.plots[[2]int{startRow, startCol}] = true
	vis[startRow][startCol] = true
	region.area++

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nR, nC := curr[0]+dx[i], curr[1]+dy[i]

			if isValid(nR, nC) {
				if field[nR][nC] == plantType {

					if !vis[nR][nC] {
						queue = append(queue, [2]int{nR, nC})
						vis[nR][nC] = true
						region.plots[[2]int{nR, nC}] = true
						region.area++
					}
				} else {

					region.peri++
				}
			} else {
				region.peri++
			}
		}
	}
	return region
}

func findAllRegions() []*Region {

	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}

	var regions []*Region
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if !vis[row][col] {
				regions = append(regions, explore(row, col, field[row][col], vis))
			}
		}
	}
	return regions
}

func first() int {

	totalPrice := 0
	regions := findAllRegions()

	for _, region := range regions {
		totalPrice += region.area * region.peri
	}
	return totalPrice
}

func countRegionSides(region *Region) int {
	sides := 0
	dx, dy := directions()

	for plot := range region.plots {
		row, col := plot[0], plot[1]

		for i := 0; i < 4; i++ {
			newRow, newCol := row+dx[i], col+dy[i]

			if !isValid(newRow, newCol) || field[newRow][newCol] != region.plantType {
				sides++
			}
		}
	}
	return sides
}

func second() int {
	totalPrice := 0
	regions := findAllRegions()

	for _, region := range regions {
		sides := countRegionSides(region)
		totalPrice += region.area * sides
	}

	return totalPrice
}

func solve() {
	fmt.Printf("First: %d\n", first())
	fmt.Printf("Second: %d\n", second())
}

func main() {
	input()
	solve()
}
