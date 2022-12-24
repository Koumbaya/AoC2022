package day8

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func Run() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type tree struct {
	height  int
	visible bool
}

func part1(in string) int {
	grid := strings.Split(in, "\n")
	// make grid
	xy := make([][]tree, len(grid[0]))
	for y := range xy {
		xy[y] = make([]tree, len(grid))
	}

	var topTree int
	nbVisible := 0
	// populate grid with heights
	for y, line := range grid {
		topTree = -1
		for x := range line {
			hgt := int(line[x] - '0') // align ascii values to uint8(0) by removing '0' from it.
			// calc X visibility
			xy[y][x] = tree{
				height:  hgt,
				visible: hgt > topTree,
			}
			if hgt > topTree {
				nbVisible++
				topTree = hgt
			}
		}
		topTree = -1
		// calc from opposite side
		for x := len(line) - 1; x >= 0; x-- {
			if xy[y][x].height > topTree {
				if !xy[y][x].visible {
					xy[y][x].visible = true
					nbVisible++
				}
				topTree = xy[y][x].height
			}
			if topTree == 9 {
				break // nothing visible further
			}
		}
	}
	// calc Y visibility
	for y := range xy {
		topTree = -1
		for x := range xy[y] {
			if xy[x][y].height > topTree {
				if !xy[x][y].visible {
					xy[x][y].visible = true
					nbVisible++
				}
				topTree = xy[x][y].height
			}
			if topTree == 9 {
				break // nothing visible further
			}
		}
		topTree = -1
		// calc from opposite side
		for x := len(xy[y]) - 1; x >= 0; x-- {
			if xy[x][y].height > topTree {
				if !xy[x][y].visible {
					xy[x][y].visible = true
					nbVisible++
				}
				topTree = xy[x][y].height
			}
			if topTree == 9 {
				break // nothing visible further
			}
		}
	}
	//printVisibility(&xy)
	return nbVisible
}

func printVisibility(forest *[][]tree) {
	for x := range *forest {
		for y := range (*forest)[x] {
			if (*forest)[x][y].visible {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
			//fmt.Printf("%t", forest[x][y].height)
		}
		fmt.Println()
	}
}

func part2(in string) int {
	grid := strings.Split(in, "\n")
	// make grid
	xy := make([][]int, len(grid[0]))
	for y := range xy {
		xy[y] = make([]int, len(grid))
	}

	// populate grid with heights
	for y, line := range grid {
		for x := range line {
			hgt := int(line[x] - '0')
			// calc X visibility
			xy[y][x] = hgt
		}
	}
	// calc scenic score
	score := 0
	for y, line := range grid {
		for x := range line {
			sc := calcScenic(x, y, xy[y][x], &xy)
			if sc > score {
				score = sc
			}
		}
	}
	return score
}

func calcScenic(x, y, height int, forest *[][]int) int {
	var dir1, dir2, dir3, dir4 int
	for i := x + 1; i < len((*forest)[x]); i++ {
		dir1++
		if (*forest)[y][i] < height {
			continue
		}
		break
	}

	for i := x - 1; i >= 0; i-- {
		dir2++
		if (*forest)[y][i] < height {
			continue
		}
		break
	}

	for i := y + 1; i < len((*forest)[y]); i++ {
		dir3++
		if (*forest)[i][x] < height {
			continue
		}
		break
	}

	for i := y - 1; i >= 0; i-- {
		dir4++
		if (*forest)[i][x] < height {
			continue
		}
		break
	}

	return dir1 * dir2 * dir3 * dir4
}
