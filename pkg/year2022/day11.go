package year2022

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day11 struct{}

type Item struct {
	Id                int
	InitialWorryLevel int
	CurrentWorryLevel int
}

type Monkey struct {
	Id          int
	Items       []Item
	Operation   rune
	Operand     string
	ModAmount   int
	TrueMonkey  int
	FalseMonkey int
	Inspections int
}

func doOperation(oldWorryLevel int, operation rune, operand string) int {
	num := 0
	if operand == "old" {
		num = oldWorryLevel
	} else {
		num, _ = strconv.Atoi(operand)
	}
	switch operation {
	case '*':
		return oldWorryLevel * num
	case '+':
		return oldWorryLevel + num
	}
	log.Printf("operation unknown %c", operation)
	return 0
}
func doRound(monkeys map[int]Monkey, moreWorried bool, lcm int) map[int]Monkey {
	monkeyCount := 0
	for range monkeys {
		monkeyCount++
	}

	for monkeyI := 0; monkeyI < monkeyCount; monkeyI++ {
		curMonkey := monkeys[monkeyI]
		for _, item := range curMonkey.Items {
			curMonkey.Inspections++
			//log.Printf("\tMonkey inspects an item with a worry level of %s", item.CurrentWorryLevel.String())
			n := doOperation(item.CurrentWorryLevel, curMonkey.Operation, curMonkey.Operand)
			item.CurrentWorryLevel = n
			//log.Printf("\t\tWorry level is %c by %s to %s", curMonkey.Operation, curMonkey.Operand, n.String())

			if !moreWorried {
				item.CurrentWorryLevel = item.CurrentWorryLevel / 3
			} else {
				item.CurrentWorryLevel = item.CurrentWorryLevel % lcm
			}
			//log.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %s", item.CurrentWorryLevel.String())

			m := item.CurrentWorryLevel % curMonkey.ModAmount
			monkeyToThrowTo := monkeys[curMonkey.FalseMonkey]
			if m == 0 {
				monkeyToThrowTo = monkeys[curMonkey.TrueMonkey]
				//log.Printf("\t\tCurrent worry level is divisible by %s", curMonkey.ModAmount.String())
				//log.Printf("\t\tItem with worry level %s is thrown to monkey %d.", item.CurrentWorryLevel.String(), curMonkey.TrueMonkey)
			} else {
				//log.Printf("\t\tCurrent worry level is not divisible by %s", curMonkey.ModAmount.String())
				//log.Printf("\t\tItem with worry level %s is thrown to monkey %d.", item.CurrentWorryLevel.String(), curMonkey.FalseMonkey)
			}
			monkeyToThrowTo.Items = append(monkeyToThrowTo.Items, item)
			monkeys[monkeyToThrowTo.Id] = monkeyToThrowTo
			curMonkey.Items = curMonkey.Items[0 : len(curMonkey.Items)-1]
			monkeys[monkeyI] = curMonkey
		}
		monkeys[monkeyI] = curMonkey
	}
	/*for _, monkey := range monkeys {
		itemIds := make([]string, 0)
		for _, item := range monkey.Items {
			itemIds = append(itemIds, item.CurrentWorryLevel.String())
		}
		//log.Printf("Monkey %d: %s", monkey.Id, strings.Join(itemIds, ","))
	}*/
	return monkeys
}

func getMonkeys(lines []string) map[int]Monkey {
	monkeys := make(map[int]Monkey, 0)

	monkey := Monkey{}
	itemId := 1
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			monkey = Monkey{}
			continue
		}
		var id int
		fmt.Sscanf(lines[i], "Monkey %d:", &id)
		monkey.Id = id
		i++
		items := make([]Item, 0)
		startingParts := strings.Split(strings.TrimSpace(lines[i]), ": ")
		startingWorryItems := strings.Split(startingParts[1], ",")
		for _, worryLevel := range startingWorryItems {
			level, _ := strconv.Atoi(strings.TrimSpace(worryLevel))
			item := Item{Id: itemId, InitialWorryLevel: level, CurrentWorryLevel: level}
			itemId++
			items = append(items, item)
		}
		monkey.Items = items
		i++
		var operation rune
		var operand string
		_, err := fmt.Sscanf(strings.TrimSpace(lines[i]), "Operation: new = old %c %s", &operation, &operand)
		if err != nil {
			log.Printf("err: %v", err)
		}
		monkey.Operand = operand
		monkey.Operation = operation
		i++
		var modAmount int
		fmt.Sscanf(strings.TrimSpace(lines[i]), "Test: divisible by %d", &modAmount)
		monkey.ModAmount = modAmount
		i++
		var trueMonkey int
		fmt.Sscanf(strings.TrimSpace(lines[i]), "If true: throw to monkey %d", &trueMonkey)
		monkey.TrueMonkey = trueMonkey
		i++
		var falseMonkey int
		fmt.Sscanf(strings.TrimSpace(lines[i]), "If false: throw to monkey %d", &falseMonkey)
		monkey.FalseMonkey = falseMonkey

		monkeys[id] = monkey
	}
	return monkeys
}

func getMonkeyLCM(monkeys map[int]Monkey) int {
	m := make([]int, 0)
	for _, monkey := range monkeys {
		m = append(m, monkey.ModAmount)
	}

	return LCM(m[0], m[1], m[2:])
}

func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[2:])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (p Day11) PartA(lines []string) any {
	rounds := 20

	monkeys := getMonkeys(lines)
	log.Printf("%#v", monkeys)

	for i := 0; i < rounds; i++ {
		monkeys = doRound(monkeys, false, 3)
	}

	inspectionCounts := make([]int, 0)
	for _, monkey := range monkeys {
		log.Printf("Monkey %d inspected items %d times.", monkey.Id, monkey.Inspections)
		inspectionCounts = append(inspectionCounts, monkey.Inspections)
	}

	return 0
}

func (p Day11) PartB(lines []string) any {
	rounds := 10000

	monkeys := getMonkeys(lines)
	lcm := getMonkeyLCM(monkeys)
	log.Printf("lcm: %d", lcm)

	for i := 0; i < rounds; i++ {
		monkeys = doRound(monkeys, true, lcm)
		if i%1000 == 0 {
			log.Printf("finished round %d", i)
		}
	}

	inspectionCounts := make([]int, 0)
	for _, monkey := range monkeys {
		log.Printf("Monkey %d inspected items %d times.", monkey.Id, monkey.Inspections)
		inspectionCounts = append(inspectionCounts, monkey.Inspections)
	}

	return 0
}
