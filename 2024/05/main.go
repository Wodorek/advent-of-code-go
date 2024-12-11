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
	inputArr := prepareInput(input)
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(inputArr rules) string {

	parsedRules := make(map[int][]int)

	for _, order := range inputArr.ordering {
		key := order[0]
		beAfter := order[1]

		if _, ok := parsedRules[key]; !ok {
			parsedRules[key] = []int{beAfter}
		} else {
			parsedRules[key] = append(parsedRules[key], beAfter)
		}
	}
	totalMiddles := 0

	for _, update := range inputArr.updates {
		isCorrect := true
		for i := 0; i < len(update)-1; i++ {
			left := update[i]
			right := update[i+1]

			if rules, ok := parsedRules[left]; ok {
				if !slices.Contains(rules, right) {
					isCorrect = false
				}
			} else {
				isCorrect = false
			}
		}

		if isCorrect {
			totalMiddles += update[int(float64(len(update)/2))]
		}
	}

	return strconv.Itoa(totalMiddles)
}

func p2(inputArr rules) string {

	parsedRules := make(map[int][]int)

	for _, order := range inputArr.ordering {
		key := order[0]
		beAfter := order[1]

		if _, ok := parsedRules[key]; !ok {
			parsedRules[key] = []int{beAfter}
		} else {
			parsedRules[key] = append(parsedRules[key], beAfter)
		}
	}

	// orderedArr := make([]int, 300)

	for k, r := range parsedRules {
		fmt.Println(k, ":", r)
	}

	return ""
}

// func fixUpdate(update []int, rules map[int][]int) []int {

// 	return []int{}
// }

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
