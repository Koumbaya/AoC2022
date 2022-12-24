package day1

import (
	_ "embed"
	"fmt"
	"sort"
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
	maxCal := 0
	currCount := 0
	for _, line := range strings.Split(in, "\n") {
		if len(line) == 0 {
			if currCount > maxCal {
				maxCal = currCount
			}
			currCount = 0
			continue
		}
		val, _ := strconv.Atoi(line)
		currCount += val
	}
	return maxCal
}

func part2(in string) int {
	elves := make([]int, 0)
	currCount := 0
	for _, line := range strings.Split(in, "\n") {
		if len(line) == 0 {
			elves = append(elves, currCount)
			currCount = 0
			continue
		}
		val, _ := strconv.Atoi(line)
		currCount += val
	}
	sort.Ints(elves)

	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
}
