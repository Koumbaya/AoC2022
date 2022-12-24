package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

const (
	OPRock     = 'A'
	OPPaper    = 'B'
	OPScissors = 'C'
	RERock     = 'X'
	Loose      = Rock
	REPaper    = 'Y'
	Draw       = Paper
	REScissors = 'Z'
	Win        = Scissors
	Rock       = 1
	Paper      = 2
	Scissors   = 3
)

var winChoice = map[int]int{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var looseChoice = map[int]int{
	Paper:    Rock,
	Scissors: Paper,
	Rock:     Scissors,
}

func val(p1 uint8) int {
	switch p1 {
	case OPRock, RERock:
		return Rock
	case OPPaper, REPaper:
		return Paper
	case OPScissors, REScissors:
		return Scissors
	default:
		panic("invalid")
	}
}

//go:embed input.txt
var input string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
func part1(in string) int {
	score := 0
	for _, line := range strings.Split(in, "\n") {
		op, me := val(line[0]), val(line[2])
		if op == me {
			score += 3
		} else if me == winChoice[op] {
			score += 6
		}
		score += me
	}
	return score
}

func part2(in string) int {
	score := 0
	for _, line := range strings.Split(in, "\n") {
		op, me := val(line[0]), val(line[2])
		switch me {
		case Draw:
			score += 3
			score += op
		case Win:
			score += winChoice[op]
			score += 6
		case Loose:
			score += looseChoice[op]
		}
	}
	return score
}
