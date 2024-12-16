package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	grid     []string
	movement string
	rows     int
	cols     int
)

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		grid = append(grid, line)
	}

	rows = len(grid)
	cols = len(grid[0])

	for scanner.Scan() {
		movement += scanner.Text()
	}
}

func input2() {
	scanner := bufio.NewScanner(os.Stdin)
	var expansion = map[string]string{
		"#": "##",
		"@": "@.",
		".": "..",
		"O": "[]",
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var str string
		for _, ele := range line {
			str += strings.Split(expansion[string(ele)], "")[0] + strings.Split(expansion[string(ele)], "")[1]
		}
		grid = append(grid, str)
	}

	rows = len(grid)
	cols = len(grid[0])

	for scanner.Scan() {
		movement += scanner.Text()
	}
}

func first() {
	var bot_x, bot_y int
	var dr, dc int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '@' {
				bot_x = row
				bot_y = col
				break
			}
		}
	}

	for _, move := range movement {
		switch move {
		case '^':
			dr, dc = -1, 0
		case '>':
			dr, dc = 0, 1
		case 'v':
			dr, dc = 1, 0
		case '<':
			dr, dc = 0, -1
		}

		currentR := bot_x
		currentC := bot_y
		targets := [][2]int{{bot_x, bot_y}}
		pass := true

		for 1 == 1 {
			currentR += dr
			currentC += dc

			curr := grid[currentR][currentC]

			if curr == '#' {
				pass = false
				break
			}
			if curr == 'O' {
				targets = append(targets, [2]int{currentR, currentC})
			}
			if curr == '.' {
				break
			}
		}

		if !pass {
			continue
		}

		// robots vacate its position
		rowRunes := []rune(grid[bot_x])
		rowRunes[bot_y] = '.'
		grid[bot_x] = string(rowRunes)

		rowRunes = []rune(grid[bot_x+dr])
		rowRunes[bot_y+dc] = '@'
		grid[bot_x+dr] = string(rowRunes)

		for i := 0; i < len(targets); i++ {
			if i == 0 {
				continue
			}
			rowRunes = []rune(grid[targets[i][0]+dr])
			rowRunes[targets[i][1]+dc] = 'O'
			grid[targets[i][0]+dr] = string(rowRunes)
		}

		bot_x += dr
		bot_y += dc
	}

	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'O' {
				sum += 100*row + col
			}
		}
	}
	print("Sum is :", sum)
}

func contains(targets [][2]int, r, c int) bool {
	for _, target := range targets {
		if target[0] == r && target[1] == c {
			return true
		}
	}
	return false
}

func second() {
	var bot_x, bot_y int
	var dr, dc int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '@' {
				bot_x = row
				bot_y = col
				break
			}
		}
	}

	for _, move := range movement {
		switch move {
		case '^':
			dr, dc = -1, 0
		case '>':
			dr, dc = 0, 1
		case 'v':
			dr, dc = 1, 0
		case '<':
			dr, dc = 0, -1
		}

		targets := [][2]int{{bot_x, bot_y}}
		pass := true

		newR, newC := 0, 0

		for _, val := range targets {
			newR = val[0] + dr
			newC = val[1] + dc

			if contains(targets, newR, newC) {
				continue
			}

			if newR >= rows {
				grid = append(grid, strings.Repeat(".", cols))
				rows++
			}
			if newC >= cols {
				for i := 0; i < rows; i++ {
					grid[i] += "."
				}
				cols++
			}

			curr := grid[newR][newC]
			if curr == '#' {
				pass = false
				break
			}
			if curr == '[' {
				targets = append(targets, [2]int{newR, newC})
				targets = append(targets, [2]int{newR, newC + 1})
			}
			if curr == ']' {
				targets = append(targets, [2]int{newR, newC})
				targets = append(targets, [2]int{newR, newC - 1})
			}
		}

		if !pass {
			continue
		}

		copyGrid := make([]string, len(grid))
		copy(copyGrid, grid)

		// Robot vacates its position
		rowRunes := []rune(grid[bot_x])
		rowRunes[bot_y] = '.'
		grid[bot_x] = string(rowRunes)

		rowRunes = []rune(grid[bot_x+dr])
		rowRunes[bot_y+dc] = '@'
		grid[bot_x+dr] = string(rowRunes)

		for i := 1; i < len(targets); i++ {
			rowRunes := []rune(grid[targets[i][0]])
			rowRunes[targets[i][1]] = '.'
			grid[targets[i][0]] = string(rowRunes)
		}

		for i := 1; i < len(targets); i++ {
			rowRunes := []rune(grid[targets[i][0]+dr])
			rowRunes[targets[i][1]+dc] = rune(copyGrid[targets[i][0]+dr][targets[i][1]+dc])
			grid[targets[i][0]+dr] = string(rowRunes)
		}

		bot_x += dr
		bot_y += dc
	}

	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'O' {
				sum += 100*row + col
			}
		}
	}
	fmt.Println("Sum is:", sum)
}

func solve() {
	// first()
	second()
}

func main() {
	// input()
	input2()
	solve()
}
