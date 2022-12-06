package year2022

type Day06 struct{}

func (p Day06) PartA(lines []string) any {
	sop := 4
	input := lines[0]
	runes := split(input)

	maybeSig := make([]rune, 0)
	for n, char := range runes {
		maybeSig = append(maybeSig, char)
		setLen := len(set(maybeSig))
		if setLen == sop {
			return n + 1
		}
		if len(maybeSig) > sop-1 {
			maybeSig = maybeSig[1:]
		}
	}
	return 0
}

func (p Day06) PartB(lines []string) any {
	som := 14
	input := lines[0]
	runes := split(input)

	maybeSig := make([]rune, 0)
	for n, char := range runes {
		maybeSig = append(maybeSig, char)
		setLen := len(set(maybeSig))
		if setLen == som {
			return n + 1
		}
		if len(maybeSig) > som-1 {
			maybeSig = maybeSig[1:]
		}
	}
	return 0
}
