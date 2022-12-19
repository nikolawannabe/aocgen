package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example17 = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestDay17PartA(t *testing.T) {
	Init()

	input := strings.Split(example17, "\n")
	p := aoc.NewPuzzle(2022, 17)

	o := p.PartA(input)
	assert.Equal(t, 3068, o)
}

func TestDay17PartB(t *testing.T) {
	Init()
	input := strings.Split(example17, "\n")
	p := aoc.NewPuzzle(2022, 17)

	p.PartB(input)
	// assert here
}
