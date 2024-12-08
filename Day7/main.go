package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const H = 850

var toEval = make(map[int][]int)

func input() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")

		key, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		operands := strings.Fields(strings.TrimSpace(split[1]))

		var values []int
		for _, s := range operands {
			num, _ := strconv.Atoi(s)
			values = append(values, num)
		}
		toEval[key] = values
	}
}

func first() {
	ans := 0

	for target, nums := range toEval {
		if sol(target, nums, 0, nums[0], 0) {
			ans += target
		}
	}
	print("First: ", ans)
}

func second() {
	ans := 0

	for target, nums := range toEval {
		if sol(target, nums, 0, nums[0], 1) {
			ans += target
		}
	}
	print("\nSecond: ", ans)
}

func sol(target int, nums []int, idx int, curr int, flag int) bool {
	if idx == len(nums)-1 {
		return curr == target
	}

	// +
	if sol(target, nums, idx+1, curr+nums[idx+1], flag) {
		return true
	}

	// *
	if sol(target, nums, idx+1, curr*nums[idx+1], flag) {
		return true
	}

	// ||
	if flag == 1 {
		concat := concate(curr, nums[idx+1])
		if sol(target, nums, idx+1, concat, flag) {
			return true
		}
	}
	return false
}

func concate(a, b int) int {
	return toInt(toString(a) + toString(b))
}

func toString(n int) string {
	return strconv.Itoa(n)
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
