package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay04PArtA(t *testing.T) {
	Init()
	input := strings.Split(example, "\n")
	p := aoc.NewPuzzle(2022, 4)

	output := p.PartA(input)
	assert.Equal(t, 2, output, "output should be 4")
}

func TestDay04PArtB(t *testing.T) {
	Init()
	input := strings.Split(example, "\n")
	p := aoc.NewPuzzle(2022, 4)

	output := p.PartB(input)
	assert.Equal(t, 4, output, "output should be 4")
}
