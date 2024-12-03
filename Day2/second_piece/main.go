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

	unsafeList, safe := parseInput(file)

	for _, unsafe := range unsafeList {
		// Try removing each level to see if it becomes safe
		for index := 0; index < len(unsafe); index++ {
			// Create a new slice without the level at 'index'
			slice := append(unsafe[:index], unsafe[index+1:]...)

			// Check if this slice is safe
			if solve(slice) == 1 {
				safe++
				break
			}
		}
	}
	print(safe + sumOfIsSafe)
}

func parseInput(file *os.File) ([][]int, int) {
	var unsafeList [][]int
	var safe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a []int
		parts := strings.Fields(scanner.Text())
		for _, part := range parts {
			x, _ := strconv.Atoi(part)
			a = append(a, x)
		}

		if solve(a) == 1 {
			safe++
		} else {
			unsafeList = append(unsafeList, a)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return unsafeList, safe
}

func isSorted(a []int) (bool, bool) {
	ascending := true
	descending := true
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			descending = false
		}
		if a[i] > a[i+1] {
			ascending = false
		}
	}
	return ascending, descending
}

func solve(a []int) int {
	ascending, descending := isSorted(a)

	if !(ascending || descending) {
		return 0
	}

	for i := 0; i < len(a)-1; i++ {
		if math.Abs(float64(a[i]-a[i+1])) < 1 || math.Abs(float64(a[i]-a[i+1])) > 3 {
			return 0
		}
	}
	return 1
}
