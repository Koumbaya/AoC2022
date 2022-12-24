package day10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

//go:embed testmin.txt
var testmin string

func Run() {
	fmt.Println(part1(input))
}

func part1(in string) int {
	reg := 1 // register starts at 1
	inst := make([]int, 0)
	strengths := 0
	for _, line := range strings.Split(in, "\n") {
		if line == "noop" {
			inst = append(inst, 0)
		} else {
			spl := strings.Split(line, " ")
			val, _ := strconv.Atoi(spl[1])
			inst = append(inst, 0)
			inst = append(inst, val)
		}
	}

	for i, val := range inst {
		cycle := i + 1

		if special(cycle) {
			strengths += cycle * reg
		}
		if (cycle-1)%40 == 0 {
			fmt.Println()
		}
		lightDark(reg, cycle) // part 2
		reg += val
	}

	return strengths
}

func lightDark(x, cycle int) {
	cycle = cycle - 1
	if (x == cycle%40) ||
		(x+1 == cycle%40) ||
		(x-1 == cycle%40) {
		fmt.Print("▇")
	} else {
		fmt.Print(" ")
	}
}

func special(i int) bool {
	switch i {
	case 20, 60, 100, 140, 180, 220:
		return true
	default:
		return false
	}
}
