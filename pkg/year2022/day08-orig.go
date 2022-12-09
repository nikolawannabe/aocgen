package year2022

import (
	"log"
	"strconv"
)

type Day08Orig struct{}

type visible bool

type tree struct {
	x int
	y int
}

func isShorter(fortTree int, hiderTree int) bool {
	if hiderTree < fortTree {
		// log.Printf("comparing %d to %d: shorter", fortTree, hiderTree)
		return true
	}
	return false
}

func isVisible(t tree, forestMap map[tree]int, xbound int, ybound int) visible {
	heightOfTarget := forestMap[t]
	if t.y == ybound-1 || t.y == 0 {
		return visible(true)
	}
	if t.x == xbound-1 || t.x == 0 {
		return visible(true)
	}

	shorter := true
	for yS := 0; yS < t.y; yS++ {
		shorter = shorter && isShorter(heightOfTarget, forestMap[tree{y: yS, x: t.x}])
	}
	if shorter {
		return visible(true)
	}

	shorter = true
	for yS := t.y + 1; yS < ybound; yS++ {
		shorter = shorter && isShorter(heightOfTarget, forestMap[tree{y: yS, x: t.x}])
	}
	if shorter {
		return visible(true)
	}

	shorter = true
	for xS := 0; xS < t.x; xS++ {
		shorter = shorter && isShorter(heightOfTarget, forestMap[tree{y: t.y, x: xS}])
	}
	if shorter {
		return visible(true)
	}

	shorter = true
	for xS := t.x + 1; xS < xbound; xS++ {
		shorter = shorter && isShorter(heightOfTarget, forestMap[tree{y: t.y, x: xS}])
	}
	if shorter {
		return visible(true)
	}

	return visible(false)
}

func lookDown(t tree, forestMap map[tree]int, xbound int, ybound int) int {
	heightOfTarget := forestMap[t]
	shorter := true
	visibleTrees := 0
	for yS := t.y + 1; yS < ybound && shorter; yS++ {
		visibleTrees++
		shorter = isShorter(heightOfTarget, forestMap[tree{y: yS, x: t.x}])
	}
	return visibleTrees
}

func lookUp(t tree, forestMap map[tree]int, xbound int, ybound int) int {
	heightOfTarget := forestMap[t]
	shorter := true
	visibleTrees := 0
	for yS := t.y - 1; yS > -1 && shorter; yS-- {
		visibleTrees++
		shorter = isShorter(heightOfTarget, forestMap[tree{y: yS, x: t.x}])
	}
	return visibleTrees
}

func lookLeft(t tree, forestMap map[tree]int, xbound int, ybound int) int {
	heightOfTarget := forestMap[t]
	shorter := true
	visibleTrees := 0
	for xS := t.x - 1; xS > -1 && shorter; xS-- {
		visibleTrees++
		shorter = isShorter(heightOfTarget, forestMap[tree{y: t.y, x: xS}])
	}
	return visibleTrees
}

func lookRight(t tree, forestMap map[tree]int, xbound int, ybound int) int {
	heightOfTarget := forestMap[t]
	shorter := true
	visibleTrees := 0
	for xS := t.x + 1; xS < xbound && shorter; xS++ {
		visibleTrees++
		shorter = isShorter(heightOfTarget, forestMap[tree{y: t.y, x: xS}])
	}
	return visibleTrees
}

func getVisibilityScore(t tree, forestMap map[tree]int, xbound int, ybound int) int {
	return lookUp(t, forestMap, xbound, ybound) *
		lookDown(t, forestMap, xbound, ybound) *
		lookLeft(t, forestMap, xbound, ybound) *
		lookRight(t, forestMap, xbound, ybound)
}

func (p Day08Orig) PartA(lines []string) any {
	forestWidth := len(lines[0])
	forestHeight := len(lines)
	heightsMap := make(map[tree]int, 0)

	for y, line := range lines {
		for x, char := range line {
			height, _ := strconv.Atoi(string(char))
			heightsMap[tree{x: x, y: y}] = height
		}
	}

	for y := 0; y < forestHeight; y++ {
		output := ""
		for x := 0; x < forestWidth; x++ {
			output += strconv.Itoa(heightsMap[tree{x: x, y: y}])
		}
		log.Printf("%s", output)
	}

	visibleTrees := 0
	for y := 0; y < forestHeight; y++ {
		output := ""
		for x := 0; x < forestWidth; x++ {
			v := isVisible(tree{x: x, y: y}, heightsMap, forestWidth, forestHeight)
			if v {
				visibleTrees++
				output += "v"
			} else {
				output += "o"
			}
		}
		log.Printf("%s", output)
	}
	return visibleTrees
}

func (p Day08Orig) PartB(lines []string) any {
	forestWidth := len(lines[0])
	forestHeight := len(lines)
	heightsMap := make(map[tree]int, 0)

	for y, line := range lines {
		for x, char := range line {
			height, _ := strconv.Atoi(string(char))
			heightsMap[tree{x: x, y: y}] = height
		}
	}

	for y := 0; y < forestHeight; y++ {
		output := ""
		for x := 0; x < forestWidth; x++ {
			output += strconv.Itoa(heightsMap[tree{x: x, y: y}])
		}
		log.Printf("%s", output)
	}

	highestScore := 0
	for y := 0; y < forestHeight; y++ {
		output := ""
		for x := 0; x < forestWidth; x++ {
			score := getVisibilityScore(tree{x: x, y: y}, heightsMap, forestWidth, forestHeight)
			if score > highestScore {
				highestScore = score
			}
			output += strconv.Itoa(score)
		}
		log.Printf("%s", output)
	}
	return highestScore
}
