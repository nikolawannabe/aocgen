package year2022

import (
	"log"
	"sort"
	"strconv"
)

func reverse[T any](s []T) []T {
	var rev []T
	for _, n := range s {
		rev = append([]T{n}, rev...)
	}
	return rev
}

func set[T comparable](s []T) []T {
	unique := make(map[T]bool, 0)
	for _, item := range s {
		unique[item] = true
	}
	out := make([]T, 0)
	for item, _ := range unique {
		out = append(out, item)
	}
	return out
}

func split(s string) []rune {
	rs := make([]rune, 0)
	for _, char := range s {
		rs = append(rs, char)
	}
	return rs
}

func concat(rs []rune) string {
	s := ""
	for _, char := range rs {
		s += string(char)
	}
	return s
}

func splitByEmpty(lines []string) [][]string {
	lineGroups := make([][]string, 0)
	lineGroup := make([]string, 0)
	for i, line := range lines {
		if len(line) == 0 {
			if len(lines)-1 == i || i == 0 {
				continue
			}
			lineGroups = append(lineGroups, lineGroup)
			lineGroup = make([]string, 0)
		}
		lineGroup = append(lineGroup, line)
	}
	if len(lineGroup) > 0 {
		lineGroups = append(lineGroups, lineGroup)
	}
	return lineGroups
}

func atoi(input string) int {
	o, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("%v", err)
	}
	return o
}

func itoa(input int) string {
	return strconv.Itoa(input)
}

type intervalsArray [][]int

func (intA intervalsArray) Len() int {
	return len(intA)
}

func (intA intervalsArray) Swap(i, j int) {
	intA[i], intA[j] = intA[j], intA[i]
}

func (intA intervalsArray) Less(i, j int) bool {
	return intA[i][0] < intA[j][0]
}

func merge(intervals [][]int) [][]int {

	intA := intervalsArray(intervals)

	sort.Sort(intA)

	intervalsSorted := [][]int(intA)

	var output [][]int
	currentIntervalStart := intervalsSorted[0][0]
	currentIntervalEnd := intervalsSorted[0][1]
	for j := 1; j < len(intervalsSorted); j++ {
		if currentIntervalEnd >= intervalsSorted[j][0] {
			if intervalsSorted[j][1] > currentIntervalEnd {
				currentIntervalEnd = intervalsSorted[j][1]
			}
		} else {
			output = append(output, []int{currentIntervalStart, currentIntervalEnd})
			currentIntervalStart = intervalsSorted[j][0]
			currentIntervalEnd = intervalsSorted[j][1]
		}
	}
	output = append(output, []int{currentIntervalStart, currentIntervalEnd})
	return output
}
