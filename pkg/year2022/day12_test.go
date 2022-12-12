package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example12 = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestDay12PartA(t *testing.T) {
	Init()
	input := strings.Split(example12, "\n")
	p := aoc.NewPuzzle(2022, 12)

	o := p.PartA(input)
	assert.Equal(t, 31, o)
}

func TestDay12PartB(t *testing.T) {
	Init()
	input := strings.Split(example12, "\n")
	p := aoc.NewPuzzle(2022, 12)

	p.PartB(input)
	// assert here
}
