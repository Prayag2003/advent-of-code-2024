package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var patterns []string
var designs []string

func input() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		patterns = strings.Split(line, ", ")
	}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			designs = append(designs, line)
		}
	}
}

func countWays(design string, memo map[string]int64) int64 {
	if len(design) == 0 {
		return 1
	}

	if count, exists := memo[design]; exists {
		return count
	}

	var total int64 = 0
	for _, pattern := range patterns {
		if len(pattern) <= len(design) && design[:len(pattern)] == pattern {
			total += countWays(design[len(pattern):], memo)
		}
	}

	memo[design] = total
	return total
}

func first() {
	possibleCount := 0

	for _, design := range designs {
		if countWays(design, make(map[string]int64)) > 0 {
			possibleCount++
		}
	}

	fmt.Printf("Part 1 - Number of possible designs: %d\n", possibleCount)
}

func second() {
	var totalWays int64 = 0

	for _, design := range designs {
		ways := countWays(design, make(map[string]int64))
		if ways > 0 {
			totalWays += ways
		}
	}

	fmt.Printf("Part 2 - Total number of different ways: %d\n", totalWays)
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
