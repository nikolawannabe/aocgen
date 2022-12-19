package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example18 = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestDay18PartA(t *testing.T) {
	Init()
	input := strings.Split(example18, "\n")
	p := aoc.NewPuzzle(2022, 18)

	o := p.PartA(input)
	assert.Equal(t, 64, o)
}

func TestDay18PartB(t *testing.T) {
	Init()
	input := strings.Split(example18, "\n")
	p := aoc.NewPuzzle(2022, 18)

	p.PartB(input)
	// assert here
}
