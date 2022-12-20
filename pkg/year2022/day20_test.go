package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example20 = `1
2
-3
3
-2
0
4`

func TestWrappingSet(t *testing.T) {
	ws := wrappingSlice([]int{1, 2, 3, 4})
	ws.set(-1, 1)
	assert.Equal(t, 1, ws[3])

	ws = wrappingSlice([]int{1, 2, 3, 4})
	ws.set(-6, 1)
	assert.Equal(t, 1, ws[2])

	ws = wrappingSlice([]int{1, 2, 3, 4})
	ws.set(5, 1)
	assert.Equal(t, 1, ws[1])

	ws = wrappingSlice([]int{1, 2, 3, 4})
	ws.set(2, 1)
	assert.Equal(t, 1, ws[2])

	ws = wrappingSlice([]int{1, 2, 3, 4})
	o := ws.get(5)
	assert.Equal(t, 2, o)
	o = ws.get(-6)
	assert.Equal(t, 3, o)
}

// -13757 wrong

func TestDay20PartA(t *testing.T) {
	Init()
	input := strings.Split(example20, "\n")
	p := aoc.NewPuzzle(2022, 20)

	o := p.PartA(input)
	assert.Equal(t, 3, o)
}

func TestDay20PartB(t *testing.T) {
	Init()
	input := strings.Split(example20, "\n")
	p := aoc.NewPuzzle(2022, 20)

	p.PartB(input)
	// assert here
}
