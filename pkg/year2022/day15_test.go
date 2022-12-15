package year2022

import (
	"strings"
	"testing"

	"aocgen/pkg/aoc"

	"github.com/stretchr/testify/assert"
)

const example15 = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestDay15PartA(t *testing.T) {
	Init()
	input := strings.Split(example15, "\n")
	p := aoc.NewPuzzle(2022, 15)

	o := p.PartA(input)
	assert.Equal(t, 25, o)
}

// 4960344 wrong
// 4210673 too low
// 5186801 too high
// 6543782
// 7130999
// 7130999
func TestDay15PartB(t *testing.T) {
	Init()
	input := strings.Split(example15, "\n")
	p := aoc.NewPuzzle(2022, 15)

	p.PartB(input)
	// assert here
}
