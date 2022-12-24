package day3

import (
	_ "embed"
	"fmt"
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
		bag1, bag2 := line[:len(line)/2], line[len(line)/2:]
		c := common(bag1, bag2)
		score += priority(c)
	}
	return score
}

func common(b1, b2 string) uint8 {
	for i := range b1 {
		for j := range b2 {
			if b1[i] == b2[j] {
				return b1[i]
			}
		}
	}
	panic("no common element")
}

// priority we remove X from the ascii value to get the right priority.
func priority(c uint8) int {
	if c <= 90 { // uppercase. A-Z = 27-52.
		return int(c - 38)
	}
	return int(c - 96) // lowercase. a-z = 1-26.
}

func part2(in string) int {
	lines := strings.Split(in, "\n")
	score := 0
	for i := 0; i < len(lines); i += 3 {
		h1 := make(map[uint8]struct{}, len(lines[i]))
		h2 := make(map[uint8]struct{}, 0)
		for a := range lines[i] {
			h1[lines[i][a]] = struct{}{}
		}
		for a := range lines[i+1] {
			if _, ok := h1[lines[i+1][a]]; ok {
				h2[lines[i+1][a]] = struct{}{}
			}
		}
		for a := range lines[i+2] {
			if _, ok := h2[lines[i+2][a]]; ok {
				score += priority(lines[i+2][a])
				break // only one element in common
			}
		}
	}
	return score
}
