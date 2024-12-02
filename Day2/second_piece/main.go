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
		for index := 0; index < len(unsafe); index++ {
			slice := append(unsafe[:index], unsafe[index+1:]...)
			for i := 0; i < len(slice); i++ {
				fmt.Print(slice[i], " ")
			}
			fmt.Println()
			if solve(slice) == 1 {
				safe++
			}
		}
	}
	print(safe + sumOfIsSafe)
}

func parseInput(file *os.File) ([][]int, int) {

	var unsafeList [][]int
	var safe int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a []int
		parts := strings.Fields(scanner.Text())
		for _, part := range parts {
			x, _ := strconv.Atoi(part)
			a = append(a, x)
		}

		sumOfIsSafe = solve(a)
		if sumOfIsSafe == 0 {
			unsafeList = append(unsafeList, a)
		} else {
			safe += 1
		}
		a = []int{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return unsafeList, safe
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
