package day6

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(in string) int {
	for i := 0; i < len(in); i++ {
		if in[i] != in[i+1] &&
			in[i] != in[i+2] &&
			in[i] != in[i+3] &&
			in[i+1] != in[i+2] &&
			in[i+1] != in[i+3] &&
			in[i+2] != in[i+3] {
			return i + 4 // 3 for last char +1 because AOC starts from 1
		}
	}
	return 0
}

func part2(in string) int {
	for i := 0; i < len(in); i++ {
		m := make(map[uint8]struct{})
		for j := 0; j < 14; j++ {
			_, exist := m[in[i+j]]
			if exist {
				break
			}
			m[in[i+j]] = struct{}{}
		}
		if len(m) == 14 {
			return i + 14
		}
	}
	return 0
}
