package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
func part1(in string) int {
	score := 0
	for _, line := range strings.Split(in, "\n") {
		spl := strings.Split(line, ",")
		range1 := strings.Split(spl[0], "-")
		range2 := strings.Split(spl[1], "-")
		st1, _ := strconv.Atoi(range1[0])
		ed1, _ := strconv.Atoi(range1[1])
		st2, _ := strconv.Atoi(range2[0])
		ed2, _ := strconv.Atoi(range2[1])
		if (st1 >= st2 && ed1 <= ed2) ||
			(st2 >= st1 && ed2 <= ed1) {
			score++
		}
	}
	return score
}

func part2(in string) int {
	score := 0
	for _, line := range strings.Split(in, "\n") {
		spl := strings.Split(line, ",")
		range1 := strings.Split(spl[0], "-")
		range2 := strings.Split(spl[1], "-")
		st1, _ := strconv.Atoi(range1[0])
		ed1, _ := strconv.Atoi(range1[1])
		st2, _ := strconv.Atoi(range2[0])
		ed2, _ := strconv.Atoi(range2[1])
		if (st1 < st2 && ed1 < ed2 && ed1 < st2) ||
			(st1 > st2 && ed1 > ed2 && st1 > ed2) {
			continue
		}
		score++
	}
	return score
}
