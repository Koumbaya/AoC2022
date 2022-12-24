package day9

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"aoc/util"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

//go:embed test2.txt
var test2 string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type pos struct {
	x int
	y int
}

func part1(in string) int {
	visited := make(map[pos]struct{})
	visited[pos{x: 0, y: 0}] = struct{}{}

	var hx, hy, tx, ty int // head and tail positions.
	for _, line := range strings.Split(in, "\n") {
		l := strings.Split(line, " ")
		dir := l[0]
		steps, _ := strconv.Atoi(l[1])
		for i := 0; i < steps; i++ {
			hx, hy = moveHead(hx, hy, dir)
			tx, ty = moveTail(hx, hy, tx, ty)
			visited[pos{x: tx, y: ty}] = struct{}{}
			//printTest(hx, hy, tx, ty)
		}
	}
	return len(visited)
}

func part2(in string) int {
	visited := make(map[pos]struct{})
	visited[pos{x: 0, y: 0}] = struct{}{}
	segments := make([]pos, 9)

	var hx, hy int // head position
	for _, line := range strings.Split(in, "\n") {
		l := strings.Split(line, " ")
		dir := l[0]
		steps, _ := strconv.Atoi(l[1])
		for i := 0; i < steps; i++ {
			hx, hy = moveHead(hx, hy, dir)
			segments[0].moveSeg(hx, hy)          // move the first segment in relation to head
			for j := 1; j < len(segments); j++ { // move the rest of the segment in relation to the previous segment
				segments[j].moveSeg(segments[j-1].x, segments[j-1].y)
			}
			//printTest2(hx, hy, segments)
			visited[pos{x: segments[8].x, y: segments[8].y}] = struct{}{}
		}
	}
	return len(visited)
}

func printTest2(hx int, hy int, segments []pos) {
	for y := 15; y >= -15; y-- {
		for x := -15; x < 15; x++ {
			val := "."
			if hx == x && hy == y {
				val = "H"
			}
			for i := range segments {
				if segments[i].x == x && segments[i].y == y {
					val = fmt.Sprint(i + 1)
					break
				}
			}
			if x == 0 && y == 0 {
				val = "s"
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printTest(hx int, hy int, tx int, ty int) {
	for y := 6; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			if hx == x && hy == y {
				fmt.Print("H")
			} else if tx == x && ty == y {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveHead(hx, hy int, dir string) (int, int) {
	switch dir {
	case "U":
		hy++
	case "D":
		hy--
	case "R":
		hx++
	case "L":
		hx--
	}
	return hx, hy
}

func moveTail(hx, hy, tx, ty int) (int, int) {
	dX := util.Abs(hx, tx)
	dY := util.Abs(hy, ty)
	if dY > 1 {
		if ty < hy {
			ty = hy - 1
		} else if ty > hy {
			ty = hy + 1
		}
		if dX == 1 {
			tx = hx
		}
	} else if dX > 1 {
		if tx < hx {
			tx = hx - 1
		} else if tx > hx {
			tx = hx + 1
		}
		if dY == 1 {
			ty = hy
		}
	}

	return tx, ty
}

func (p *pos) moveSeg(rx, ry int) {
	dX := util.Abs(rx, p.x)
	dY := util.Abs(ry, p.y)
	if dX < 2 && dY < 2 {
		return // adjacent
	}
	if p.x == rx {
		if p.y < ry {
			p.y++
		} else {
			p.y--
		}
	} else if p.y == ry {
		if p.x < rx {
			p.x++
		} else {
			p.x--
		}
	} else if p.y < ry && p.x < rx {
		p.x++
		p.y++
	} else if p.y < ry && p.x > rx {
		p.x--
		p.y++
	} else if p.y > ry && p.x < rx {
		p.x++
		p.y--
	} else {
		p.x--
		p.y--
	}
}
