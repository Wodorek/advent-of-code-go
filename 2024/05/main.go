package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := prepareInput(input)
	util.PrintSolution(1, p1(inputArr))
}

func p1(inputArr rules) string {
	return ""
}

func prepareInput(inputStr string) rules {
	rules := rules{ordering: [][]int{}, updates: [][]int{}}

	split := strings.Split(inputStr, "\n\n")

	for _, order := range strings.Split(split[0], "\n") {
		newPair := make([]int, 0)
		splitOrder := strings.Split(order, "|")
		left, _ := strconv.Atoi(splitOrder[0])
		right, _ := strconv.Atoi(splitOrder[1])

		newPair = append(newPair, left, right)
		rules.ordering = append(rules.ordering, newPair)
	}

	for _, update := range strings.Split(split[1], "\n") {
		splitUpdate := strings.Split(update, ",")
		newUpdate := make([]int, 0)

		for _, num := range splitUpdate {
			asInt, _ := strconv.Atoi(num)
			newUpdate = append(newUpdate, asInt)
		}
		rules.updates = append(rules.updates, newUpdate)
	}

	return rules
}

type rules struct {
	ordering [][]int
	updates  [][]int
}
