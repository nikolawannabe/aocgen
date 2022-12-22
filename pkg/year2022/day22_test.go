package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example22 = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

func TestDay22PartA(t *testing.T) {
	Init()
	input := strings.Split(example22, "\n")
	p := aoc.NewPuzzle(2022, 22)

	o := p.PartA(input)
	assert.Equal(t, 6032, o)
}

func TestDay22PartB(t *testing.T) {
	Init()
	input := strings.Split(example22, "\n")
	p := aoc.NewPuzzle(2022, 22)

	p.PartB(input)
	// assert here
}
