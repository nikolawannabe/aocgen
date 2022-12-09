package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example09 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestDay09PartA(t *testing.T) {
	Init()
	input := strings.Split(example09, "\n")
	p := aoc.NewPuzzle(2022, 9)

	o := p.PartA(input)
	assert.Equal(t, 13, o)
}

const example09PartB = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestDay09PartB(t *testing.T) {
	Init()
	input := strings.Split(example09PartB, "\n")
	p := aoc.NewPuzzle(2022, 9)

	o := p.PartB(input)
	assert.Equal(t, 36, o)
}

func TestDay09PartB1(t *testing.T) {
	Init()
	input := strings.Split(example09, "\n")
	p := aoc.NewPuzzle(2022, 9)

	o := p.PartB(input)
	assert.Equal(t, 1, o)
}
