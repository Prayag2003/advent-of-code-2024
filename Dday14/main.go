package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const (
	cols = 101
	rows = 103
)

var quadrants = [4]int{0, 0, 0, 0}
var re = regexp.MustCompile(`-?\d+`)
var robots [][4]int

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		if len(matches) >= 4 {
			var robot [4]int
			for i := 0; i < 4; i++ {
				robot[i], _ = strconv.Atoi(matches[i])
			}
			robots = append(robots, robot)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}

func first() {
	for _, robot := range robots {
		x := (robot[0] + (100 * robot[2])) % cols
		y := (robot[1] + (100 * robot[3])) % rows

		if x < 0 {
			x += cols
		}
		if y < 0 {
			y += rows
		}

		if x < 50 && y < 51 {
			quadrants[0]++
		} else if x > 50 && y < 51 {
			quadrants[1]++
		} else if x < 50 && y > 51 {
			quadrants[2]++
		} else if x > 50 && y > 51 {
			quadrants[3]++
		}
	}
	result := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	fmt.Println(result)
}

func second() {
	for i := 0; i < 1000000; i++ {
		for j := 0; j < rows; j++ {
			row := []int{}
			for _, robot := range robots {
				if robot[1] == j {
					row = append(row, robot[0])
				}
			}
			sort.Ints(row)
			n := 0
			for k := 0; k < len(row)-1; k++ {
				if row[k+1] == row[k]+1 {
					n++
				} else {
					n = 0
				}
				if n > 10 {
					fmt.Println(i)
					return
				}
			}
		}
		for idx := range robots {
			robots[idx][0] = (robots[idx][0] + robots[idx][2]) % cols
			robots[idx][1] = (robots[idx][1] + robots[idx][3]) % rows
			if robots[idx][0] < 0 {
				robots[idx][0] += cols
			}
			if robots[idx][1] < 0 {
				robots[idx][1] += rows
			}
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
