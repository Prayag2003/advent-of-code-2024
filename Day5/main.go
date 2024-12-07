package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pairs [][2]int
var queries [][]*int

func input() ([][2]int, [][]*int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		splits := strings.Split(line, "|")
		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])
		pairs = append(pairs, [2]int{x, y})
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		splits := strings.Split(line, ",")
		var query []*int
		for _, s := range splits {
			num, _ := strconv.Atoi(s)
			query = append(query, &num)
		}
		queries = append(queries, query)
	}
	return pairs, queries
}

func first() int {
	pairs, queries := input()
	before := make(map[int]map[int]bool)
	for _, pair := range pairs {
		if before[pair[1]] == nil {
			before[pair[1]] = make(map[int]bool)
		}
		before[pair[1]][pair[0]] = true
	}

	var ans int
	for _, query := range queries {
		ok := true
		for i, x := range query {
			for j, y := range query {
				if i < j && before[*x] != nil && before[*x][*y] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			ans += *query[len(query)/2]
		}
	}
	return ans
}

func second() int {
	before := make(map[int]map[int]bool)
	after := make(map[int]map[int]bool)

	for _, pair := range pairs {
		if before[pair[1]] == nil {
			before[pair[1]] = make(map[int]bool)
		}
		before[pair[1]][pair[0]] = true
		if after[pair[0]] == nil {
			after[pair[0]] = make(map[int]bool)
		}
		after[pair[0]][pair[1]] = true
	}

	var res int
	for _, query := range queries {
		valid := true
		for i, x := range query {
			for j, y := range query {
				if i < j && before[*x] != nil && before[*x][*y] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		var mid int
		if valid {
			mid = *query[len(query)/2]
		} else {
			deg := make(map[int]int)
			qSet := make(map[int]bool)
			for _, p := range query {
				qSet[*p] = true
				deg[*p] = 0
			}

			for _, x := range query {
				for _, y := range query {
					if before[*y] != nil && before[*y][*x] {
						deg[*x]++
					}
				}
			}

			q := []int{}
			for p := range qSet {
				if deg[p] == 0 {
					q = append(q, p)
				}
			}

			var corrOrder []int
			for len(q) > 0 {
				curr := q[0]
				q = q[1:]
				corrOrder = append(corrOrder, curr)

				for next := range after[curr] {
					deg[next]--
					if deg[next] == 0 {
						q = append(q, next)
					}
				}
			}

			if len(corrOrder) > 0 {
				mid = corrOrder[len(corrOrder)/2]
			}
		}

		res += mid
	}
	return res
}

func main() {
	fmt.Println("first :", first())
	fmt.Println("second:", second())
}
