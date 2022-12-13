package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example13 = ``

func TestDay13PartA(t *testing.T) {
	Init()
	input := strings.Split(example13, "\n")
	p := aoc.NewPuzzle(2022, 13)

	o := p.PartA(input)
	assert.Equal(t, 13, o)
}

func TestDay13PartB(t *testing.T) {
	Init()
	input := strings.Split(example13, "\n")
	p := aoc.NewPuzzle(2022, 13)

	o := p.PartB(input)
	assert.Equal(t, 140, o)
}
