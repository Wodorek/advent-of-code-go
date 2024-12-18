package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)

	util.PrintSolution(1, p1(inputArr))
}

func p1(inputArr []equation) string {
	totalValue := 0
	for _, equation := range inputArr {
		if equation.checkPossibleP1() {
			totalValue += equation.value
		}
	}

	return strconv.Itoa(totalValue)
}

func parseInput(inputStr string) []equation {
	equations := make([]equation, 0)

	split := strings.Split(inputStr, "\n")

	for _, line := range split {
		parsedEquation := equation{}

		splitLine := strings.Split(line, ": ")
		val, err := strconv.Atoi(splitLine[0])

		if err != nil {
			fmt.Println(err)
			return nil
		}

		splitOperators := strings.Split(splitLine[1], " ")
		for _, operator := range splitOperators {
			converted, err := strconv.Atoi(operator)

			if err != nil {
				fmt.Println(err)
				return nil
			}

			parsedEquation.operators = append(parsedEquation.operators, converted)
		}

		parsedEquation.value = val
		equations = append(equations, parsedEquation)
	}

	return equations
}

type equation struct {
	value     int
	operators []int
}

func (e *equation) checkPossibleP1() bool {
	possibleValues := make([]int, 0)
	possibleValues = append(possibleValues, e.operators[0])

	for i := 1; i < len(e.operators); i++ {
		newValues := make([]int, 0)

		for _, val := range possibleValues {
			newValues = append(newValues, val*e.operators[i])
			newValues = append(newValues, val+e.operators[i])
		}
		possibleValues = newValues
	}

	return slices.Contains(possibleValues, e.value)
}
