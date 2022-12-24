package day7

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
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

type Node struct {
	isDir  bool
	size   int64
	parent *Node
	// we store the name of the files/directory in the map for quick access on `cd` commands, since name are unique in a given folder
	childNames map[string]*Node
}

func NewNode(name string, size int64, parent *Node) {
	n := &Node{
		isDir:      size == 0,
		size:       size,
		parent:     parent,
		childNames: make(map[string]*Node),
	}
	if size != 0 {
		AddSizeUpstream(parent, size)
	}
	parent.AddChild(n, name)
}

func (n *Node) AddChild(ch *Node, name string) {
	n.childNames[name] = ch
}

// AddSizeUpstream add the size to all parents folder recursively.
func AddSizeUpstream(ch *Node, size int64) {
	ch.size += size
	if ch.parent == nil {
		return // main folder
	}
	AddSizeUpstream(ch.parent, size)
}

func part1(in string) int64 {
	mainDir := &Node{
		isDir:      true,
		parent:     nil, // should be the only parentless dir
		childNames: make(map[string]*Node),
	}
	currDir := mainDir
	lines := strings.Split(in, "\n")
	for _, line := range lines[1:] { // skip first line
		spl := strings.Split(line, " ")
		switch spl[0] {
		case "$":
			if spl[1] == "ls" {
				continue // we don't care
			}
			if spl[1] == "cd" && spl[2] == ".." {
				currDir = currDir.parent // go up one level
				continue
			}
			// else current directory is one of the subdirectory of currDir
			currDir = currDir.childNames[spl[2]]
		case "dir":
			// we discover a new subdirectory, add it to the current directory
			NewNode(spl[1], 0, currDir)
		default:
			// must be a file
			size, _ := strconv.ParseInt(spl[0], 10, 64)
			NewNode(spl[1], size, currDir)
		}
	}
	//Print(mainDir, "")
	return calcSize(mainDir)
}

func Print(n *Node, space string) {
	fmt.Printf("%s (%d)", space, n.size)
	fmt.Println()
	space += "  "
	for _, node := range n.childNames {
		Print(node, space)
	}
}

func calcSize(n *Node) int64 {
	sum := int64(0)
	if n.size <= 100000 {
		sum += n.size
	}
	for _, node := range n.childNames {
		if node.isDir {
			sum += calcSize(node)
		}
	}
	return sum
}

func part2(in string) int64 {
	const systemSize = 70000000
	const neededSize = 30000000

	mainDir := &Node{
		isDir:      true,
		parent:     nil, // should be the only parentless dir
		childNames: make(map[string]*Node),
	}
	currDir := mainDir

	lines := strings.Split(in, "\n")
	for _, line := range lines[1:] { // skip first line
		spl := strings.Split(line, " ")
		switch spl[0] {
		case "$":
			if spl[1] == "ls" {
				continue // we don't care
			}
			if spl[1] == "cd" && spl[2] == ".." {
				currDir = currDir.parent // go up one level
				continue
			}
			// else current directory is one of the child
			currDir = currDir.childNames[spl[2]]
		case "dir":
			// we discover a new subdirectory, add it to the parent
			NewNode(spl[1], 0, currDir)
		default:
			// must be a file
			size, _ := strconv.ParseInt(spl[0], 10, 64)
			NewNode(spl[1], size, currDir)
		}
	}

	toFreeUp := neededSize - (systemSize - mainDir.size)
	return findSmallest(mainDir, toFreeUp)
}

// findSmallest traverses the directories recursively to find the smallest dir
// that if erased would free at least toFree.
func findSmallest(n *Node, toFree int64) int64 {
	min := int64(math.MaxInt64)
	if n.size >= toFree {
		min = n.size
	}
	for _, node := range n.childNames {
		if node.isDir {
			childSize := findSmallest(node, toFree)
			if childSize < min {
				min = childSize
			}
		}
	}
	return min
}
