/*
In the first blueprint example:
o := 4t
c := 8t (2o) // costs 2 ore
obs := 12t + 14*8t // (3 ore and 14 clay.)
g := 4t + 7*14*8t // costs 2 ore and 7 obsidian.

convert each ore resources to a base unit.  In this case "t".
switch to the next up resource type in the chain when the total t is equal or less than the production per t
*/
package year2022

import (
	"fmt"
	"log"
	"strings"
)

type Day19 struct{}

const simMin = 24

type oreType string

type oreCollectorType string

type UnitBlueprint struct {
	Name string
	Cost map[oreType]int
}

type Blueprint struct {
	Id         int
	Blueprints map[oreCollectorType]UnitBlueprint
}

type Unit struct {
	Name           string
	Production     map[oreType]int
	BaseProduction int
}

type MineralCollectorSystem struct {
	Inventory   map[oreType]int
	Blueprint   Blueprint
	Units       map[oreCollectorType]int
	WaitingTime int
	WaitingUnit *oreCollectorType
}

func NewMineralCollectorSystem(bp Blueprint) MineralCollectorSystem {
	i := make(map[oreType]int, 0)
	u := make(map[oreCollectorType]int, 0)
	u[oreCollectorType("ore")] = 1
	m := MineralCollectorSystem{Inventory: i, Units: u, Blueprint: bp}
	return m
}

func (m *MineralCollectorSystem) collectResource(o oreType) {
	// waterproof obsidian-collecting robots collect geodes
	robotCount, pres := m.Units[oreCollectorType(o)]
	if !pres {
		return
	}
	m.Inventory[o] += robotCount
	log.Printf("%d %s-collecting robot collects %d %s; you now have %d %s.", robotCount, o, robotCount, o, m.Inventory[o], o)
}

func (m *MineralCollectorSystem) tryBeginProduceCollectingRobot(o oreCollectorType) bool {
	// if we have more of each resources than is needed for the blueprint cost of a
	// that robot
	for resourceType, resourceAmount := range m.Blueprint.Blueprints[o].Cost {
		if m.Inventory[resourceType] < resourceAmount {
			return false
		}
	}
	collector := oreCollectorType(o)
	m.WaitingUnit = &collector
	m.WaitingTime = 1
	log.Printf("Spent resources to start building a %s-collecting robot.", o)
	for resourceType, resourceAmount := range m.Blueprint.Blueprints[o].Cost {
		m.Inventory[resourceType] -= resourceAmount
	}
	return true
}

func (m *MineralCollectorSystem) finishProduceCollectingRobot() {
	if m.WaitingUnit == nil {
		return
	}
	collector := *m.WaitingUnit
	m.WaitingUnit = nil
	m.WaitingTime = 0
	m.Units[collector] = m.Units[collector] + 1
	log.Printf("The new %s-collecting robot is ready; you now have %d of them.", collector, m.Units[collector])
}

/*
func (m *MineralCollectorSystem) getCurrentUnitMinuteProduction(o oreType) {
	total := 0

	for resourceType, resourceCost := range m.Blueprint.Blueprints[o].Cost {
		total +
	}
} */

func (m *MineralCollectorSystem) beginProduceRobot() {
	done := m.tryBeginProduceCollectingRobot(oreCollectorType("geode"))
	if done {
		return
	}
	done = m.tryBeginProduceCollectingRobot(oreCollectorType("obsidian"))
	if done {
		return
	}
	if m.Blueprint.Blueprints[oreCollectorType("obsidian")].Cost[oreType("clay")]*m.Units[oreCollectorType("clay")] <
		m.Blueprint.Blueprints[oreCollectorType("obsidian")].Cost[oreType("clay")]*m.Units[oreCollectorType("obsidian")] {
		done = m.tryBeginProduceCollectingRobot(oreCollectorType("clay"))
		if done {
			return
		}
	}
	done = m.tryBeginProduceCollectingRobot(oreCollectorType("ore"))
	if done {
		return
	}
}

func (m *MineralCollectorSystem) simulateMinute() {
	m.beginProduceRobot()
	m.collectResource(oreType("geode"))
	m.collectResource(oreType("obsidian"))
	m.collectResource(oreType("clay"))
	m.collectResource(oreType("ore"))
	m.finishProduceCollectingRobot()
}

func (p Day19) PartA(lines []string) any {
	allBP := make([]Blueprint, 0)
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		bp := Blueprint{Id: i + 1}
		parts := strings.Split(line, ":")
		bp.Blueprints = make(map[oreCollectorType]UnitBlueprint, 0)
		//baseProduction := 1
		for _, unitStatement := range strings.Split(parts[1], ".") {
			if len(unitStatement) == 0 {
				continue
			}
			unitType := strings.Split(unitStatement, " costs ")
			var unitName string
			fmt.Sscanf(strings.TrimSpace(unitType[0]), "Each %s", &unitName)
			resourceCosts := strings.Split(unitType[1], " and ")
			totalCosts := make(map[oreType]int)
			for _, cost := range resourceCosts {
				var amount int
				var resourceType string
				fmt.Sscanf(cost, "%d %s.", &amount, &resourceType)
				totalCosts[oreType(resourceType)] = amount
			}

			bp.Blueprints[oreCollectorType(unitName)] = UnitBlueprint{Name: unitName, Cost: totalCosts}
		}
		allBP = append(allBP, bp)
	}

	mcs := NewMineralCollectorSystem(allBP[0])
	for i := 1; i <= simMin; i++ {
		log.Printf("== Minute %d == ", i)
		mcs.simulateMinute()
	}
	return mcs.Inventory["geode"]
}

func (p Day19) PartB(lines []string) any {
	return "implement_me"
}
