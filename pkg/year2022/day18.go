package year2022

import (
	"fmt"
)

type Day18 struct{}

type interval struct {
	Start int `json:"s"`
	End   int `json:"e"`
}

type face struct {
	X interval `json:"x"`
	Y interval `json:"y"`
	Z interval `json:"z"`
}

/* - `[x-1,x],[y-1,y],[z=z-1]` = bottom
- `[x-1,x],[y=y-1],[z-1,z]` = front
- `[x=x-1],[y-1,y],[z-1,z]` = right
- `[x-1,x],[y-1,y],[z=z]` = top
- `[x-1,x],[y=y],[z-1,z]` = behind
- `[x=x],[y-1,y],[z-1,z]` = left */

func makeFaces(x int, y int, z int) []face {
	faces := make([]face, 0)
	bottom := face{X: interval{Start: x - 1, End: x},
		Y: interval{Start: y - 1, End: y},
		Z: interval{Start: z - 1, End: z - 1}}
	front := face{X: interval{Start: x - 1, End: x},
		Y: interval{Start: y - 1, End: y - 1},
		Z: interval{Start: z - 1, End: z}}
	right := face{X: interval{Start: x - 1, End: x - 1},
		Y: interval{Start: y - 1, End: y},
		Z: interval{Start: z - 1, End: z}}
	top := face{X: interval{Start: x - 1, End: x},
		Y: interval{Start: y - 1, End: y},
		Z: interval{Start: z, End: z}}
	behind := face{X: interval{Start: x - 1, End: x},
		Y: interval{Start: y, End: y},
		Z: interval{Start: z - 1, End: z}}
	left := face{X: interval{Start: x, End: x},
		Y: interval{Start: y - 1, End: y},
		Z: interval{Start: z - 1, End: z}}
	faces = append(faces, bottom, front, right, top, behind, left)
	return faces
}

func (p Day18) PartA(lines []string) any {
	faceMap := make(map[face]int, 0)
	allFaces := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var x int
		var y int
		var z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		faces := makeFaces(x, y, z)
		//f, _ := json.MarshalIndent(faces, "", "  ")
		//log.Printf("%s", string(f))

		for _, face := range faces {
			allFaces++
			seenCount, pres := faceMap[face]
			if !pres {
				faceMap[face] = 1
			} else {
				seenCount++
				faceMap[face] = seenCount
			}
		}
	}

	surfaceArea := 0
	notSurfaceArea := 0
	for _, seenCount := range faceMap {
		if seenCount == 1 {
			surfaceArea++
		} else {
			notSurfaceArea++
		}
	}
	return surfaceArea
}

func (p Day18) PartB(lines []string) any {
	return "implement_me"
}
