package main

import (
	_ "embed"
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
	return ""
}

func p2() string {
	return ""
}

func prepareInput() [][]int {
	split := strings.Split(input, "\n")

	result := make([][]int, 0)

	for _, val := range split {

	}

	return result
}
