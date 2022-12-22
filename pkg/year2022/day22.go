package year2022

import "log"

type Day22 struct{}

type puzzleMap struct {
	boardGrid   map[position]rune
	visited     map[position]rune
	start       position
	width       int
	height      int
	current     position
	orientation rune
}

var (
	orientationMap      map[rune]int
	rotationMap         map[rune]int
	orientationScoreMap map[int]rune
)

func init() {
	// facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	orientationMap = make(map[rune]int, 0)
	orientationScoreMap = make(map[int]rune, 0)
	rotationMap = make(map[rune]int, 0)
	orientationMap['>'] = 0
	orientationScoreMap[0] = '>'
	orientationMap['v'] = 1
	orientationScoreMap[1] = 'v'
	orientationMap['<'] = 2
	orientationScoreMap[2] = '<'
	orientationMap['^'] = 3
	orientationScoreMap[3] = '^'
	rotationMap['R'] = 1
	rotationMap['L'] = -1

}
func (board *puzzleMap) print() {
	output := ""

	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			p := position{x: x, y: y}
			item, pres := board.boardGrid[p]
			if !pres {
				item = ' '
			}
			if pres {
				visited, hasVisited := board.visited[p]
				if hasVisited {
					item = visited
				}
			}
			output += string(item)
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n")
}

func (b *puzzleMap) SetStart() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			item, pres := b.boardGrid[position{x: x, y: y}]
			if !pres || item == '#' {
				continue
			}
			p := position{x: x, y: y}
			b.start = p
			b.current = p
			b.orientation = '>'
			b.visited[p] = b.orientation
			return
		}
	}
}
func (b *puzzleMap) Turn(direction rune) {
	o := orientationMap[b.orientation]
	o += rotationMap[direction]
	b.orientation = orientationScoreMap[o]
}

func (b *puzzleMap) moveDirection(amount int, compass rune) {
	for i := 0; i < amount; i++ {
		newP := getNext(b.current, compass)
		item, pres := b.boardGrid[newP]
		if pres && item == '.' {
			b.current = newP
			b.visited[newP] = b.orientation
			continue
		}
		if !pres {
			log.Printf("need to wrap!")
			continue
		}
		if item == '#' {
			break
		}
	}
}
func getNext(p position, compass rune) position {
	switch compass {
	case 'N':
		return position{x: p.x, y: p.y - 1}
	case 'S':
		return position{x: p.x, y: p.y + 1}
	case 'W':
		return position{x: p.x - 1, y: p.y}
	case 'E':
		return position{x: p.x + 1, y: p.y}
	}
	return p
}

func (b *puzzleMap) Move(amount int) {
	switch b.orientation {
	case '>':
		b.moveDirection(amount, 'E')
	case 'v':
		b.moveDirection(amount, 'S')
	case '<':
		b.moveDirection(amount, 'W')
	case '^':
		b.moveDirection(amount, 'N')
	}
}

func parseBoard(lines []string) puzzleMap {
	board := puzzleMap{}
	boardMap := make(map[position]rune, 0)
	visited := make(map[position]rune, 0)

	board.width = 0
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		if len(line) > board.width {
			board.width = len(line)
		}
		for x, char := range line {
			p := position{x: x, y: y}
			if char == '#' {
				boardMap[p] = '#'
				continue
			}
			if char == '.' {
				boardMap[p] = '.'
				continue
			}
		}
	}
	board.boardGrid = boardMap
	board.height = len(lines)
	board.visited = visited
	return board
}

func (p Day22) PartA(lines []string) any {
	board := parseBoard(lines[0 : len(lines)-2])
	board.print()
	board.SetStart()
	board.Move(5)
	board.Turn('R')
	board.Move(10)
	board.Turn('R')
	board.Move(3)
	board.Turn('R')
	board.Move('1')
	board.Turn('L')
	board.Turn('L')
	board.Move(1)
	board.Turn('R')
	board.Move(10)
	board.print()
	return 0
}

func (p Day22) PartB(lines []string) any {
	return "implement_me"
}
