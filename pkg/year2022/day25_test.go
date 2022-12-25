package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example25 = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func TestDay25PartA(t *testing.T) {
	Init()
	input := strings.Split(example25, "\n")
	p := aoc.NewPuzzle(2022, 25)

	o := p.PartA(input)
	assert.Equal(t, "2=-1=0", o)
}

func TestSnafuConversion(t *testing.T) {
	cases := []struct {
		snafu   string
		decimal int64
	}{
		{snafu: "1", decimal: 1},
		{snafu: "2", decimal: 2},
		{snafu: "1=", decimal: 3},
		{snafu: "1-", decimal: 4},
		{snafu: "1=-0-2", decimal: 1747},
		{snafu: "12111", decimal: 906},
		{snafu: "2=0=", decimal: 198},
		{snafu: "21", decimal: 11},
		{snafu: "2=01", decimal: 201},
		{snafu: "111", decimal: 31},
		{snafu: "20012", decimal: 1257},
		{snafu: "112", decimal: 32},
		{snafu: "1=-1=", decimal: 353},
		{snafu: "1-12", decimal: 107},
		{snafu: "12", decimal: 7},
		{snafu: "1=", decimal: 3},
		{snafu: "122", decimal: 37},
	}

	for _, test := range cases {
		o := snafuToDecimal(test.snafu)
		assert.Equal(t, test.decimal, o)
	}
}

func TestDay25PartB(t *testing.T) {
	Init()
	input := strings.Split(example25, "\n")
	p := aoc.NewPuzzle(2022, 25)

	p.PartB(input)
	// assert here
}
