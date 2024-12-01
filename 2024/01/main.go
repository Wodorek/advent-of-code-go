package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func parseInput(inputStr string) ([][]int, error) {

	lists := make([][]int, 2)

	split := strings.Split(inputStr, "\n")
	for _, pair := range split {
		splitPair := strings.Split(pair, "   ")
		leftNum, err := strconv.Atoi(splitPair[0])
		if err != nil {
			return [][]int{}, err
		}
		rightNum, err := strconv.Atoi(splitPair[1])
		if err != nil {
			return [][]int{}, err
		}

		lists[0] = append(lists[0], leftNum)
		lists[1] = append(lists[1], rightNum)
	}

	return lists, nil
}

func p1(inputArr [][]int) string {
	fmt.Println(inputArr)
	sort.Slice(inputArr[0], func(a, b int) bool {
		return inputArr[0][a] < inputArr[0][b]
	})
	sort.Slice(inputArr[1], func(a, b int) bool {
		return inputArr[1][a] < inputArr[1][b]
	})

	totalDistance := 0.0

	for i, _ := range inputArr[0] {
		diff := math.Abs(float64(inputArr[0][i]) - float64(inputArr[1][i]))

		totalDistance += diff
	}
	return strconv.Itoa(int(totalDistance))
}
