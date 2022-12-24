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
	fmt.Println(part2(input))
}

func part1(in string) int {
	x := 1 // register
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
			strengths += cycle * x
		}
		if (cycle-1)%40 == 0 {
			fmt.Println()
		}
		litDark(x, cycle)
		x += val
	}

	return strengths
}

func litDark(x, cycle int) {
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
	if i == 20 ||
		i == 60 ||
		i == 100 ||
		i == 140 ||
		i == 180 ||
		i == 220 {
		return true
	}
	return false
}

func part2(in string) int {

	return 0
}
