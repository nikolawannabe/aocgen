package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example07 = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestDay07PartA(t *testing.T) {
	Init()
	input := strings.Split(example07, "\n")
	p := aoc.NewPuzzle(2022, 7)

	output := p.PartA(input)
	assert.Equal(t, 95437, output)
}

func TestDay07PartB(t *testing.T) {
	Init()
	input := strings.Split(example07, "\n")
	p := aoc.NewPuzzle(2022, 7)

	output := p.PartB(input)
	assert.Equal(t, 24933642, output)
	// assert here
}
