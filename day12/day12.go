package day12

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"aoc/util"
)

//go:embed input.txt
var input string

func Run() {
	// TODO : fix A*
	fmt.Println(part1(input))
}

func part1(in string) int {
	grid := strings.Split(in, "\n")
	// make grid
	xy := make([][]uint8, len(grid))
	for y := range xy {
		xy[y] = make([]uint8, len(grid[0]))
	}

	var start, end pos
	for y, line := range grid {
		for x := range line {
			xy[y][x] = grid[y][line[x]]
			if xy[y][x] == 'S' {
				xy[y][x] = 'a'
				start = pos{y: y, x: x}
			}
			if xy[y][x] == 'E' {
				xy[y][x] = 'z'
				end = pos{y: y, x: x}
			}
		}
	}

	res := aStar(start, end, xy)

	return len(res)
}

type pos struct {
	y int
	x int
}

func in(p pos, ps []pos) bool {
	for i := range ps {
		if ps[i] == p {
			return true
		}
	}
	return false
}

func neighbors(p pos) []pos {
	return []pos{{y: p.y - 1, x: p.x}, pos{y: p.y + 1, x: p.x}, pos{y: p.y, x: p.x - 1}, pos{y: p.y, x: p.x + 1}}
}

func canGotoNeighbor(p, n pos, grid [][]uint8) bool {
	if n.x < 0 || n.y < 0 || n.x > len(grid) || n.y > len((grid)[0]) {
		return false
	}
	return grid[n.y][n.x] < grid[p.y][p.x] || grid[n.y][n.x]+1 == grid[p.y][p.x]
}

func manhattanDistance(s, e pos) int {
	return util.Abs(e.x, s.x) + util.Abs(e.y, s.y)
}

func lowestF(open []pos, fscore map[pos]int) (pos, int) {
	low := math.MaxInt64
	lowestPos := pos{}
	idx := 0
	for i := range open {
		if fscore[open[i]] < low {
			low = fscore[open[i]]
			lowestPos = open[i]
			idx = i
		}
	}
	return lowestPos, idx
}

// can do return len(map) directly ?
func reconstructPath(cameFrom map[pos]pos, current pos) []pos {
	totalPath := make([]pos, 0)
	for i := 0; i < len(cameFrom); i++ {
		current = cameFrom[current]
		totalPath = append(totalPath, current)
	}
	return totalPath
}

func aStar(start, end pos, grid [][]uint8) []pos {
	openSet := make([]pos, 0)
	openSet = append(openSet, start)
	camefrom := make(map[pos]pos)
	gScore := make(map[pos]int)
	fScore := make(map[pos]int)

	gScore[start] = 0
	fScore[start] = manhattanDistance(start, end)

	for len(openSet) > 0 {
		current, idx := lowestF(openSet, fScore)
		if current == end {
			return reconstructPath(camefrom, current)
		}

		openSet = append(openSet[:idx], openSet[idx+1:]...)

		for _, n := range neighbors(current) {
			if canGotoNeighbor(current, n, grid) {
				tentativeGscore := gScore[current] + 1 //fixed weight
				prevGscore, exist := gScore[n]
				if !exist || tentativeGscore < prevGscore {
					gScore[n] = tentativeGscore
					fScore[n] = tentativeGscore + manhattanDistance(n, end)
					camefrom[n] = current
					if !in(n, openSet) {
						openSet = append(openSet, n)
					}
				}
			}
		}
	}
	return []pos{}
}
