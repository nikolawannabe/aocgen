package year2022

import (
	"sort"
	"strconv"
)

type Day01 struct{}

func getElfsCalories(lines []string, startIndex int) (int, int) {
	currentCalories := 0
	for index := startIndex; index < len(lines); index++ {
		num, err := strconv.Atoi(lines[index])
		if err != nil {
			return currentCalories, index + 1
		}
		currentCalories = currentCalories + num
	}
	return currentCalories, len(lines)
}

func getMostCalories(lines []string) int {
	mostCalories := 0
	for index := 0; index < len(lines); {
		elfsCalories, newIndex := getElfsCalories(lines, index)
		index = newIndex
		if elfsCalories > mostCalories {
			mostCalories = elfsCalories
		}
	}
	return mostCalories
}

func getEachElfsCalories(lines []string) []int {
	allElfCalories := make([]int, 0)
	for index := 0; index < len(lines); {
		elfCalories, newIndex := getElfsCalories(lines, index)
		index = newIndex
		allElfCalories = append(allElfCalories, elfCalories)
	}
	return allElfCalories
}

func getTop3ElfsCalories(lines []string) int {
	allElfCalories := getEachElfsCalories(lines)
	sort.Sort(sort.Reverse(sort.IntSlice(allElfCalories)))

	if len(allElfCalories) > 2 {
		return allElfCalories[0] + allElfCalories[1] + allElfCalories[2]
	}
	return 0
}

func (p Day01) PartA(lines []string) any {
	mostCalories := getMostCalories(lines)
	return mostCalories
}

func (p Day01) PartB(lines []string) any {
	return getTop3ElfsCalories(lines)
}
