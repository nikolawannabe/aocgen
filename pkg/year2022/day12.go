package year2022

import (
	"log"
	"math"
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
}

type Node struct {
	self     position
	distance int
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

func (f *forestMap) searchTree() (int, bool) {
	n := Node{self: f.start, distance: 0}
	return f.BFS(n)
}

func (f *forestMap) searchTreeFrom(pos position) (int, bool) {
	n := Node{self: pos, distance: 0}
	return f.BFS(n)
}

func (f *forestMap) BFS(node Node) (int, bool) {
	queue := []Node{}
	queue = append(queue, node)
	return f.processQueue(queue)
}

func (f *forestMap) processQueue(queue []Node) (int, bool) {
	if len(queue) == 0 {
		return 0, false
	}

	for len(queue) > 0 {
		if queue[0].self.x == f.end.x && queue[0].self.y == f.end.y {
			log.Printf("found goal: %d, %t", queue[0].distance, true)
			return queue[0].distance, true
		}
		x := queue[0].self.x
		y := queue[0].self.y
		options := []position{{x: x + 1, y: y}, {x: x, y: y + 1}, {x: x - 1, y: y}, {x, y - 1}}
		for _, option := range options {
			_, ok := f.heightGridInt[option]
			if ok &&
				(f.heightGridInt[option]-f.heightGridInt[queue[0].self] < 2 ||
					f.heightGridInt[queue[0].self] > f.heightGridInt[option]) {
				if _, pres := f.visited[option]; !pres {
					if _, pres := f.visited[option]; !pres {
						f.visited[option] = true
						n := Node{self: option, distance: queue[0].distance + 1}
						//log.Printf("added %#v", n)
						queue = append(queue, n)
					}
				} else {
					//log.Printf("already visited %#v", option)
				}
			}
		}
		queue = queue[1:]
	}
	return 0, false

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
			heightMap[p] = height
		}
	}
	f.heightGridInt = heightMap
}

func setup(lines []string) forestMap {
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
				visited[p] = true
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
	forest.convertToInt()
	return forest
}
func (p Day12) PartA(lines []string) any {
	forest := setup(lines)
	printGrid(forest)

	steps, isFound := forest.searchTree()
	if !isFound {
		log.Printf("not found")
		return 0
	}

	return steps
}

func (p Day12) PartB(lines []string) any {
	forest := setup(lines)
	printGrid(forest)

	starts := make([]position, 0)
	for pos, height := range forest.heightGridInt {
		if height == 0 {
			starts = append(starts, pos)
		}
	}
	log.Printf("found %d starts", len(starts))

	pathLengths := make(map[position]int, 0)
	for _, start := range starts {
		forest.visited = make(map[position]bool, 0)
		log.Printf("testing start: %#v", start)
		steps, isFound := forest.searchTreeFrom(start)
		if isFound {
			log.Printf("found")
			pathLengths[start] = steps
		} else {
			log.Printf("not found")
		}
	}

	minSteps := math.MaxInt
	for _, steps := range pathLengths {
		if steps < minSteps {
			minSteps = steps
		}
	}
	return minSteps
}
