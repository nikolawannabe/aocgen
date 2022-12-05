package year2022

import (
	"fmt"
	"log"
)

type Day05 struct{}

func findCrateStackHeight(lines []string) int {
	for i, line := range lines {
		if len(line) == 0 {
			return i
		}
	}
	return 0
}

func getCrateStacks(crateLines []string) [][]rune {
	stackCount := len(crateLines[0])/4 + 1
	stackMaxHeight := len(crateLines)
	stacks := make([][]rune, stackCount)
	for crateLevel := stackMaxHeight - 1; crateLevel > -1; crateLevel-- {
		for charCounter := 1; charCounter < len(crateLines[crateLevel]); charCounter += 4 {
			crateStack := charCounter / 4
			if rune(crateLines[crateLevel][charCounter]) == ' ' {
				continue
			}
			stacks[crateStack] = append(stacks[crateStack], rune(crateLines[crateLevel][charCounter]))
		}
	}
	return stacks
}

type instruction struct {
	MoveHowMany  int
	FromStackNum int
	ToStackNum   int
}

func getStackInstructions(instructionLines []string) []instruction {
	instructions := make([]instruction, 0)
	for _, line := range instructionLines {
		if len(line) == 0 {
			continue
		}
		var howMany, fromStack, toStack int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &howMany, &fromStack, &toStack)
		if err != nil {
			log.Printf("err: %v", err)
		}
		instruction := instruction{MoveHowMany: howMany, FromStackNum: fromStack, ToStackNum: toStack}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func printStack(stack []rune) {
	stringOut := ""
	for _, c := range stack {
		stringOut = stringOut + string(c)
	}
	log.Printf("stack: `%s`", stringOut)

}

func printStacks(stacks [][]rune) {
	for _, stack := range stacks {
		printStack(stack)
	}
}

func reverseSlice(slice []rune) []rune {
	var rev []rune
	for _, n := range slice {
		rev = append([]rune{n}, rev...)
	}
	return rev
}
func runInstruction(stacks [][]rune, ins instruction, newCrane bool) [][]rune {
	tsl := len(stacks[ins.FromStackNum-1])
	stuffToMove := stacks[ins.FromStackNum-1][tsl-ins.MoveHowMany : tsl]
	if !newCrane {
		stuffToMove = reverseSlice(stuffToMove)
	}
	stacks[ins.ToStackNum-1] = append(stacks[ins.ToStackNum-1], stuffToMove...)
	stacks[ins.FromStackNum-1] = stacks[ins.FromStackNum-1][:tsl-ins.MoveHowMany]
	return stacks
}

func getStackTops(stacks [][]rune) string {
	outputStr := ""
	for _, stack := range stacks {
		outputStr = outputStr + string(stack[len(stack)-1])
	}
	return outputStr
}
func (p Day05) PartA(lines []string) any {
	stackHeight := findCrateStackHeight(lines) - 1
	log.Printf("stack height: %d", stackHeight)

	stacks := getCrateStacks(lines[0:stackHeight])
	printStacks(stacks)

	instructions := getStackInstructions(lines[stackHeight+2:])
	log.Printf("instructions: %v", instructions)

	for _, instruction := range instructions {
		stacks = runInstruction(stacks, instruction, false)
		log.Printf("instruction: %d", instruction)
		printStacks(stacks)

	}
	return getStackTops(stacks)
}

func (p Day05) PartB(lines []string) any {
	stackHeight := findCrateStackHeight(lines) - 1
	log.Printf("stack height: %d", stackHeight)

	stacks := getCrateStacks(lines[0:stackHeight])
	printStacks(stacks)

	instructions := getStackInstructions(lines[stackHeight+2:])
	log.Printf("instructions: %v", instructions)

	for _, instruction := range instructions {
		stacks = runInstruction(stacks, instruction, true)
		log.Printf("instruction: %d", instruction)
		printStacks(stacks)

	}
	return getStackTops(stacks)
}
