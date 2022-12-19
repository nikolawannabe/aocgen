package year2022

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitByEmpty(t *testing.T) {
	s := []string{"aaaaa", "bbbb", "", "cccc", "dddd"}
	groups := splitByEmpty(s)
	assert.Equal(t, 2, len(groups), "there should be 2 groups of strings since there's 1 empty line")

	s = []string{"aaaaa", "bbbb", "", "cccc", "dddd", "", "eee"}
	groups = splitByEmpty(s)
	assert.Equal(t, 3, len(groups), "there should be 3 groups of strings since there's 2 empty lines")

	s = []string{"", "aaaaa", "bbbb", "", "cccc", "dddd", "", "eee"}
	groups = splitByEmpty(s)
	assert.Equal(t, 3, len(groups), "newline at beginning should not affect group count")

	s = []string{"", "aaaaa", "bbbb", "", "cccc", "dddd", "", "eee", ""}
	groups = splitByEmpty(s)
	assert.Equal(t, 3, len(groups), "newline at beginning should not affect group count")

	assert.Equal(t, 2, len(groups[0]), "first group should have two lines")
	log.Printf("groups1: %v", groups[1])
	assert.Equal(t, 2, len(groups[1]), "second group should have two lines")
	assert.Equal(t, 1, len(groups[2]), "third group should have 1 line")
}

func TestMerge(t *testing.T) {
	i := [][]int{{3, 5}, {7, 9}, {1, 4}}
	e := [][]int{{1, 5}, {7, 9}}
	o := merge(i)
	assert.Equal(t, e, o)

	i = [][]int{{5, 3}, {7, 9}, {1, 4}}
	e = [][]int{{1, 5}, {7, 9}}
	o = merge(i)
	assert.Equal(t, e, o)

	i = [][]int{{230600, 41034}, {229367, 2541377}}
	e = [][]int{{41034, 2541377}}
	o = merge(i)
	assert.Equal(t, e, o)

	i = [][]int{{-1, 5}, {7, 9}, {1, 4}}
	e = [][]int{{-1, 5}, {7, 9}}
	o = merge(i)
	assert.Equal(t, e, o)

	i = [][]int{{-1, -5}, {-3, -5}}
	e = [][]int{{-5, -1}}
	o = merge(i)
	assert.Equal(t, e, o)

	i = [][]int{{-3, -5}, {7, 9}, {1, 4}}
	e = [][]int{{-5, -3}, {1, 4}, {7, 9}}
	o = merge(i)
	assert.Equal(t, e, o)
}
