package year2022

import (
	"fmt"
	"log"
	"strings"
)

type Day21 struct{}

type expression struct {
	operator rune
	operand1 string
	operand2 string
	value    *int
}

type statement struct {
	symbol     string
	expression expression
	value      *int
}

func parse(lines []string) []statement {
	statements := make([]statement, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line := strings.ReplaceAll(line, ":", " : ")
		var statementSymbol string
		var value int
		_, err := fmt.Sscanf(line, "%s : %d", &statementSymbol, &value)
		if err != nil {
			var operand1 string
			var operand2 string
			var operator rune
			_, err := fmt.Sscanf(line, "%s : %s %c %s", &statementSymbol, &operand1, &operator, &operand2)
			if err != nil {
				log.Printf("couldn't read line %s: %v", line, err)
				continue
			}
			statement := statement{symbol: statementSymbol, expression: expression{operator: operator, operand1: operand1, operand2: operand2}}
			statements = append(statements, statement)
			continue
		}
		statement := statement{symbol: statementSymbol, value: &value}
		statements = append(statements, statement)
	}
	return statements
}

type symbols map[string]statement

func (s *symbols) eval(sme statement) *int {
	if sme.value != nil {
		return sme.value
	}

	exp := sme.expression
	op1 := map[string]statement(*s)[exp.operand1]
	op2 := map[string]statement(*s)[exp.operand2]
	operator := exp.operator

	ans1 := s.eval(op1)
	ans2 := s.eval(op2)
	if ans1 == nil {
		return nil
	}
	if ans2 == nil {
		return nil
	}
	switch operator {
	case '/':
		value := *ans1 / *ans2
		log.Printf("%d = %s / %s", value, op1.symbol, op2.symbol)
		sme.value = &value
	case '+':
		value := *ans1 + *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
	case '-':
		value := *ans1 - *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
	case '*':
		value := *ans1 * *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
	default:
		log.Printf("%s, %s", op1.symbol, op2.symbol)
	}
	return sme.value
}

func (s *symbols) evalCmp(sme statement, eq int) (*int, *int) {
	if sme.value != nil {
		return sme.value, nil
	}

	exp := sme.expression
	op1 := map[string]statement(*s)[exp.operand1]
	op2 := map[string]statement(*s)[exp.operand2]
	operator := exp.operator

	switch operator {
	case '/':
		ans1, solution := s.evalCmp(op1, eq)
		ans2, solution := s.evalCmp(op2, eq)
		if ans1 == nil {
			*solution = eq * *ans2
			return nil, solution
		}
		if ans2 == nil {
			*solution = eq * *ans1
			return nil, solution
		}
		value := *ans1 / *ans2
		log.Printf("%d = %s / %s", value, op1.symbol, op2.symbol)
		sme.value = &value
		return sme.value, solution
	case '+':
		ans1, solution := s.evalCmp(op1, eq)
		ans2, solution := s.evalCmp(op2, eq)
		if ans1 == nil {
			*solution = eq - *ans2
			return nil, solution
		}
		if ans2 == nil {
			*solution = eq - *ans1
			return nil, solution
		}
		value := *ans1 + *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
		return sme.value, solution
	case '-':
		ans1, solution := s.evalCmp(op1, eq)
		ans2, solution := s.evalCmp(op2, eq)
		if ans1 == nil {
			*solution = eq + *ans2
			return nil, solution
		}
		if ans2 == nil {
			*solution = eq + *ans1
			return nil, solution
		}
		value := *ans1 - *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
		return sme.value, solution
	case '*':
		ans1, solution := s.evalCmp(op1, eq)
		ans2, solution := s.evalCmp(op2, eq)
		if ans1 == nil {
			*solution = eq / *ans2
			return nil, solution
		}
		if ans2 == nil {
			*solution = eq / *ans1
			return nil, solution
		}
		value := *ans1 * *ans2
		log.Printf("%d = %s + %s", value, op1.symbol, op2.symbol)
		sme.value = &value
		return sme.value, solution
	default:
		log.Printf("%s, %s", op1.symbol, op2.symbol)
	}
	return sme.value, nil
}

func evalP1Root(rootNum int, lines []string) int {
	statements := parse(lines)
	assignments := make(map[string]statement, 0)

	for _, statement := range statements {
		assignments[statement.symbol] = statement
	}

	s := symbols(assignments)
	ans := s.eval(statements[rootNum])
	return *ans
}

func evalP2Root(rootNum int, lines []string) int {
	statements := parse(lines)
	assignments := make(map[string]statement, 0)

	for _, statement := range statements {
		if statement.symbol == "root" {
			statement.expression = expression{operator: '='}
		}
		if statement.symbol == "humn" {
			statement.value = nil
		}
		assignments[statement.symbol] = statement
	}

	s := symbols(assignments)

	sme := statements[rootNum]
	exp := sme.expression
	op1 := map[string]statement(s)[exp.operand1]
	op2 := map[string]statement(s)[exp.operand2]

	//var ans *int
	s["humn"] = statement{symbol: "humn", value: nil}
	target := s.eval(op2)
	//log.Printf("seeking %d", target)
	_, solution := s.evalCmp(op1, *target)

	return *solution
}

func (p Day21) PartA(lines []string) any {
	return evalP1Root(65, lines)
}

func (p Day21) PartB(lines []string) any {
	return evalP2Root(65, lines)
}
