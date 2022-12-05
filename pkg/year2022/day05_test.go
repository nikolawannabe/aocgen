package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const exampleDay05 = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestDay05PartA(t *testing.T) {
	Init()
	input := strings.Split(exampleDay05, "\n")
	p := aoc.NewPuzzle(2022, 5)

	output := p.PartA(input)
	assert.Equal(t, "CMZ", output)
}

func TestDay05PartB(t *testing.T) {
	Init()
	input := strings.Split(exampleDay05, "\n")
	p := aoc.NewPuzzle(2022, 5)

	output := p.PartB(input)
	assert.Equal(t, "MCD", output)
}
