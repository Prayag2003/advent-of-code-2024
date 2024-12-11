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

	fmt.Println("First:", len(arr))
}

func solve2(x string, t int) int {
	key := x + "_" + strconv.Itoa(t)
	if val, ok := memo[key]; ok {
		return val
	}

	if t == 0 {
		return 1
	}

	if x == "0" {
		return solve2("1", t-1)
	}

	if len(x)%2 == 0 {
		mid := len(x) / 2
		left := strings.TrimLeft(x[:mid], "0")
		right := strings.TrimLeft(x[mid:], "0")

		if left == "" {
			left = "0"
		}
		if right == "" {
			right = "0"
		}

		ret := solve2(left, t-1) + solve2(right, t-1)
		memo[key] = ret
		return ret
	}

	num, _ := strconv.Atoi(x)
	ret := solve2(strconv.Itoa(num*2024), t-1)
	memo[key] = ret
	return ret
}

func second() {
	memo = make(map[string]int)

	total := 0
	for _, x := range arr2 {
		total += solve2(x, 75)
	}

	fmt.Println("Second:", total)
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
