package main

import (
	_ "embed"
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

		if !isCorrect {
			sortedUpdate := sortUpdate(update, parsedRules)
			totalMiddles += sortedUpdate[len(sortedUpdate)/2]
		}
	}

	return strconv.Itoa(totalMiddles)

}

func sortUpdate(update []int, rules map[int][]int) []int {

	for i := 0; i < len(update)-1; i++ {
		for j := 0; j < len(update)-i-1; j++ {
			rule := rules[update[j]]
			if !slices.Contains(rule, update[j+1]) {
				update[j], update[j+1] = update[j+1], update[j]
			}
		}
	}

	return update
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
