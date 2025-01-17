package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example11 = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func TestDay11PartA(t *testing.T) {
	Init()
	input := strings.Split(example11, "\n")
	p := aoc.NewPuzzle(2022, 11)

	o := p.PartA(input)
	assert.Equal(t, 10605, o)
}

func TestDay11PartB(t *testing.T) {
	Init()
	input := strings.Split(example11, "\n")
	p := aoc.NewPuzzle(2022, 11)

	o := p.PartB(input)
	assert.Equal(t, 2713310158, o)
}
