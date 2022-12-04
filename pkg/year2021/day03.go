package year2021

import (
	"log"
	"strconv"
)

type Day03 struct{}

func (p Day03) PartA(lines []string) any {
	zeroBits := make([]int, 12)
	oneBits := make([]int, 12)
	for _, line := range lines {
		for index, char := range line {
			if char == '0' {
				zeroBits[index] = zeroBits[index] + 1
			} else {
				oneBits[index] = oneBits[index] + 1
			}

		}
	}

	gammaBits := make([]rune, 12)
	epsilonBits := make([]rune, 12)
	for index, _ := range zeroBits {
		if zeroBits[index] > oneBits[index] {
			gammaBits[index] = '0'
			epsilonBits[index] = '1'
		} else {
			gammaBits[index] = '1'
			epsilonBits[index] = '0'
		}
	}

	log.Printf("gammaBits: %s", string(gammaBits))
	log.Printf("epsilonBits: %s", string(epsilonBits))
	gamma, err := strconv.ParseInt(string(gammaBits), 2, 64)
	if err != nil {
		log.Printf("err: %v", err)
	}
	epsilon, err := strconv.ParseInt(string(epsilonBits), 2, 64)
	if err != nil {
		log.Printf("err: %v", err)
	}
	log.Printf("gamma: %d", gamma)
	log.Printf("epsilon: %d", epsilon)
	return gamma * epsilon
}

func findCriteriaLinesInBit(lines []string, bitToCheck int, reverseCriteria bool) []string {
	zeroBits := 0
	oneBits := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		char := line[bitToCheck]
		if char == '0' {
			zeroBits++
		} else {
			oneBits++
		}
	}
	newLines := make([]string, 0)
	favorZero := false
	if zeroBits > oneBits {
		favorZero = true
	}
	favorZero = favorZero && reverseCriteria
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		char := line[bitToCheck]
		if !reverseCriteria {
			if zeroBits > oneBits && char == '0' {
				newLines = append(newLines, line)
			}
			if oneBits >= zeroBits && char == '1' {
				newLines = append(newLines, line)
			}

		} else {
			if zeroBits <= oneBits && char == '0' {
				newLines = append(newLines, line)
			}
			if oneBits < zeroBits && char == '1' {
				newLines = append(newLines, line)
			}
		}
	}
	return newLines
}

func findOxygenBits(lines []string) string {
	for index := 0; index < 12 && len(lines) > 1; index++ {
		lines = findCriteriaLinesInBit(lines, index, false)
	}
	return lines[0]
}

func findScrubberBits(lines []string) string {
	for index := 0; index < 12 && len(lines) > 1; index++ {
		lines = findCriteriaLinesInBit(lines, index, true)
	}
	return lines[0]
}

func (p Day03) PartB(lines []string) any {
	oxygenBits := findOxygenBits(lines)
	scrubberBits := findScrubberBits(lines)
	oxygen, err := strconv.ParseInt(string(oxygenBits), 2, 64)
	if err != nil {
		log.Printf("error: %v", err)
	}
	scrubber, err := strconv.ParseInt(string(scrubberBits), 2, 64)
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("oxygen: %d, scrubber: %d", oxygen, scrubber)
	return oxygen * scrubber
}
