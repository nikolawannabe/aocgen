package year2022

import (
	"fmt"
	"log"
)

const (
	eliminated = '#'
	sensorIcon = 'S'
	beacon     = 'B'
)

type Day15 struct{}

type sensor struct {
	at     position
	beacon position
}

type tunnels struct {
	xpairs     [][]int
	sensors    map[position]sensor
	beacons    map[position]bool
	eliminated map[position]bool
	upperLeft  position
	lowerRight position
	y          int
	clamp      *int
}

func (t *tunnels) printTunnel() {
	output := ""
	if t.clamp != nil {
		t.upperLeft.x = 0
		t.upperLeft.y = 0
		t.lowerRight.x = *t.clamp
		t.lowerRight.y = *t.clamp
	}
	for y := t.upperLeft.y - 1; y < t.lowerRight.y; y++ {
		for x := t.upperLeft.x - 1; x < t.lowerRight.x; x++ {
			if y == t.upperLeft.y-1 {
				xStr := itoa(x)
				output += xStr[len(xStr)-1:]
				continue
			}
			if x == t.upperLeft.x-1 {
				yStr := itoa(y)
				output += yStr[len(yStr)-1:]
				continue
			}
			p := position{x: x, y: y}
			if _, pres := t.sensors[p]; pres {
				output += string(sensorIcon)
			} else if _, pres := t.beacons[p]; pres {
				output += string(beacon)
			} else if _, pres := t.eliminated[p]; pres {
				output += string(eliminated)
			} else {
				output += "."
			}
		}
		log.Printf("%s", output)
		output = ""
	}
	log.Printf("\n")
}

func (t *tunnels) maybeExpandBounds(p position) {
	if p.x < t.upperLeft.x {
		t.upperLeft.x = p.x
	}
	if p.y < t.upperLeft.y {
		t.upperLeft.y = p.y
	}
	if p.x > t.lowerRight.x {
		t.lowerRight.x = p.x
	}
	if p.y > t.lowerRight.y {
		t.lowerRight.y = p.y
	}
}

func (t *tunnels) updateBounds(s sensor) {
	t.maybeExpandBounds(s.at)
	t.maybeExpandBounds(s.beacon)
}

func getMhd(s sensor) int {
	return abs(s.at.x-s.beacon.x) + abs(s.at.y-s.beacon.y)
}

func (t *tunnels) eliminateArea(s sensor) {
	mhd := getMhd(s)
	y := s.at.y - mhd
	for xd := 0; xd <= mhd; xd++ {
		x1 := s.at.x - xd
		x2 := s.at.x + xd
		for x := x1; x <= x2; x++ {
			if t.clamp != nil && (x > *t.clamp || x < 0) {
				continue
			}
			p := position{x: x, y: y}
			if y == t.y {
				t.eliminated[p] = true
			}

		}
		y++
	}
	for xd := mhd - 1; xd >= 0; xd-- {
		x1 := s.at.x - xd
		x2 := s.at.x + xd
		for x := x1; x <= x2; x++ {
			if t.clamp != nil && (x > *t.clamp || x < 0) {
				continue
			}
			p := position{x: x, y: y}
			if y == t.y {
				t.eliminated[p] = true
			}
		}
		y++
	}
}

func (t *tunnels) countEliminated() int {
	sumEliminated := 0
	//output := ""

	for pos, _ := range t.eliminated {
		//log.Printf("p: %d, %d", pos.x, pos.y)
		_, beaconPres := t.beacons[pos]
		_, sensorPres := t.sensors[pos]
		if !beaconPres && !sensorPres {
			sumEliminated++
		} else {
			log.Printf("skipping %#v because it was a thing", pos)
		}
	}
	return sumEliminated
}

func (t *tunnels) getXd(s sensor, mhd int) int {
	return mhd - abs(s.at.y-t.y)
}

func (t *tunnels) getXPair(s sensor, xd int) []int {
	return []int{s.at.x - xd, s.at.x + xd}
}

