package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {

	util.PrintSolution(1, p1())
	util.PrintSolution(2, p2())
}

func p1() string {

	moves := map[string][]int{
		"^": {0, 1},
		"v": {0, -1},
		">": {1, 0},
		"<": {-1, 0},
	}

	inputArr := prepareInput()
	locations := make(map[string]int)
	locations["0,0"] = 1
	currX := 0
	currY := 0

	for _, move := range inputArr {
		coords := moves[move]
		x, y := coords[0], coords[1]
		currY += y
		currX += x
		moveTo(locations, currX, currY)
	}

	return fmt.Sprintf("%v", len(locations))
}

func moveTo(locations map[string]int, x, y int) {

	key := fmt.Sprintf("%d,%d", x, y)
	if _, ok := locations[key]; ok {
		locations[key]++
	} else {
		locations[key] = 0
	}
}

func p2() string {
	return ""
}

func prepareInput() []string {
	return strings.Split(input, "")

}
