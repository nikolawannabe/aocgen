package year2022

import "log"

type Day20 struct{}

type wrappingSlice []int

func (w *wrappingSlice) set(position int, value int) {
	length := len(*w)
	if value == 0 {
		return
	}
	newPos := length + (position % length)
	if position > -1 {
		newPos = (position % length)
	}
	tail := []int{}
	if newPos < len(*w) {
		tail = []int(*w)[newPos+1:]
	}
	head := []int(*w)[1 : newPos+1]
	log.Printf("head %#v", head)
	log.Printf("value %d", value)
	log.Printf("tail %#v", tail)
	*w = append(head, append([]int{value}, tail...)...)
}

func (w *wrappingSlice) get(position int) int {
	length := len(*w)
	if position == 0 {
		return []int(*w)[position]
	}
	if position > 0 {
		newPos := (position % length)
		return []int(*w)[newPos]
	}
	newPos := length + (position % length)
	return []int(*w)[newPos]
}

func mix(index int, in []int, out wrappingSlice) wrappingSlice {
	log.Printf("out: %#v", out)
	value := in[index]
	if value == 0 {
		return out
	}
	newPosition := index + value
	out.set(newPosition, value)
	return out
}

func getZeroIndex(out []int) int {
	for i, val := range out {
		if val == 0 {
			return i
		}
	}
	return 0
}

func (p Day20) PartA(lines []string) any {
	count := len(lines)
	out := wrappingSlice(make([]int, count))
	in := make([]int, 0)
	for i, line := range lines {
		in = append(in, atoi(line))
		out[i] = atoi(line)
	}

	log.Printf("%#v", out)
	for i, _ := range in {
		log.Printf("== i %d ==", i)
		out = mix(i, in, out)
		log.Printf("%#v", out)
	}
	zeroIndex := getZeroIndex(out)
	return out.get(zeroIndex+1000) + out.get(zeroIndex+2000) + out.get(zeroIndex+3000)
}

func (p Day20) PartB(lines []string) any {
	return "implement_me"
}
