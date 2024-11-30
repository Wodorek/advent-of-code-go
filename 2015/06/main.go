package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type operation struct {
	what string
	x    struct {
		from int
		to   int
	}
	y struct {
		from int
		to   int
	}
}

func parseInputStr(str string) operation {
	var parsed operation

	split := strings.Split(str, " ")
	mod := 0
	if split[0] == "turn" {
		mod++
	}

	parsed.what = split[0+mod]
	startingPoints := strings.Split(split[1+mod], ",")
	xfrom, _ := strconv.Atoi(startingPoints[0])
	yfrom, _ := strconv.Atoi(startingPoints[1])
	parsed.x.from = xfrom
	parsed.y.from = yfrom
	endingPoints := strings.Split(split[3+mod], ",")

	xto, _ := strconv.Atoi(endingPoints[0])
	yto, _ := strconv.Atoi(endingPoints[1])
	parsed.x.to = xto
	parsed.y.to = yto

	return parsed
}

func main() {
	operations := []operation{}

	for _, operation := range strings.Split(input, "\n") {
		operations = append(operations, parseInputStr(operation))
	}

	util.PrintSolution(1, p1(operations))
	util.PrintSolution(2, p2(operations))
}

func p1(ops []operation) string {
	matrix := [1000][1000]int{}
	for _, op := range ops {
		for y := op.y.from; y <= op.y.to; y++ {
			for x := op.x.from; x <= op.x.to; x++ {
				switch op.what {
				case "on":
					matrix[y][x] = 1
				case "off":
					matrix[y][x] = 0
				case "toggle":
					if matrix[y][x] == 0 {
						matrix[y][x] = 1
					} else {
						matrix[y][x] = 0
					}
				}
			}
		}
	}

	totalBrightness := 0

	for i := range matrix {
		for _, y := range matrix[i] {
			totalBrightness += y
		}
	}
	return strconv.Itoa(totalBrightness)
}

func p2(ops []operation) string {
	matrix := [1000][1000]int{}
	for _, op := range ops {
		for y := op.y.from; y <= op.y.to; y++ {
			for x := op.x.from; x <= op.x.to; x++ {
				switch op.what {
				case "on":
					matrix[y][x]++
				case "off":
					if matrix[y][x] > 0 {
						matrix[y][x]--
					}

				case "toggle":
					if matrix[y][x] == 0 {
						matrix[y][x] = 1
					} else {
						matrix[y][x] = 0
					}
				}
			}
		}
	}

	totalBrightness := 0

	for i := range matrix {
		for _, y := range matrix[i] {
			totalBrightness += y
		}
	}
	return strconv.Itoa(totalBrightness)
}
