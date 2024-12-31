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

type trail struct {
	value     int
	x         int
	y         int
	arrivesTo map[string]int
}

func main() {
	inputArr := parseInput(input)
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(inputArr [][]int) string {
	trailHeads := make([]trail, 0)

	for y := 0; y < len(inputArr); y++ {
		for x := 0; x < len(inputArr[0]); x++ {
			if inputArr[y][x] == 0 {
				trailHeads = append(trailHeads, trail{0, x, y, make(map[string]int)})
			}
		}
	}

	total := 0
	for _, trailHead := range trailHeads {
		traverseTrailp1(inputArr, &trailHead, trailHead, 1)
	}

	for _, t := range trailHeads {
		total += len(t.arrivesTo)
	}

	return strconv.Itoa(total)
}

func traverseTrailp1(inputArr [][]int, currHead *trail, currNode trail, nextVal int) {
	neigbors := getNeighborsCardinal(inputArr, currNode.x, currNode.y)
	matchingNeighbors := make([]trail, 0)
	for _, neigh := range neigbors {
		if neigh.value == nextVal {
			matchingNeighbors = append(matchingNeighbors, neigh)
		}
	}

	if nextVal == 9 {
		for _, match := range matchingNeighbors {
			key := fmt.Sprintf("%d,%d", match.x, match.y)
			(*currHead).arrivesTo[key] = 1
		}
	} else {
		for _, match := range matchingNeighbors {
			traverseTrailp1(inputArr, currHead, match, nextVal+1)
		}
	}

}

func traverseTrailp2(inputArr [][]int, currHead *trail, currNode trail, nextVal int) {
	neigbors := getNeighborsCardinal(inputArr, currNode.x, currNode.y)
	matchingNeighbors := make([]trail, 0)
	for _, neigh := range neigbors {
		if neigh.value == nextVal {
			matchingNeighbors = append(matchingNeighbors, neigh)
		}
	}

	if nextVal == 9 {
		for _, match := range matchingNeighbors {
			key := fmt.Sprintf("%d,%d", match.x, match.y)
			(*currHead).arrivesTo[key] = 1
		}
	} else {
		for _, match := range matchingNeighbors {
			traverseTrailp2(inputArr, currHead, match, nextVal+1)
		}
	}

}

func p2(inputArr [][]int) string {
	trailHeads := make([]trail, 0)

	for y := 0; y < len(inputArr); y++ {
		for x := 0; x < len(inputArr[0]); x++ {
			if inputArr[y][x] == 0 {
				trailHeads = append(trailHeads, trail{0, x, y, make(map[string]int)})
			}
		}
	}

	total := 0
	for _, trailHead := range trailHeads {
		traverseTrailp2(inputArr, &trailHead, trailHead, 1)
	}

	for _, t := range trailHeads {
		total += len(t.arrivesTo)
	}

	return strconv.Itoa(total)
}

func parseInput(inputString string) [][]int {
	split := strings.Split(inputString, "\n")
	parsed := make([][]int, 0)

	for _, line := range split {
		newLine := make([]int, 0)
		splitLine := strings.Split(line, "")

		for _, num := range splitLine {
			asInt, err := strconv.Atoi(num)

			if err != nil {
				fmt.Println(err)
				return [][]int{}
			}
			newLine = append(newLine, asInt)
		}
		parsed = append(parsed, newLine)
	}
	return parsed
}

func getNeighborsCardinal(arr [][]int, x, y int) []trail {
	neighbors := []trail{
		{-1, x - 1, y, make(map[string]int)},
		{-1, x, y - 1, make(map[string]int)},
		{-1, x + 1, y, make(map[string]int)},
		{-1, x, y + 1, make(map[string]int)}}

	//you are outside of the array, you don't exits, so you don't have neighbors
	if x > len(arr[0])-1 || y > len(arr)-1 || x < 0 || y < 0 {
		return neighbors
	}

	if y > 0 {
		neighbors[1].value = arr[y-1][x]
	}

	if x > 0 {
		neighbors[0].value = arr[y][x-1]
	}

	if x < len(arr[0])-1 {
		neighbors[2].value = arr[y][x+1]
	}

	if y < len(arr)-1 {
		neighbors[3].value = arr[y+1][x]

	}

	return neighbors
}
