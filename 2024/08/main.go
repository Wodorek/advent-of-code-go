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

var maxY int
var maxX int

func main() {
	inputArr := parseInput(input)

	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(inputArr map[string][][]int) string {

	allNodes := make(map[string]string)

	for _, v := range inputArr {
		antinodes := makeAntinodes(v)
		for _, node := range antinodes {
			if checkNodeInRange(node, maxX, maxY) {
				key := fmt.Sprintf("%d,%d", node[0], node[1])
				allNodes[key] = "x"
			}
		}
	}

	return strconv.Itoa(len(allNodes))
}

func p2(inputArr map[string][][]int) string {
	return ""
}

func makeAntinode(a1, a2 []int) [][]int {
	dy := a1[0] - a2[0]
	dx := a1[1] - a2[1]

	return [][]int{{a1[0] + dy, a1[1] + dx}, {a2[0] - dy, a2[1] - dx}}
}

func makeAntinodes(antennas [][]int) [][]int {
	nodes := make([][]int, 0)

	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {

			newNodes := makeAntinode(antennas[i], antennas[j])
			nodes = append(nodes, newNodes...)
		}
	}

	return nodes
}

func checkNodeInRange(node []int, x, y int) bool {

	return node[0] >= 0 && node[1] >= 0 && node[0] < x && node[1] < y

}

func parseInput(inputString string) map[string][][]int {

	antennas := make(map[string][][]int)

	lines := strings.Split(inputString, "\n")
	maxY = len(lines)
	maxX = len(lines[0])
	for y, line := range lines {
		splitline := strings.Split(line, "")
		for x, char := range splitline {
			if char != "." {
				if _, ok := antennas[char]; !ok {
					antennas[char] = make([][]int, 0)
				}
				antennas[char] = append(antennas[char], []int{x, y})
			}
		}

	}

	return antennas
}
