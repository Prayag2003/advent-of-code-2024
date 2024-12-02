package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sumOfIsSafe int = 0

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	parseInput(file)
	fmt.Println("Sum of isSafe:", sumOfIsSafe)
}

func parseInput(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a []int
		parts := strings.Fields(scanner.Text())
		for _, part := range parts {
			x, _ := strconv.Atoi(part)
			a = append(a, x)
		}

		sumOfIsSafe += solve(a)
		a = []int{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func isSortedDesc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			return false
		}
	}
	return true
}

func isSortedAsc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func solve(a []int) int {
	if !(isSortedAsc(a) || isSortedDesc(a)) {
		return 0
	}
	if isSortedAsc(a) || isSortedDesc(a) {
		for i := 0; i < len(a)-1; i++ {
			if !(math.Abs(float64(a[i]-a[i+1])) >= 1 && math.Abs(float64(a[i]-a[i+1])) <= 3) {
				return 0
			}
		}
	}
	return 1
}
