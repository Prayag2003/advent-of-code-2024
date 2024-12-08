package main

import (
	"fmt"
	"os"
)

var antMap [][2]interface{}
var totalMap [][]rune
var freqs = make(map[string]struct{})
var al = make(map[[2]int]struct{})
var al2 = make(map[[2]int]struct{})

func input() {
	file, _ := os.Open("input.txt")

	var y int
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}

		var chari []rune
		for x, char := range line {
			if char != '.' && char != '\n' {
				freqs[string(char)] = struct{}{}
				antMap = append(antMap, [2]interface{}{string(char), [2]int{x, y}})
			}
			chari = append(chari, char)
		}
		totalMap = append(totalMap, chari)
		y++
	}

}

func slope(antA, antB [2]int) (int, int) {
	x1, y1 := antA[0], antA[1]
	x2, y2 := antB[0], antB[1]
	return abs(y2 - y1), abs(x2 - x1)
}

func first(antA, antB [2]int) ([2]int, [2]int) {
	rise, run := slope(antA, antB)
	x1, y1 := antA[0], antA[1]
	x2, y2 := antB[0], antB[1]

	if x1 > x2 {
		x1 += run
		x2 -= run
	} else if x1 < x2 {
		x1 -= run
		x2 += run
	}

	if y1 < y2 {
		y1 -= rise
		y2 += rise
	}

	return [2]int{x1, y1}, [2]int{x2, y2}
}

func second(antA, antB [2]int, xLimit, yLimit int) [][2]int {
	rise, run := slope(antA, antB)
	x1, y1 := antA[0], antA[1]
	x2, y2 := antB[0], antB[1]

	antinodes := [][2]int{{x1, y1}, {x2, y2}}
	inBounds, inBounds2 := true, true

	for inBounds || inBounds2 {
		if x1 > x2 {
			x1 += run
			x2 -= run
		} else if x1 < x2 {
			x1 -= run
			x2 += run
		}

		if y1 < y2 {
			y1 -= rise
			y2 += rise
		}

		if x1 < 0 || x1 > xLimit || y1 < 0 || y1 > yLimit {
			inBounds = false
		} else {
			antinodes = append(antinodes, [2]int{x1, y1})
		}

		if x2 < 0 || x2 > xLimit || y2 < 0 || y2 > yLimit {
			inBounds2 = false
		} else {
			antinodes = append(antinodes, [2]int{x2, y2})
		}
	}

	return antinodes
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve() {
	for f := range freqs {
		var antennas [][2]int
		for _, a := range antMap {
			if a[0] == f {
				antennas = append(antennas, a[1].([2]int))
			}
		}

		for i, antA := range antennas {
			for b := i + 1; b < len(antennas); b++ {
				antinodeA, antinodeB := first(antA, antennas[b])
				al[antinodeA] = struct{}{}
				al[antinodeB] = struct{}{}
			}
		}
	}

	for f := range freqs {
		var antennas [][2]int
		for _, a := range antMap {
			if a[0] == f {
				antennas = append(antennas, a[1].([2]int))
			}
		}

		for i, antA := range antennas {
			for b := i + 1; b < len(antennas); b++ {
				antinodes := second(antA, antennas[b], len(totalMap[0]), len(totalMap))
				for _, antis := range antinodes {
					al2[antis] = struct{}{}
				}
			}
		}
	}

	part1, part2 := 0, 0
	for antinode := range al {
		x, y := antinode[0], antinode[1]
		if x >= 0 && x < len(totalMap[0]) && y >= 0 && y < len(totalMap) {
			part1++
			totalMap[y][x] = '#'
		}
	}

	for antinode := range al2 {
		x, y := antinode[0], antinode[1]
		if x >= 0 && x < len(totalMap[0]) && y >= 0 && y < len(totalMap) {
			part2++
			totalMap[y][x] = '#'
		}
	}

	fmt.Printf("Part 1: %d \n", part1)
	fmt.Printf("Part 2: %d ", part2)
}

func main() {
	input()
	solve()
}
