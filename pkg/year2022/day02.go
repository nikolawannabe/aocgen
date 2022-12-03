package year2022

import (
	"strings"
)

type Day02 struct {
}

type Data struct {
	MoveMap      map[string]string
	TypeScoreMap map[string]int
	OutcomeMap   map[string]string
	Defeats      map[string]string
	DefeatedBy   map[string]string
}

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
	Lose     = "lose"
	Win      = "win"
	Draw     = "draw"
)

func (p Day02) PartA(lines []string) any {
	maps := encodeCriteriaToMaps()
	totalScore := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		choices := strings.Split(line, " ")
		opponentMove := maps.MoveMap[choices[0]]
		selfMove := maps.MoveMap[choices[1]]
		selfScore := maps.scoreType(selfMove)
		winScore := maps.scoreWin(opponentMove, selfMove)
		totalScore = totalScore + selfScore + winScore
	}
	return totalScore
}

func (p Day02) PartB(lines []string) any {
	maps := encodeCriteriaToMaps()
	totalScore := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		choices := strings.Split(line, " ")
		opponentMove := maps.MoveMap[choices[0]]
		moveType := maps.OutcomeMap[choices[1]]
		selfMove := maps.getMove(opponentMove, moveType)
		selfScore := maps.scoreType(selfMove)
		winScore := maps.scoreWin(opponentMove, selfMove)
		totalScore = totalScore + selfScore + winScore
	}
	return totalScore
}

func encodeCriteriaToMaps() Data {
	d2 := Data{}
	// A: Rock
	// B: paper
	// C: scissors
	// X: rock
	// Y : paper
	// Z : scissors
	m := make(map[string]string, 0)
	m["A"] = Rock
	m["B"] = Paper
	m["C"] = Scissors
	m["X"] = Rock
	m["Y"] = Paper
	m["Z"] = Scissors
	d2.MoveMap = m

	// 1:Rock
	// 2: paper
	// 3: scissors
	t := make(map[string]int, 0)
	t[Rock] = 1
	t[Paper] = 2
	t[Scissors] = 3
	d2.TypeScoreMap = t

	// X means you need to lose,
	// Y means you need to end the round in a draw
	// Z means you need to win. Good luck!
	o := make(map[string]string, 0)
	o["X"] = Lose
	o["Y"] = Draw
	o["Z"] = Win
	d2.OutcomeMap = o

	// Rock defeats Scissors,
	// Scissors defeats Paper,
	// Paper defeats Rock.

	defeats := make(map[string]string, 0)
	defeats[Rock] = Scissors
	defeats[Scissors] = Paper
	defeats[Paper] = Rock
	d2.Defeats = defeats

	db := make(map[string]string, 0)
	db[Scissors] = Rock
	db[Paper] = Scissors
	db[Rock] = Paper
	d2.DefeatedBy = db
	return d2
}

func (p Data) scoreWin(opponent string, self string) int {
	// (0 if you lost,
	// 3 if the round was a draw,'
	// and 6 if you won

	if opponent == self {
		return 3
	}

	moveToWin, _ := p.DefeatedBy[opponent]
	if self == moveToWin {
		return 6
	}
	return 0
}

func (p Data) getMove(opponent string, outcome string) string {
	if outcome == Lose {
		return p.Defeats[opponent]
	}

	if outcome == Draw {
		return opponent
	}

	if outcome == Win {
		return p.DefeatedBy[opponent]
	}
	return ""
}

func (p Data) scoreType(choice string) int {
	return p.TypeScoreMap[choice]
}
