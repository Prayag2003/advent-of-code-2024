package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sections []string

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	sections = strings.Split(input, "\n\n")
}

func first() {
	intRegex := regexp.MustCompile(`-?\d+`)
	total := float64(0)

	for _, section := range sections {
		matches := intRegex.FindAllString(section, -1)

		ax, _ := strconv.Atoi(matches[0])
		ay, _ := strconv.Atoi(matches[1])
		bx, _ := strconv.Atoi(matches[2])
		by, _ := strconv.Atoi(matches[3])
		px, _ := strconv.Atoi(matches[4])
		py, _ := strconv.Atoi(matches[5])

		curr := math.Inf(1)
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if ax*i+bx*j == px && ay*i+by*j == py {
					curr = math.Min(curr, float64(i*3+j))
				}
			}
		}

		if curr != math.Inf(1) {
			total += curr
		}
	}
	fmt.Println("First -> Min cost is:", total)
}

func second() {
	total := 0
	offset := 10000000000000

	intRegex := regexp.MustCompile(`-?\d+`)

	for _, section := range sections {
		matches := intRegex.FindAllString(section, -1)

		ax, _ := strconv.Atoi(matches[0])
		ay, _ := strconv.Atoi(matches[1])
		bx, _ := strconv.Atoi(matches[2])
		by, _ := strconv.Atoi(matches[3])
		x, _ := strconv.Atoi(matches[4])
		y, _ := strconv.Atoi(matches[5])

		x += offset
		y += offset

		denominator := float64(bx*ay - by*ax)
		if denominator == 0 {
			continue
		}

		A := (float64(bx*y) - float64(by*x)) / denominator
		B := (float64(x) - float64(ax)*A) / float64(bx)

		if float64(int(A)) == A && float64(int(B)) == B {
			total += int(3*A + B)
		}
	}
	fmt.Println("Second -> Min cost is:", total)
}

func main() {
	input()
	first()
	second()
}
