package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example08 = `30373
25512
65332
33549
35390`

func TestDay08PartA(t *testing.T) {
	Init()
	input := strings.Split(example08, "\n")
	p := aoc.NewPuzzle(2022, 8)

	output := p.PartA(input)
	assert.Equal(t, 21, output)
}

func TestDay08PartB(t *testing.T) {
	Init()
	input := strings.Split(example08, "\n")
	p := aoc.NewPuzzle(2022, 8)

	output := p.PartB(input)
	assert.Equal(t, 8, output)
}
