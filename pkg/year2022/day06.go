package year2022

type Day06 struct{}

func findMarker(runes []rune, markerLen int) int {
	maybeSig := make([]rune, 0)
	for n, char := range runes {
		maybeSig = append(maybeSig, char)
		setLen := len(set(maybeSig))
		if setLen == markerLen {
			return n + 1
		}
		if len(maybeSig) > markerLen-1 {
			maybeSig = maybeSig[1:]
		}
	}
	return 0
}

func (p Day06) PartA(lines []string) any {
	sop := 4
	input := lines[0]
	runes := split(input)

	return findMarker(runes, sop)
}

func (p Day06) PartB(lines []string) any {
	som := 14
	input := lines[0]
	runes := split(input)

	return findMarker(runes, som)
}
