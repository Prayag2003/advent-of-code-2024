package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var diskMap string

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		diskMap = scanner.Text()
	}
}

func first() {
	encodedStr := encode(diskMap)
	compactStr := compact(encodedStr)
	checksum := checkSum(compactStr)
	fmt.Println("Original:", diskMap)
	fmt.Println("Length:", len(diskMap))
	fmt.Println("Checksum:", checksum)
	fmt.Println("Encoded String:", encodedStr)
	fmt.Println("Compacted String:", compactStr)
}

func second() {
}

func encode(diskMap string) string {
	encodedStr := ""
	fileId := 0

	for i := 0; i < len(diskMap); i++ {
		length, _ := strconv.Atoi(string(diskMap[i]))

		for j := 0; j < length; j++ {
			if i%2 == 0 {
				encodedStr += strconv.Itoa(fileId)
			} else {
				encodedStr += "."
			}
		}

		if i%2 == 0 {
			fileId++
		}
	}
	return encodedStr
}

func compact(encodedStr string) string {
	left := 0
	right := len(encodedStr) - 1

	for left <= right && right >= 0 {
		if encodedStr[right] != '.' && encodedStr[left] == '.' {
			encodedStr = encodedStr[:left] + string(encodedStr[right]) + encodedStr[left+1:right] + "." + encodedStr[right+1:]
			left++
			right--
		} else if encodedStr[right] == '.' {
			right--
		} else {
			left++
		}
	}

	return encodedStr
}

func checkSum(compactStr string) int {
	checksum := 0

	for i, char := range compactStr {
		if char == '.' {
			break
		}
		checksum += int(char-'0') * i
	}

	return checksum
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
