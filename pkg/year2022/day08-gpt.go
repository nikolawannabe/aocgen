package year2022

import (
	"log"
	"strconv"
)

type Day08 struct{}

type Tree struct {
	Height  int
	Visible bool
}

// Forest represents the grid of trees
type Forest struct {
	Trees [][]Tree
	Width int
}

// CountVisibleTrees returns the number of trees in the forest that are visible from the outside of the grid
func (f *Forest) CountVisibleTrees() int {
	// Initialize the count of visible trees to the number of trees on the edge of the grid
	count := f.Width*2 + len(f.Trees)*2 - 2

	// Iterate over the rows and columns of the grid of trees
	for row := 1; row < len(f.Trees)-1; row++ {
		for col := 1; col < f.Width-1; col++ {
			// Check if the current tree is visible from the outside of the grid
			if f.isTreeVisible(row, col) {
				// If the tree is visible, increment the count of visible trees
				f.Trees[row][col].Visible = true
				count++
			}
		}
	}

	// Return the final count of visible trees
	return count
}

// isTreeVisible returns true if the tree at the given row and column is visible from the outside of the grid
func (f *Forest) isTreeVisible(row, col int) bool {
	// Get the height of the current tree
	height := f.Trees[row][col].Height

	// Iterate over the other trees in the same row and column as the current tree
	for r := 0; r < len(f.Trees); r++ {
		if r == row {
			continue
		}
		if f.Trees[r][col].Height > height {
			return false
		}
	}
	for c := 0; c < f.Width; c++ {
		if c == col {
			continue
		}
		if f.Trees[row][c].Height > height {
			return false
		}
	}

	// If all other trees in the same row and column are shorter than the current tree, return true
	return true
}

func (p Day08) PartA(lines []string) any {
	forestHeight := len(lines)
	forestWidth := len(lines[0])
	trees := make([][]Tree, len(lines))

	for y, line := range lines {
		for x, char := range line {
			if trees[x] == nil {
				trees[x] = make([]Tree, len(line))
			}
			height, _ := strconv.Atoi(string(char))
			t := Tree{Height: height, Visible: false}
			trees[x][y] = t
		}
	}
	forest := Forest{Trees: trees, Width: len(trees[0])}

	count := forest.CountVisibleTrees()

	for y := 0; y < forestHeight; y++ {
		output := ""
		for x := 0; x < forestWidth; x++ {
			v := forest.Trees[x][y].Visible
			if v {
				output += "v"
			} else {
				output += "o"
			}
		}
		log.Printf("%s", output)
	}
	return count
}

func (p Day08) PartB(lines []string) any {
	return 0
}
