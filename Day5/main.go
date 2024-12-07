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
	beforeMap := make(map[int]map[int]bool)
	for _, pair := range pairs {
		if beforeMap[pair[1]] == nil {
			beforeMap[pair[1]] = make(map[int]bool)
		}
		beforeMap[pair[1]][pair[0]] = true
	}

	var ans int
	for _, query := range queries {
		ok := true
		for i, x := range query {
			for j, y := range query {
				if i < j && beforeMap[*x] != nil && beforeMap[*x][*y] {
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
	beforeMap := make(map[int]map[int]bool)
	afterMap := make(map[int]map[int]bool)

	// Build before and after relationship maps
	for _, pair := range pairs {
		if beforeMap[pair[1]] == nil {
			beforeMap[pair[1]] = make(map[int]bool)
		}
		beforeMap[pair[1]][pair[0]] = true

		if afterMap[pair[0]] == nil {
			afterMap[pair[0]] = make(map[int]bool)
		}
		afterMap[pair[0]][pair[1]] = true
	}

	var ans int
	for _, query := range queries {
		ok := true
		for i, x := range query {
			for j, y := range query {
				if i < j && beforeMap[*x] != nil && beforeMap[*x][*y] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}

		var middlePage int
		if ok {
			middlePage = *query[len(query)/2]
		} else {
			querySet := make(map[int]bool)
			degrees := make(map[int]int)
			for _, page := range query {
				querySet[*page] = true
				degrees[*page] = 0
			}

			// Build graph and in-degree map
			for _, x := range query {
				for _, y := range query {
					if beforeMap[*y] != nil && beforeMap[*y][*x] {
						degrees[*x]++
					}
				}
			}

			// Topological sort using Kahn's algorithm
			queue := []int{}
			for page := range querySet {
				if degrees[page] == 0 {
					queue = append(queue, page)
				}
			}

			var good []int
			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				good = append(good, curr)

				for next := range afterMap[curr] {
					degrees[next]--
					if degrees[next] == 0 {
						queue = append(queue, next)
					}
				}
			}

			if len(good) > 0 {
				middlePage = good[len(good)/2]
			}
		}

		ans += middlePage
	}
	return ans
}

func main() {
	fmt.Println("first res:", first())
	fmt.Println("second res:", second())
}
