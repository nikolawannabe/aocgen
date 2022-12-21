package year2022

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example21 = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestDay21PartA(t *testing.T) {
	Init()
	input := strings.Split(example21, "\n")

	o := evalP1Root(0, input)
	assert.Equal(t, 152, o)
}

func TestDay21PartB(t *testing.T) {
	Init()
	input := strings.Split(example21, "\n")
	o := evalP2Root(0, input)
	assert.Equal(t, 301, o)
}
