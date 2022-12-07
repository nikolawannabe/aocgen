package year2022

import (
	"log"
	"strings"
)

const upperAsciiDiff = 38
const lowerAsciiDiff = 97

type Day03 struct{}

func getPriority(char rune) int {
	asciiVal := int(char)
	outputVal := 0
	if asciiVal < 91 && asciiVal > 64 {
		//"A" = 65, want 27
		outputVal = asciiVal - upperAsciiDiff
	} else {
		outputVal = asciiVal - 97 + 1
	}
	return outputVal
}
func (p Day03) PartA(lines []string) any {
	totalPriorities := 0
	foundPriorities := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		midpoint := len(line) / 2
		if len(lines)%2 != 0 {
			midpoint = len(line) / 2
		}

		for i := midpoint; i < len(line); i++ {
			char := rune(line[i])
			if strings.Contains(line[0:midpoint], string(char)) {
				outputVal := getPriority(char)
				totalPriorities = totalPriorities + outputVal
				foundPriorities = foundPriorities + 1
				break
			}
		}
	}
	return totalPriorities
}

func reduceToUnique(line string) string {
	return concat(set(split(line)))
}

func (p Day03) PartB(lines []string) any {
	totalPriority := 0
	foundMap := make(map[int]int, 0)
	foundPriority := false
	foundPriorities := 0
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, char := range reduceToUnique(line) {
			foundMap[getPriority(char)] = foundMap[getPriority(char)] + 1
		}
		n := i + 1
		if n%3 == 0 {
			for priority, foundIn := range foundMap {
				if foundIn == 3 {
					totalPriority = totalPriority + priority
					foundPriority = true
					foundPriorities = foundPriorities + 1
					break
				}
			}
			if !foundPriority {
				log.Printf("didn't find priorities lines %d-%d", i-2, i)
			}
			foundMap = make(map[int]int, 0)
		}

	}
	return totalPriority
}
