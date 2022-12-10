package year2022

import (
	"fmt"
	"log"
)

type Day10 struct{}

type cpu struct {
	register       int
	cycle          int
	horizontalLine string
}

func (c *cpu) doCycle() {
	screenX := c.cycle % 40

	if abs(screenX-c.register-1) < 2 {
		c.horizontalLine += "#"
	} else {
		c.horizontalLine += "."
	}
	if c.cycle%40 == 0 {
		log.Printf("%s", c.horizontalLine)
		c.horizontalLine = ""
	}
	c.cycle++
}

func (c *cpu) getStrength() int {
	return c.register * c.cycle
}

func (c *cpu) noOp() *int {
	var xAtFreq *int
	c.doCycle()
	if (c.cycle-20)%40 == 0 {
		s := c.getStrength()
		xAtFreq = &s
	}
	//log.Printf("End cycle %d: CRT draws pixel in position %d", c.cycle, c.register)
	return xAtFreq
}

func (c *cpu) addX(operand int) *int {
	var xAtFreq *int
	c.doCycle()
	if (c.cycle-20)%40 == 0 {
		s := c.getStrength()
		xAtFreq = &s
	}

	//log.Printf("End cycle %d: CRT draws pixel in position %d", c.cycle, c.register)
	c.doCycle()
	if (c.cycle-20)%40 == 0 {
		c.register += operand
		s := c.getStrength()
		xAtFreq = &s
	} else {
		c.register += operand
	}
	//log.Printf("End cycle %d: CRT draws pixel in position %d", c.cycle, c.register)
	return xAtFreq
}

func process(lines []string) any {
	freqSignals := make([]int, 0)
	c := cpu{register: 1, cycle: 1}
	for _, line := range lines {
		var command string
		var operand int
		fmt.Sscanf(line, "%s %d", &command, &operand)
		switch command {
		case "noop":
			xAtFreq := c.noOp()
			if xAtFreq != nil {
				freqSignals = append(freqSignals, *xAtFreq)
			}
		case "addx":
			xAtFreq := c.addX(operand)
			if xAtFreq != nil {
				freqSignals = append(freqSignals, *xAtFreq)
			}
		default:
		}
	}

	log.Printf("cycles: %d, register: %d", c.cycle, c.register)
	log.Printf("%v", freqSignals)
	totalSignalStrengths := 0
	for _, signal := range freqSignals {
		totalSignalStrengths += signal
	}
	return totalSignalStrengths
}

func (p Day10) PartA(lines []string) any {
	return process(lines)
}

func (p Day10) PartB(lines []string) any {
	return process(lines)
}
