package main

import (
	"fmt"
)

var prog []int
var rega, regb, regc int

func input() {
	rega = 63281501
	regb = 0
	regc = 0
	arr := []int{2, 4, 1, 5, 7, 5, 4, 5, 0, 3, 1, 6, 5, 5, 3, 0}
	i := 0

	for i < len(arr) {
		prog = append(prog, arr[i])
		i++
	}
}

func getCombo(oper, rega, regb, regc int) int {
	switch oper {
	case 0, 1, 2, 3:
		return oper
	case 4:
		return rega
	case 5:
		return regb
	case 6:
		return regc
	default:
		return 0
	}
}

func run(prog []int, rega, regb, regc int) []int {
	ip := 0
	out := []int{}

	for ip < len(prog) {
		oper := prog[ip+1]
		combo := getCombo(oper, rega, regb, regc)

		switch prog[ip] {
		case 0:
			rega /= 1 << uint(combo)
		case 1:
			regb ^= oper
		case 2:
			regb = combo % 8
		case 3:
			if rega != 0 {
				ip = oper
				continue
			}
		case 4:
			regb ^= regc
		case 5:
			out = append(out, combo%8)
		case 6:
			regb = rega / (1 << uint(combo))
		case 7:
			regc = rega / (1 << uint(combo))
		}
		ip += 2
	}
	return out
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func first() {
	result := run(prog, rega, regb, regc)
	fmt.Print("Part 1: ")
	for i, val := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func second() {
	rega := 0
	j := 1
	istart := 0

	for j <= len(prog) && j >= 0 {
		rega <<= 3
		foundMatch := false

		for i := istart; i < 8; i++ {
			suffix := prog[len(prog)-j:]
			result := run(prog, rega+i, regb, regc)

			if slicesEqual(suffix, result) {
				rega += i
				j++
				istart = 0
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			j--
			rega >>= 3
			istart = rega%8 + 1
			rega >>= 3
		}
	}

	fmt.Println("Part 2:", rega)
}

func solve() {
	first()
	second()
}

func main() {
	input()
	solve()
}
