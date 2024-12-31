package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)
	fmt.Println(inputArr)
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(inputArr []int) string {

	stones := make([]int, len(inputArr))
	copy(stones, inputArr)

	for i := 0; i < 25; i++ {
		newStones := make([]int, 0)
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
				asStr := strconv.Itoa(stone)
				left, right := asStr[0:len(asStr)/2], asStr[len(asStr)/2:]
				leftInt, err := strconv.Atoi(left)
				if err != nil {
					fmt.Println(err)
				}
				rightInt, err := strconv.Atoi(right)
				if err != nil {
					fmt.Println(err)
				}
				newStones = append(newStones, leftInt, rightInt)
			} else {
				newStones = append(newStones, (stone * 2024))
			}
		}

		stones = newStones

	}

	return fmt.Sprintf("%d", len(stones))
}

func p2(inputArr []int) string {
	return ""
}

func parseInput(inputString string) []int {
	split := strings.Split(inputString, " ")

	parsed := make([]int, 0)

	for _, val := range split {
		asInt, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println(err)
			return nil
		}
		parsed = append(parsed, asInt)
	}

	return parsed
}
