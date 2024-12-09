package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func makeFilesystem(disk string) []interface{} {
	blocks := []interface{}{}
	isFile := true
	id := 0

	for _, char := range disk {
		x, _ := strconv.Atoi(string(char))
		if isFile {
			for j := 0; j < x; j++ {
				blocks = append(blocks, id)
			}
			id++
			isFile = false
		} else {
			for j := 0; j < x; j++ {
				blocks = append(blocks, nil)
			}
			isFile = true
		}
	}
	return blocks
}

func move(arr []interface{}) []interface{} {
	firstFree := 0
	for arr[firstFree] != nil {
		firstFree++
	}

	i := len(arr) - 1
	for arr[i] == nil {
		i--
	}

	for i > firstFree {
		arr[firstFree] = arr[i]
		arr[i] = nil
		for arr[i] == nil {
			i--
		}
		for arr[firstFree] != nil {
			firstFree++
		}
	}
	return arr
}

func checksum(arr []interface{}) int {
	ans := 0
	for i, x := range arr {
		if x != nil {
			ans += i * x.(int)
		}
	}
	return ans
}

func solve() {
	diskMap := input()
	diskMap = strings.TrimSpace(diskMap)

	filesystem := makeFilesystem(diskMap)
	finalState := move(filesystem)
	result := checksum(finalState)

	fmt.Println("Length:", len(diskMap))
	fmt.Println("Checksum:", result)
}

func main() {
	solve()
}
