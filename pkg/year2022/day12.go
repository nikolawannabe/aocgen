package year2022

import (
	"log"
	"strconv"
)

type Day12 struct{}

type forestMap struct {
	heightGrid    map[position]rune
	heightGridInt map[position]int
	visited       map[position]bool
	width         int
	height        int
	current       position
	start         position
	end           position
	tree          *Node
	levelCount    int
}

type Node struct {
	selfHeight int
	self       position
	neighbors  []Node
}

func printMap(forest forestMap) {
	output := ""

	for y := 0; y < forest.height; y++ {
		for x := 0; x < forest.width; x++ {
			output += string(forest.heightGrid[position{x: x, y: y}])
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n")
}

func (n *Node) addNeighbors(forest forestMap) {
	ns := make([]Node, 0)
	x := n.self.x
	y := n.self.y
	options := []position{{x: x + 1, y: y}, {x: x, y: y + 1}, {x: x - 1, y: y}, {x, y - 1}}
	for _, option := range options {
		_, ok := forest.heightGridInt[option]
		if ok &&
			(forest.heightGridInt[option]-n.selfHeight < 2 ||
				n.selfHeight > forest.heightGridInt[option]) {
			if !forest.visited[option] {
				neighbor := Node{selfHeight: forest.heightGridInt[option], self: option}
				forest.visited[option] = true
				neighbor.addNeighbors(forest)
				ns = append(ns, neighbor)
				continue
			}

		}
	}
	n.neighbors = ns
}

func (f *forestMap) buildTree() {
	start := f.start
	nodeStart := Node{selfHeight: f.heightGridInt[start], self: start}
	nodeStart.addNeighbors(*f)
	f.tree = &nodeStart
}

func (f *forestMap) visitNode(n Node) bool {
	f.levelCount++
	//log.Printf("entered depth: %d", f.levelCount)
	for _, child := range n.neighbors {
		if child.self.x == f.end.x && child.self.y == f.end.y {
			log.Printf("%v", child.self)
			return true
		}
	}

	for _, child := range n.neighbors {
		found := f.visitNode(child)
		if found {
			return true
		}

	}
	f.levelCount--
	//log.Printf("left, at depth: %d", f.levelCount)
	return false
}

func (f *forestMap) searchTree() bool {
	return f.visitNode(*f.tree)
}

func printGrid(forest forestMap) {
	output := ""

	for y := 0; y < forest.height; y++ {
		for x := 0; x < forest.width; x++ {
			p := position{x: x, y: y}
			intStr := strconv.Itoa(forest.heightGridInt[p])
			if forest.heightGridInt[p] < 10 {
				output += "  " + intStr
			} else {
				output += " " + intStr
			}
		}
		log.Printf("%s", output)
		output = ""
		log.Printf("\n")
	}
	log.Printf("\n")
}

func (f *forestMap) convertToInt() {
	heightMap := make(map[position]int, 0)

	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			p := position{x: x, y: y}
			height := int(f.heightGrid[p]) - 'a'
			log.Printf("p: %d, height: %d", p, height)
			heightMap[p] = height
		}
	}
	f.heightGridInt = heightMap
}

func (p Day12) PartA(lines []string) any {
	forest := forestMap{}
	rMap := make(map[position]rune, 0)
	visited := make(map[position]bool, 0)

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		for x, char := range line {
			p := position{x: x, y: y}
			if char == 'S' {
				forest.start = p
				forest.current = p
				rMap[p] = 'a'
				continue
			}
			if char == 'E' {
				forest.end = p
				rMap[p] = 'z'
				continue
			}
			rMap[p] = char
		}
	}
	forest.heightGrid = rMap
	//forest.levelCount = 1
	forest.width = len(lines[0])
	forest.height = len(lines)
	forest.visited = visited
	printMap(forest)
	forest.convertToInt()
	printGrid(forest)
	forest.buildTree()
	isFound := forest.searchTree()
	if !isFound {
		log.Printf("did not find node")
	}
	return forest.levelCount
}

func (p Day12) PartB(lines []string) any {
	return "implement_me"
}