func (t *tunnels) sumIntervals(intervals [][]int) int {
	sum := 0
	for _, pair := range intervals {
		if pair[0] <= 0 && pair[1] >= 0 {
			pairSum := abs(pair[1]-pair[0]) + 1
			//log.Printf("%d - %d = %d", pair[1], pair[0], pairSum)
			sum += pairSum
		} else {
			pairSum := abs(pair[1]-pair[0]) + 1
			//log.Printf("%d - %d = %d", pair[1], pair[0], pairSum)
			sum += pairSum
		}
	}
	return sum
}

func (t *tunnels) checkPair(x int, y int, start int, end int) int {
	if y == t.y && start <= x && x <= end {
		return 1
	}
	return 0
}

func (t *tunnels) findAffectedStuff(intervals [][]int) int {
	//log.Printf("merged pairs: %v", intervals)
	itemsOnLine := make(map[position]bool, 0)
	for _, s := range t.sensors {
		if t.clamp != nil && (s.at.x > *t.clamp || s.at.x < 0) {
			continue
		}
		for _, intervalPair := range intervals {
			found := t.checkPair(s.at.x, s.at.y, intervalPair[0], intervalPair[1])
			if found == 1 {
				itemsOnLine[position{x: s.at.x, y: s.at.y}] = true
				//log.Printf("sensor %d, %d; interval %d, %d", s.at.x, s.at.y, intervalPair[0], intervalPair[1])
				break
			}
		}
		for _, intervalPair := range intervals {
			found := t.checkPair(s.beacon.x, s.beacon.y, intervalPair[0], intervalPair[1])
			if found == 1 {
				itemsOnLine[position{x: s.beacon.x, y: s.beacon.y}] = true
				//log.Printf("beacon %d, %d; interval %d, %d", s.beacon.x, s.beacon.y, intervalPair[0], intervalPair[1])
				break
			}
		}
	}
	return map_len(itemsOnLine)
}

func (t *tunnels) getTotalByBruteForce() int {
	t.eliminated = make(map[position]bool, 0)
	for _, s := range t.sensors {
		t.updateBounds(s)

		mhd := getMhd(s)

		t.eliminateArea(s)
		//log.Printf("completed sensor %#v", s.at)
		if s.at.y-mhd <= t.y && t.y <= s.at.y+mhd {
			t.maybeExpandBounds(position{x: s.at.x + mhd, y: s.at.y})
			t.maybeExpandBounds(position{x: s.at.x - mhd, y: s.at.y})
			//log.Printf("updating eliminations of sensor: %#v", s)
		} else {
			t.maybeExpandBounds(position{x: s.beacon.x, y: s.beacon.y})
			//log.Printf("skipping sensor as it doesn't cross output %d %#v", t.y, s)
		}
	}

	t.printTunnel()

	totalByEliminated := t.countEliminated()
	//log.Printf("total by eliminated: %d", totalByEliminated)
	return totalByEliminated
}

func getTunnels(yLine int, lines []string) tunnels {
	t := tunnels{}
	sensors := make(map[position]sensor, 0)
	beacons := make(map[position]bool, 0)
	t.sensors = sensors
	t.beacons = beacons
	t.y = yLine
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var sx int
		var sy int
		var bx int
		var by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		b := position{x: bx, y: by}
		s := sensor{at: position{x: sx, y: sy}, beacon: b}
		//t.updateBounds(s)

		mhd := getMhd(s)

		t.sensors[s.at] = s
		t.beacons[b] = true

		//log.Printf("line %d; completed sensor at %d, %d", i, s.at.x, s.at.y)
		if s.at.y-mhd <= t.y && t.y <= s.at.y+mhd {
			t.maybeExpandBounds(position{x: s.at.x + mhd, y: s.at.y})
			t.maybeExpandBounds(position{x: s.at.x - mhd, y: s.at.y})
			//log.Printf("updating eliminations of %d sensor of %d: %#v", i+1, len(lines), s)
		} else {
			t.maybeExpandBounds(position{x: s.beacon.x, y: s.beacon.y})
			//log.Printf("skipping sensor %d as it doesn't cross output %d %#v", i, t.y, s)
		}
	}
	return t
}

