package day11

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/util"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func Run() {
	fmt.Println(part1(input, 20))
}

type monkey struct {
	items     []int
	operation func(old int) int
	test      int
	mkTrue    int
	mkFalse   int
}

func parseMonkey(lines []string) (mk monkey) {
	mk.items = getItems(lines[1][strings.Index(lines[1], ": ")+2:])
	mk.operation = getOperation(lines[2][strings.Index(lines[2], ": ")+2:])
	mk.test = util.GetFirstInt(lines[3])
	mk.mkTrue = util.GetFirstInt(lines[4])
	mk.mkFalse = util.GetFirstInt(lines[5])
	return mk
}

func getItems(s string) []int {
	s = strings.ReplaceAll(s, ",", "")
	s2 := strings.Split(s, " ")
	var it []int
	for _, s3 := range s2 {
		it = append(it, util.MustAtoi(s3))
	}
	return it
}

func getOperation(s string) func(old int) int {
	s2 := strings.Split(s, " ")
	switch s2[3] {
	case "*":
		if s2[4] == "old" {
			return func(old int) int {
				return old * old
			}
		}
		val := util.MustAtoi(s2[4])
		return func(old int) int {
			return old * val
		}
	case "+":
		val := util.MustAtoi(s2[4])
		return func(old int) int {
			return old + val
		}
	default:
		panic("no operation for this")
	}
}

func part1(in string, rounds int) int {
	var monkeys []monkey
	lines := strings.Split(in, "\n")
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, parseMonkey(lines[i:i+7]))
	}
	inspected := make(map[int]int, len(monkeys))

	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			// for each item carried pop it and treat it
			length := len(monkeys[m].items)
			for it := 0; it < length; it++ {
				var item int
				item, monkeys[m].items = monkeys[m].items[0], monkeys[m].items[1:]
				// inspect
				inspected[m]++
				item = monkeys[m].operation(item)
				// worry level decrease
				item = item / 3
				// give item to another monkey depending on test
				if item%monkeys[m].test == 0 {
					monkeys[monkeys[m].mkTrue].items = append(monkeys[monkeys[m].mkTrue].items, item)
				} else {
					monkeys[monkeys[m].mkFalse].items = append(monkeys[monkeys[m].mkFalse].items, item)
				}
			}
		}
	}

	return topTwo(inspected)
}

func topTwo(m map[int]int) int {
	var t0, t1 int
	for _, v := range m {
		if v > t0 {
			t1 = t0
			t0 = v
		} else if v > t1 {
			t1 = v
		}
	}
	return t0 * t1
}
