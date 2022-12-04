package year2022

import (
	"strconv"
	"strings"
)

type Day04 struct{}

func getIntPair(strPair []string) (int, int) {
	fst, _ := strconv.Atoi(strPair[0])
	snd, _ := strconv.Atoi(strPair[1])
	return fst, snd
}
func (p Day04) PartA(lines []string) any {
	fullyContained := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		strPairs := strings.Split(line, ",")
		fstStrPair := strings.Split(strPairs[0], "-")
		sndStrPair := strings.Split(strPairs[1], "-")
		fst1, snd1 := getIntPair(fstStrPair)
		fst2, snd2 := getIntPair(sndStrPair)
		if snd2 >= snd1 {
			if fst1 >= fst2 {
				fullyContained = fullyContained + 1
				continue
			}
		}
		if snd2 <= snd1 {
			if fst1 <= fst2 {
				fullyContained = fullyContained + 1
				continue
			}
		}
	}
	return fullyContained
}

func (p Day04) PartB(lines []string) any {

	overlapping := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		strPairs := strings.Split(line, ",")
		fstStrPair := strings.Split(strPairs[0], "-")
		sndStrPair := strings.Split(strPairs[1], "-")
		fst1, snd1 := getIntPair(fstStrPair)
		fst2, snd2 := getIntPair(sndStrPair)
		if fst2 <= snd2 && snd1 >= fst2 && snd1 <= snd2 {
			overlapping = overlapping + 1
			continue
		}
		if fst1 <= snd1 && snd2 >= fst1 && snd2 <= snd1 {
			overlapping = overlapping + 1
			continue
		}
	}
	return overlapping
}
