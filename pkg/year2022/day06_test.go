package year2022

import (
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

func TestDay06PartA(t *testing.T) {
	Init()
	input := []string{"bvwbjplbgvbhsrlpgdmjqwftvncz"}
	p := aoc.NewPuzzle(2022, 6)

	output := p.PartA(input)
	assert.Equal(t, 5, output, "first marker should be 5")

	input = []string{"nppdvjthqldpwncqszvftbrmjlhg"}
	output = p.PartA(input)
	assert.Equal(t, 6, output)

	input = []string{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}
	output = p.PartA(input)
	assert.Equal(t, 10, output)

	input = []string{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}
	output = p.PartA(input)
	assert.Equal(t, 11, output)
}

func TestDay06PartB(t *testing.T) {
	Init()
	input := []string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}
	p := aoc.NewPuzzle(2022, 6)

	output := p.PartB(input)
	assert.Equal(t, 19, output)
}
