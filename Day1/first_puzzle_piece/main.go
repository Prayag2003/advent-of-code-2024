package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	a, b := parseInput(file)

	// sort
	sort.Ints(a)
	sort.Ints(b)

	// sum of abs diffs
	sum := calculateSum(a, b)
	fmt.Println(sum)
}

func parseInput(file *os.File) (a, b []int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			a = append(a, x)
			b = append(b, y)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return
}

func calculateSum(a, b []int) int64 {
	var sum int64
	for i := 0; i < len(a); i++ {
		sum += int64(math.Abs(float64(a[i] - b[i])))
	}
	return sum
}
