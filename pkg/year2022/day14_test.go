package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example14 = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestDay14PartA(t *testing.T) {
	Init()
	input := strings.Split(example14, "\n")
	p := aoc.NewPuzzle(2022, 14)

	o := p.PartA(input)
	assert.Equal(t, 24, o)
}

func TestDay14PartB(t *testing.T) {
	Init()
	input := strings.Split(example14, "\n")
	p := aoc.NewPuzzle(2022, 14)

	o := p.PartB(input)
	assert.Equal(t, 93, o)
	// assert here
}
