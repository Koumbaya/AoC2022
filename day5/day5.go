package day5

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Stack []uint8

func (s *Stack) Push(r uint8) {
	*s = append(*s, r)
}

func (s *Stack) PushN(r []uint8) {
	*s = append(*s, r...)
}

func (s *Stack) Pop() uint8 {
	l := len(*s)
	if l == 0 {
		return 0
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

func (s *Stack) PopN(n int) []uint8 {
	l := len(*s)
	if l == 0 {
		return []uint8{}
	}

	res := (*s)[l-n:]
	*s = (*s)[:l-n]

	return res
}

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseInitState(lines []string) []Stack {
	stacks := make([]Stack, 0, 8)
	for i := 0; i < 9; i++ {
		stacks = append(stacks, Stack{})
	}
	// we parse the lines (levels) in reverse so that we can use FILO on the stacks.
	for levelIdx := 7; levelIdx >= 0; levelIdx-- {
		level := lines[levelIdx][1:] // ignore first [
		// columns are 4 chars apart, so we check letterIdx every 4 space
		for letterIdx, colIdx := 0, 0; letterIdx < len(level); letterIdx, colIdx = letterIdx+4, colIdx+1 {
			if level[letterIdx] != ' ' { // we don't push empty, it means we're at the top of this column
				stacks[colIdx].Push(level[letterIdx])
			}
		}
	}

	return stacks
}

func part1(in string) string {
	total := strings.Split(in, "\n")
	setup := total[:8]
	stacks := parseInitState(setup)

	for _, l := range total[10:] {
		l = strings.Replace(l, "move ", "", 1)
		l = strings.Replace(l, " from ", " ", 1)
		l = strings.Replace(l, " to ", " ", 1)
		ls := strings.Split(l, " ")
		nb, _ := strconv.Atoi(ls[0])
		from, _ := strconv.Atoi(ls[1])
		from--
		to, _ := strconv.Atoi(ls[2])
		to--

		var r uint8
		for i := 0; i < nb; i++ {
			r = stacks[from].Pop()
			stacks[to].Push(r)
		}
	}
	code := ""
	for i := range stacks {
		code += string(stacks[i].Pop())
	}
	return code
}

func part2(in string) string {
	total := strings.Split(in, "\n")
	setup := total[:8]
	stacks := parseInitState(setup)

	for _, l := range total[10:] {
		l = strings.Replace(l, "move ", "", 1)
		l = strings.Replace(l, " from ", " ", 1)
		l = strings.Replace(l, " to ", " ", 1)
		ls := strings.Split(l, " ")
		nb, _ := strconv.Atoi(ls[0])
		from, _ := strconv.Atoi(ls[1])
		from--
		to, _ := strconv.Atoi(ls[2])
		to--

		r := stacks[from].PopN(nb)
		stacks[to].PushN(r)
	}
	code := ""
	for i := range stacks {
		code += string(stacks[i].Pop())
	}
	return code
}
