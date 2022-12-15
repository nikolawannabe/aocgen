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
}

func (t *tunnels) printTunnel() {
	output := ""

	for y := t.upperLeft.y; y < t.lowerRight.y; y++ {
		for x := t.upperLeft.x - 1; x < t.lowerRight.x; x++ {
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

	for range t.eliminated {
		sumEliminated++
	}
	return sumEliminated
	/*
			for x := t.upperLeft.x; x < t.lowerRight.x; x++ {
				p := position{x: x, y: t.y}
				if _, pres := t.sensors[p]; pres {
					output += string(sensorIcon)
				} else if _, pres := t.beacons[p]; pres {
					output += string(beacon)
				} else if _, pres := t.eliminated[p]; pres {
					output += string(eliminated)
					sumEliminated++
				} else {
					output += "."
				}
			}
		log.Printf("%s", output)
		return sumEliminated */
}

func (t *tunnels) getXd(s sensor, mhd int) int {
	return mhd - abs(t.y-s.at.y)
}

func (t *tunnels) getXPair(s sensor, xd int) []int {
	return []int{s.at.x - xd, s.at.x + xd}
}

func (t *tunnels) sumIntervals(intervals [][]int) int {
	sum := 0
	for _, pair := range intervals {
		sum += abs(pair[1] - pair[0])
	}
	return sum
}

func (p Day15) PartA(lines []string) any {
	t := tunnels{}
	sensors := make(map[position]sensor, 0)
	beacons := make(map[position]bool, 0)
	t.sensors = sensors
	t.beacons = beacons
	t.y = 2000000
	eliminated := make(map[position]bool, 0)
	t.eliminated = eliminated
	t.xpairs = make([][]int, 0)
	for i, line := range lines {
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
		t.updateBounds(s)

		mhd := getMhd(s)

		xd := t.getXd(s, mhd)

		t.sensors[s.at] = s
		t.beacons[b] = true

		xpair := t.getXPair(s, xd)
		t.xpairs = append(t.xpairs, xpair)
		//t.eliminateArea(s)
		if s.at.y-mhd <= t.y && t.y <= s.at.y+mhd {
			t.maybeExpandBounds(position{x: s.at.x + mhd, y: s.at.y})
			t.maybeExpandBounds(position{x: s.at.x - mhd, y: s.at.y})
			log.Printf("updating eliminations of %d sensor of %d: %#v", i+1, len(lines), s)
		} else {
			t.maybeExpandBounds(position{x: s.beacon.x, y: s.beacon.y})
			log.Printf("skipping sensor %d as it doesn't cross output %d %#v", i, t.y, s)
		}
	}
	log.Printf("%#v", t.xpairs)
	mergedPairs := merge(t.xpairs)
	log.Printf("%#v", mergedPairs)

	//t.printTunnel()
	//return t.countEliminated()

	return t.sumIntervals(mergedPairs)
}

func (p Day15) PartB(lines []string) any {
	return "implement_me"
}
