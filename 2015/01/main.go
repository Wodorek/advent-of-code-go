package main

import (
	"fmt"

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
	currLvl := 0

	split := strings.Split(input, "")

	for _, move := range split {
		if move == "(" {
			currLvl += 1
		} else {
			currLvl -= 1
		}
	}

	return fmt.Sprint(currLvl)
}

func p2() string {
	currLvl := 0

	split := strings.Split(input, "")

	for i, move := range split {
		if move == "(" {
			currLvl += 1
		} else {
			currLvl -= 1
		}

		if currLvl == -1 {

			return fmt.Sprint(i + 1)
		}
	}
	return "0"
}
