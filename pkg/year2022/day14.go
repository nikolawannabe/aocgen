package year2022

import (
	"log"
	"strings"
)

type Day14 struct{}

const (
	rock     = '#'
	sediment = 'o'
	sand     = 's'
)

type cave struct {
	width       int
	height      int
	obstacles   map[position]icon
	sand        *position
	sandCounter int
}
type icon rune

func (c *cave) printWindow(windowX int, windowY int, windowLen int) {
	output := ""

	for y := windowY; y < c.height && y < windowY+windowLen; y++ {
		for x := windowX - 1; x < c.width && x < windowX+windowLen; x++ {
			if x == windowX-1 {
				yStr := itoa(y)
				output += yStr[len(yStr)-1:]
				continue
			}
			p := position{x: x, y: y}
			if i, pres := c.obstacles[p]; pres {
				output += string(i)
			} else if c.sand != nil && *c.sand == p {
				output += string(sand)
			} else {
				output += "."
			}
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n")
}

func (c *cave) tickSand() bool {
	if c.sand.x+1 > c.width || c.sand.y+1 > c.width {
		log.Printf("exceeded bounds: %d, %d", c.sand.x, c.sand.y)
	}
	if _, pres := c.obstacles[position{x: c.sand.x, y: c.sand.y + 1}]; pres {
		if _, pres := c.obstacles[position{x: c.sand.x - 1, y: c.sand.y + 1}]; pres {
			if _, pres := c.obstacles[position{x: c.sand.x + 1, y: c.sand.y + 1}]; pres {
				c.obstacles[*c.sand] = sediment
				if _, pres := c.obstacles[position{500, 0}]; pres {
					return true
				}
				c.sand = &position{500, 0}
				c.sandCounter++
			} else {
				c.sand = &position{x: c.sand.x + 1, y: c.sand.y + 1}
			}
		} else {
			c.sand = &position{x: c.sand.x - 1, y: c.sand.y + 1}
		}
	} else {
		c.sand = &position{x: c.sand.x, y: c.sand.y + 1}
	}
	return false
}

func (c *cave) renderObstacle(coordinateList []position) {
	head := coordinateList[0]
	for i := 1; i < len(coordinateList); i++ {
		tail := coordinateList[i]
		if head.x != tail.x {
			if head.x < tail.x {
				for x := head.x; x <= tail.x; x++ {
					c.obstacles[position{x: x, y: head.y}] = rock
				}
			} else {
				for x := head.x; x >= tail.x; x-- {
					c.obstacles[position{x: x, y: head.y}] = rock
				}
			}
		} else if head.y != tail.y {
			if head.y < tail.y {
				for y := head.y; y <= tail.y; y++ {
					c.obstacles[position{x: head.x, y: y}] = rock
				}
			} else {
				for y := head.y; y >= tail.y; y-- {
					c.obstacles[position{x: head.x, y: y}] = rock
				}
			}
		}
		head = tail
	}
}

func (c *cave) renderFloor(floorY int) {
	for x := 0; x < c.width; x++ {
		c.obstacles[position{x: x, y: floorY}] = rock
	}
}

func (p Day14) PartA(lines []string) any {
	floorY := 0
	obstacles := make(map[position]icon, 0)
	c := cave{obstacles: obstacles, sand: nil, width: 900, height: 170}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		coordinateStrings := strings.Split(line, " -> ")
		coordinateList := make([]position, 0)
		for _, coordinate := range coordinateStrings {
			pair := strings.Split(coordinate, ",")
			y := atoi(pair[1])
			if y > floorY {
				floorY = y
			}
			p := position{x: atoi(pair[0]), y: y}
			coordinateList = append(coordinateList, p)
		}
		c.renderObstacle(coordinateList)
	}
	c.printWindow(494, 0, 10)

	pos := position{500, 0}
	c.sand = &pos
	c.printWindow(494, 0, 10)
	iter := 0
	for c.sand.y < floorY {
		c.tickSand()
		iter++
		if iter%10000 == 0 {
			log.Printf("===== iter %d ====", iter)
			c.printWindow(300, 0, 300)
		}
	}
	c.printWindow(494, 0, 10)
	return c.sandCounter
}

func (p Day14) PartB(lines []string) any {
	lastY := 0
	lastX := 0
	obstacles := make(map[position]icon, 0)
	c := cave{obstacles: obstacles, sand: nil, width: 900, height: 170}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		coordinateStrings := strings.Split(line, " -> ")
		coordinateList := make([]position, 0)
		for _, coordinate := range coordinateStrings {
			pair := strings.Split(coordinate, ",")
			y := atoi(pair[1])
			if y > lastY {
				lastY = y
			}
			x := atoi(pair[0])
			if x > lastX {
				lastX = x
			}
			p := position{x: x, y: y}
			coordinateList = append(coordinateList, p)
		}
		c.renderObstacle(coordinateList)
	}
	floorY := lastY + 2
	c.height = floorY + 2
	c.width = lastX * 3
	c.renderFloor(floorY)
	log.Printf("moo")

	pos := position{500, 0}
	c.sand = &pos
	c.printWindow(0, 0, c.width)
	iter := 0
	found := false
	c.sandCounter = 1
	for !found {
		found = c.tickSand()
		iter++
		if iter%10000 == 0 {
			log.Printf("===== iter %d ====", iter)
			c.printWindow(0, 0, c.width)
		}
	}
	c.printWindow(0, 0, c.width)
	return c.sandCounter
}
