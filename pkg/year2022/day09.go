package year2022

import (
	"fmt"
	"log"
	"strconv"
)

type Day09 struct{}

type position struct {
	x int
	y int
}

func move(h position, direction []rune) position {
	for _, dirPart := range direction {
		switch dirPart {
		case 'R':
			h.x++
		case 'L':
			h.x--
		case 'U':
			h.y++
		case 'D':
			h.y--
		}
	}
	return h
}

func straightDirection(t position, h position) rune {
	if t.x == h.x {
		if h.y > t.y {
			return 'U'
		}
		if h.y < t.y {
			return 'D'
		}
	}
	if t.y == h.y {
		if h.x > t.x {
			return 'R'
		}
		if h.x < t.x {
			return 'L'
		}
	}
	return 0
}

func moveDiagonal(t position, h position) (position, []rune) {
	dir := make([]rune, 0)
	if h.x > t.x {
		t.x++
		dir = append(dir, 'R')
	} else if h.x < t.x {
		t.x--
		dir = append(dir, 'L')
	}

	if h.y > t.y {
		t.y++
		dir = append(dir, 'U')
	} else if h.y < t.y {
		t.y--
		dir = append(dir, 'D')
	}
	return t, dir
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lookDiagonal(t position, h position) bool {
	xd := h.x - t.x
	yd := h.y - t.y
	if abs(xd)+abs(yd) > 2 {
		return true
	}
	return false
}

func lookStraight(t position, h position) bool {
	if t.x == h.x {
		yd := h.y - t.y
		if abs(yd) > 1 {
			return true
		}
	}
	if t.y == h.y {
		xd := h.x - t.x
		if abs(xd) > 1 {
			return true
		}
	}
	return false
}

func moveTail(t position, h position) (position, []rune) {
	if t.x == h.x && t.y == h.y {
		return t, []rune{0}
	}
	if lookStraight(t, h) {
		dir := straightDirection(t, h)
		return move(t, []rune{dir}), []rune{dir}
	}
	if lookDiagonal(t, h) {
		return moveDiagonal(t, h)
	}
	return t, []rune{0}
}

func (p Day09) PartA(lines []string) any {
	v := make(map[position]bool, 0)
	h := position{x: 0, y: 0}
	t := position{x: 0, y: 0}
	v[t] = true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var direction rune
		var amount int
		fmt.Sscanf(line, "%c %d", &direction, &amount)
		log.Printf("== %c %d ==", direction, amount)
		for step := 0; step < amount; step++ {
			h = move(h, []rune{direction})
			log.Printf("h: %v", h)
			t, _ = moveTail(t, h)
			log.Printf("t: %v\n\n", t)
			v[t] = true

		}
		printStep(v, 30, 30)
	}

	l := 0
	for coord, visited := range v {
		if visited {
			log.Printf("visited %d, %d", coord.x, coord.y)
			l++
		}
	}
	return l
}

func moveHeadPair(h position, t position, hDir []rune) (position, position, []rune) {
	h = move(h, hDir)
	t, dir := moveTail(t, h)
	return h, t, dir
}

func moveTrio(prev position, h position, t position) (position, position, []rune) {
	h, _ = moveTail(h, prev)
	t, dir := moveTail(t, h)
	return h, t, dir
}

func printStep(visited map[position]bool, xb int, yb int) {
	output := ""

	for y := 15; y > -15; y-- {
		for x := -15; x < 15; x++ {
			if visited[position{x: x, y: y}] == true {
				output += "#"
			} else {
				output += "."
			}
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n\n")
}

func printChar(char map[position]int, xb int, yb int) {
	output := ""

	for y := 15; y > -15; y-- {
		for x := -15; x < 15; x++ {
			if char[position{x: x, y: y}] > 0 {
				output += strconv.Itoa(char[position{x: x, y: y}])
			} else {
				output += "."
			}
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n\n")
}

type character struct {
	num int
	pos position
}

func (p Day09) PartB(lines []string) any {
	v := make(map[position]bool, 0)
	c := make(map[position]int, 0)
	rope := make([]position, 0)

	for i := 0; i < 10; i++ {
		rope = append(rope, position{})
	}
	v[position{0, 0}] = true

	for _, line := range lines {
		c = make(map[position]int, 0)
		if len(line) == 0 {
			continue
		}
		var direction rune
		var amount int
		fmt.Sscanf(line, "%c %d", &direction, &amount)
		dir := []rune{direction}
		//log.Printf("== %s %d ==", concat(dir), amount)
		for step := 0; step < amount; step++ {
			//log.Printf("%s", concat(dir))
			h, t, _ := moveHeadPair(rope[0], rope[1], dir)
			rope[0] = h
			rope[1] = t
			for i := 1; i < len(rope)-1; i++ {
				// log.Printf("= %s =", concat(dir))
				h, t, _ := moveTrio(rope[i-1], rope[i], rope[i+1])
				rope[i] = h
				rope[i+1] = t
			}
			v[rope[9]] = true
		}
		for i, pos := range rope {
			c[pos] = i
		}
		printStep(v, 30, 30)
		printChar(c, 30, 30)
	}

	l := 0
	for _, visited := range v {
		if visited {
			l++
		}
	}
	return l
}
