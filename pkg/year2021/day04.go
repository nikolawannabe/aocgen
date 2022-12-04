package year2021

import (
	"log"
	"strconv"
	"strings"
)

type Day04 struct{}

func (p Day04) PartA(lines []string) any {
	return "implement_me"
}

func splitToInts(line string) []int {
	stringNums := strings.Split(line, ",")
	ints := make([]int, 0)
	for _, stringNum := range stringNums {
		num, err := strconv.Atoi(stringNum)
		if err != nil {
			log.Printf("err: %v", err)
		}
		ints = append(ints, num)
	}
	return ints
}

func splitLinesToGroups(lines []string) [][][]int {
	outputGroups := make([][][]int, 0)
	currentLineGroup := make([][]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			outputGroups = append(outputGroups, currentLineGroup)
			currentLineGroup = make([][]int, 0)
		} else {
			lineGroup := splitToInts(line)
			currentLineGroup = append(currentLineGroup, lineGroup)
		}
	}
	return outputGroups
}

func makeMarkGroups(width int, n int) [][][]bool {
	groups := make([][][]bool, 0)
	for i := 0; i < n; i++ {
		group := make([][]bool, width)
		groups[i] = group
	}
	return groups
}

func markGroup(selected int, board [][]int, markBoard [][]bool) [][]bool {
	for rowNum, row := range board {
		for colNum, colVal := range row {
			if selected == colVal {
				markBoard[rowNum][colNum] = true
			}
		}
	}
	return markBoard
}

func testWin(markBoard [][]bool) bool {
	for _, row := range markBoard {
		rowWin := true
		for _, colVal := range row {
			if colVal == false {
				rowWin = false
			}
		}
		if rowWin {
			return true
		}
	}

	width := 5
	height := 5
	for x := 0; x < width; x++ {
		colWin := true
		for y := 0; y < height; y++ {
			if colWin == false {
				colWin = false
			}
		}
		if colWin {
			return true
		}
	}
	return false
}

func runCall(selected int, boards [][][]int, markerBoards [][][]bool) bool {
	for boardNum, board := range boards {
		markerBoards[boardNum] = markGroup(selected, board, markerBoards[boardNum])
		won := testWin(markerBoards[boardNum])
		if won {
			return true
		}
	}
	return false
}

func (p Day04) PartB(lines []string) any {
	drawOrder := splitToInts(lines[0])

	boards := splitLinesToGroups(lines[2:])
	markerBoards := makeMarkGroups(5, len(boards))

	wonSelected := 0
	for _, selected := range drawOrder {
		won := runCall(selected, boards, markerBoards)
		if won {
			wonSelected = selected
			break
		}
	}

	log.Printf("drawOrder: %v, boards: %d", drawOrder, len(boards))
	log.Printf("board1: %v", boards[0])
	log.Printf("win selection: %d", wonSelected)
	return 0
}