func (t *tunnels) getIntervals() [][]int {
	eliminated := make(map[position]bool, 0)
	t.eliminated = eliminated
	t.xpairs = make([][]int, 0)
	for _, s := range t.sensors {
		mhd := getMhd(s)

		xd := t.getXd(s, mhd)

		if xd > 0 {
			xpair := t.getXPair(s, xd)
			t.xpairs = append(t.xpairs, xpair)
		}

		if s.at.y-mhd <= t.y && t.y <= s.at.y+mhd {
			t.maybeExpandBounds(position{x: s.at.x + mhd, y: s.at.y})
			t.maybeExpandBounds(position{x: s.at.x - mhd, y: s.at.y})
			//log.Printf("updating eliminations of sensor: %#v", s.at)
		} else {
			t.maybeExpandBounds(position{x: s.beacon.x, y: s.beacon.y})
			//log.Printf("skipping sensor %#v as it doesn't cross output ", s.at)
		}

		//log.Printf("completed sensor at %#v", s.at)
	}
	//log.Printf("%#v", t.xpairs)
	mergedPairs := merge(t.xpairs)
	//log.Printf("%#v", mergedPairs)
	//log.Printf("intervals: %v", mergedPairs)
	return mergedPairs
}

func (t *tunnels) getTotalByIntervals() int {
	mergedPairs := t.getIntervals()
	stuffOnLine := t.findAffectedStuff(mergedPairs)
	log.Printf("y: %d", t.y)
	log.Printf("sensors+beacons on y: %d", stuffOnLine)
	totalByIntervals := t.sumIntervals(mergedPairs) - stuffOnLine
	log.Printf("total by intervals: %d", totalByIntervals)
	return totalByIntervals
}

func getTotal(y int, input []string) (int, int) {
	t := getTunnels(y, input)
	bruteForce := t.getTotalByBruteForce()
	byIntervals := t.getTotalByIntervals()
	return bruteForce, byIntervals
}

func findIntervalCountAndY(bound int, input []string) ([][]int, int) {
	t := getTunnels(0, input)
	t.clamp = &bound
	outputPairs := make([][][]int, 0)
	outputY := make([]int, 0)
	for y := 0; y < bound; y++ {
		t.y = y
		pairs := t.getIntervals()
		//log.Printf("%d: %#v", y, pairs)
		if len(pairs) == 2 {
			outputPairs = append(outputPairs, pairs)
			outputY = append(outputY, y)
		}
	}

	reducedPairs, reducedYs := reduceIntervals(outputPairs, outputY)
	if len(reducedPairs) < 1 {
		log.Printf("nope couldn't find it")
		return [][]int{}, 0
	}
	if len(reducedPairs) > 1 {
		log.Printf("found multiple! %#v, %#v", reducedPairs, reducedYs)
		return [][]int{}, 0
	}
	return reducedPairs[0], reducedYs[0]
}

func reduceIntervals(pairs [][][]int, ys []int) ([][][]int, []int) {
	output := make([][][]int, 0)
	outputys := make([]int, 0)
	for i, pair := range pairs {
		if pair[0][1]+1 != pair[1][0] {
			output = append(output, pair)
			outputys = append(outputys, ys[i])
		}
	}
	return output, outputys
}

func findXIntersection(pairs [][]int) []int {
	exes := make([]int, 0)
	for x := pairs[0][1] + 1; x < pairs[1][0]; x++ {
		exes = append(exes, x)
	}
	return exes
}

func findFrequency(x int, y int) int {
	return (x * 4000000) + y
}

func (p Day15) PartA(lines []string) any {
	y := 2000000

	t := getTunnels(y, lines)
	byIntervals := t.getTotalByIntervals()
	return byIntervals
}

func (p Day15) PartB(lines []string) any {
	pairs, y := findIntervalCountAndY(4000000, lines)
	log.Printf("found pairs %#v", pairs)
	exes := findXIntersection(pairs)
	if len(exes) > 1 {
		log.Printf("too many exes: %#v", exes)
	}
	frequency := findFrequency(exes[0], y)
	return frequency
}
