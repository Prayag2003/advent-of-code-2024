package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	enabled := true
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		// debug: fmt.Println("Line:", line)

		commands := regexp.MustCompile(`(?:do\(\)|don't\(\)|mul\(\d+,\d+\))`).FindAllString(line, -1)

		for _, cmd := range commands {
			if doRegex.MatchString(cmd) {
				enabled = true
			} else if dontRegex.MatchString(cmd) {
				enabled = false
			} else if enabled && mulRegex.MatchString(cmd) {
				matches := mulRegex.FindStringSubmatch(cmd)
				a, _ := strconv.Atoi(matches[1])
				b, _ := strconv.Atoi(matches[2])
				ans += a * b
			}
		}
	}

	fmt.Println(ans)
}
