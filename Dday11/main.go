package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arr []string
var arr2 []string
var memo = make(map[string]int)

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	arr = strings.Split(text, " ")
	arr2 = arr
}

func first() {
	for i := 0; i < 25; i++ {
		newStones := []string{}

		for _, stone := range arr {
			if stone == "0" {
				newStones = append(newStones, "1")
				continue
			}

			if len(stone) > 0 && len(stone)%2 == 0 {
				mid := len(stone) / 2

				left := strings.TrimLeft(stone[:mid], "0")
				right := strings.TrimLeft(stone[mid:], "0")

				if left == "" {
					left = "0"
				}
				if right == "" {
					right = "0"
				}

				newStones = append(newStones, left)
				newStones = append(newStones, right)
				continue
			}

			num, _ := strconv.Atoi(stone)
			newStones = append(newStones, strconv.Itoa(num*2024))
		}

		arr = newStones
	}

	fmt.Println("First: ", len(arr))
}

func second() {
	sum := 0
	for _, stone := range arr2 {
		sum += count(stone, 25)
	}
	fmt.Println("Second:", sum)
}

func count(stone string, steps int) int {
	key := stone + ":" + strconv.Itoa(steps)
	if val, ok := memo[key]; ok {
		return val
	}

	if steps == 0 {
		return 1
	}

	if stone == "0" {
		result := count("1", steps-1)
		memo[key] = result
		return result
	}

	length := len(stone)
	if length%2 == 0 {
		mid := length / 2
		left := count(stone[:mid], steps-1)
		right := count(stone[mid:], steps-1)
		result := left + right
		memo[key] = result
		return result
	}

	temp, _ := strconv.Atoi(stone)
	result := count(strconv.Itoa(temp*2024), steps-1)
	memo[key] = result
	return result
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
