package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example19 = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestDay19PartA(t *testing.T) {
	Init()
	input := strings.Split(example19, "\n")
	p := aoc.NewPuzzle(2022, 19)

	o := p.PartA(input)
	assert.Equal(t, 33, o)
}

func TestDay19PartB(t *testing.T) {
	Init()
	input := strings.Split(example19, "\n")
	p := aoc.NewPuzzle(2022, 19)

	p.PartB(input)
	// assert here
}
