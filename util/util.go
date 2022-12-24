package util

import (
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	s = strings.Trim(s, "\n")
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func GetFirstInt(s string) int {
	s = strings.Trim(s, "\n")
	spl := strings.Split(s, " ")
	for _, s2 := range spl {
		val, err := strconv.Atoi(s2)
		if err == nil {
			return val
		}
	}
	panic("nostring")
}

func GetInts(s string) (ret []int) {
	spl := strings.Split(s, ",")
	for _, s2 := range spl {
		val, _ := strconv.Atoi(s2)
		ret = append(ret, val)
	}
	return ret
}

func Abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
