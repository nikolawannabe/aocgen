package year2022

import (
	"log"
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
