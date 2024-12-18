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

var initial []int

func main() {
	inputArr := parseInput(input)
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))

}

func p1(inputArr [][]string) string {
	guard := guard{exited: false}
	guard.facing.x = 0
	guard.facing.y = -1

	for y, row := range inputArr {
		for x, char := range row {
			if char == "^" {
				guard.setPosition(x, y)
				initial = []int{x, y}
				fmt.Println(x, y)
			}
		}
	}

	for {
		if guard.exited {
			break
		}
		guard.move(inputArr)
	}

	totalVisited := 1

	for _, row := range inputArr {
		for _, char := range row {
			if char == "X" || char == "^" {
				totalVisited += 1
			}
		}
	}

	return strconv.Itoa(totalVisited)
}

func p2(inputArr [][]string) string {

	for y, line := range inputArr {
		for x, char := range line {
			if char == "X" {

				if x > 0 && inputArr[y][x-1] != "#" {
					inputArr[y][x-1] = "X"
				}
				if x < len(inputArr[y])-1 && inputArr[y][x+1] != "#" {
					inputArr[y][x+1] = "X"
				}
				if y > 0 && inputArr[y-1][x] != "#" {
					inputArr[y-1][x] = "X"
				}
				if y < len(inputArr)-1 && inputArr[y+1][x] != "#" {
					inputArr[y+1][x] = "X"
				}
			}
		}
	}

	totalLoops := 0

	guard := guard{exited: false}
	guard.facing.x = 0
	guard.facing.y = -1
	guard.initial.x = initial[0]
	guard.initial.y = initial[1]

	for y, row := range inputArr {
		for x, char := range row {
			if char == "^" {
				guard.setInitialPosition(x, y)
			}
		}
	}

	for y, line := range inputArr {
		for x, char := range line {
			if char == "X" {
				guard.reset()
				inputArr[y][x] = "#"
				for i := 0; i < 10000; i++ {
					guard.move(inputArr)

				}

				if !guard.exited {
					totalLoops++
				}
				inputArr[y][x] = "X"
				guard.reset()
			}
		}
	}

	return strconv.Itoa(totalLoops)
}

type guard struct {
	position struct {
		x int
		y int
	}
	facing struct {
		x int
		y int
	}
	initial struct {
		x int
		y int
	}
	exited bool
}

func (g *guard) rotate() {
	positions := [][2]int{
		{0, -1}, //N
		{1, 0},  //E
		{0, 1},  //S
		{-1, 0}, //W
	}

	currIdx := -1

	for i := 0; i < 3; i++ {
		if positions[i][0] == g.facing.x && positions[i][1] == g.facing.y {
			currIdx = i
		}
	}
	g.facing.x, g.facing.y = positions[currIdx+1%4][0], positions[currIdx+1%4][1]

}

func (g *guard) setPosition(x, y int) {
	g.position.x = x
	g.position.y = y
}

func (g *guard) setInitialPosition(x, y int) {
	g.initial.x = x
	g.initial.y = y
}

func (g *guard) reset() {
	g.position.x = g.initial.x
	g.position.y = g.initial.y
	g.exited = false
	g.facing.x = 0
	g.facing.y = -1
}

func (g *guard) move(matrix [][]string) {
	if g.position.x+g.facing.x > len(matrix[0])-1 ||
		g.position.x+g.facing.x < 0 ||
		g.position.y+g.facing.y > len(matrix)-1 ||
		g.position.y+g.facing.y < 0 {
		g.exited = true
		return

	}

	if matrix[g.position.y+g.facing.y][g.position.x+g.facing.x] == "#" {
		g.rotate()
		return
	}

	matrix[g.position.y][g.position.x] = "X"
	g.position.x += g.facing.x
	g.position.y += g.facing.y
}

func parseInput(inputString string) [][]string {
	parsed := make([][]string, 0)
	lines := strings.Split(inputString, "\n")

	for _, line := range lines {
		splitLine := strings.Split(line, "")
		parsed = append(parsed, splitLine)
	}
	return parsed
}
