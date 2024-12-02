package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	left, right := parseInput(file)

	fmt.Println(score(left, right))
}

func parseInput(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var left, right []int

	// read each line & split into two integers
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing integers:", err1, err2)
				continue
			}

			left = append(left, num1)
			right = append(right, num2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return left, right
}

func score(left []int, right []int) int {

	score := 0

	for _, num := range left {
		freq := 0

		for _, val := range right {
			if val == num {
				freq++
			}
		}
		score += num * freq
	}
	return score
}
