package day18

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
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type pos struct {
	x int
	y int
	z int
}

func part1(in string) int {
	cubes := make(map[pos]struct{})
	for _, line := range strings.Split(in, "\n") {
		spl := util.GetInts(line)
		cubes[pos{
			x: spl[0],
			y: spl[1],
			z: spl[2],
		}] = struct{}{}
	}

	faces := 0
	for p := range cubes {
		for _, n := range getNeighbors(p) {
			_, exist := cubes[n]
			if !exist {
				faces++
			}
		}
	}
	return faces
}

var relativeNeighbors = [...]pos{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func (p pos) add(p2 pos) pos {
	return pos{
		x: p.x + p2.x,
		y: p.y + p2.y,
		z: p.z + p2.z,
	}
}

func getNeighbors(p pos) (n []pos) {
	for _, neighbor := range relativeNeighbors {
		n = append(n, p.add(neighbor))
	}
	return n
}

func getNeighborsInBounds(p pos, max int) (n []pos) {
	for _, neighbor := range relativeNeighbors {
		pn := p.add(neighbor)
		if pn.x <= max && pn.y <= max && pn.z <= max &&
			pn.x >= -2 && pn.y >= -2 && pn.z >= -2 {
			n = append(n, p.add(neighbor))
		}
	}
	return n
}

func part2(in string) int {
	cubes := make(map[pos]struct{})
	maxSize := 0
	for _, line := range strings.Split(in, "\n") {
		spl := util.GetInts(line)
		p := pos{
			x: spl[0],
			y: spl[1],
			z: spl[2],
		}
		cubes[p] = struct{}{}
		maxSize = max(maxSize, p)
	}
	maxSize += 2

	faces := 0
	toExplore := []pos{{}}                 // insert pos 0,0,0
	visited := map[pos]struct{}{pos{}: {}} // insert pos 0,0,0
	for len(toExplore) > 0 {
		current := toExplore[0]
		toExplore = toExplore[1:]
		for _, n := range getNeighborsInBounds(current, maxSize) {
			if _, exist := cubes[n]; exist {
				faces++
			} else if _, seen := visited[n]; !seen {
				visited[n] = struct{}{}
				toExplore = append(toExplore, n)
			}
		}
	}
	return faces
}

func max(max int, p pos) int {
	if p.x > max {
		max = p.x
	}
	if p.y > max {
		max = p.y
	}
	if p.z > max {
		max = p.z
	}
	return max
}
