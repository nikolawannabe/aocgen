package year2021

import (
	"log"
	"strconv"
	"strings"
)

type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	horizontal := 0
	depth := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		commands := strings.Split(line, " ")
		direction := commands[0]
		amount, _ := strconv.Atoi(commands[1])
		log.Printf("direction: %s", direction)
		log.Printf("amount: %d", amount)
		if direction == "forward" {
			horizontal = horizontal + amount
		} else if direction == "down" {
			depth = depth + amount
		} else if direction == "up" {
			depth = depth - amount
		}
		log.Printf("horizontal: %d, depth: %d", horizontal, depth)
	}
	final := horizontal * depth
	return final
}

func (p Day02) PartB(lines []string) any {
	horizontal := 0
	aim := 0
	depth := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		commands := strings.Split(line, " ")
		direction := commands[0]
		amount, _ := strconv.Atoi(commands[1])
		log.Printf("direction: %s", direction)
		log.Printf("amount: %d", amount)
		if direction == "forward" {
			horizontal = horizontal + amount
			depth = depth + (aim * amount)
		} else if direction == "down" {
			aim = aim + amount
		} else if direction == "up" {
			aim = aim - amount
		}
		log.Printf("horizontal: %d, depth: %d, aim; %d", horizontal, depth, aim)
	}
	final := horizontal * depth
	return final
}
