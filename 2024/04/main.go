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
	inputArr := prepareInput(input)

	directions := [][][2]int{
		{{-1, -1}, {-1, 0}, {-1, 1}},
		{{0, -1}, {0, 0}, {0, 1}},
		{{1, -1}, {1, 0}, {1, 1}}}

	util.PrintSolution(1, p1(inputArr, directions))
	util.PrintSolution(2, p2(inputArr, directions))

}

func prepareInput(inputStr string) [][]string {
	split := strings.Split(inputStr, "\n")
	matrix := make([][]string, 0)

	for _, line := range split {
		matrix = append(matrix, strings.Split(line, ""))

	}
	return matrix
}

func p1(inputArr [][]string, directions [][][2]int) string {

	totalStrings := 0

	for y := range inputArr {
		for x := range inputArr[y] {
			if inputArr[y][x] == "X" {
				neighs := util.GetNeighborsDiagonal(inputArr, x, y, nil)
				for nx := range neighs {
					for ny := range neighs[nx] {
						if neighs[nx][ny] == "M" {
							isFullXMAS := checkNextLetters(inputArr, x+directions[ny][nx][0], y+directions[ny][nx][1], &directions[ny][nx])

							if isFullXMAS {
								totalStrings++
							}
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(totalStrings)
}

func p2(inputArr [][]string, directions [][][2]int) string {
	totalStrings := 0

	for y := range inputArr {
		for x := range inputArr[y] {
			if inputArr[y][x] == "A" {
				str1 := "A"
				str2 := "A"

				neighs := util.GetNeighborsDiagonal(inputArr, x, y, nil)
				str1 = neighs[0][0] + str1 + neighs[2][2]
				str2 = neighs[2][0] + str2 + neighs[0][2]

				fmt.Println(str1, str2)

				if (str1 == "MAS" || str1 == "SAM") && (str2 == "MAS" || str2 == "SAM") {
					totalStrings++
				}
			}
		}
	}

	return strconv.Itoa(totalStrings)
}

func checkNextLetters(inputArr [][]string, x, y int, direction *[2]int) bool {
	finalStr := ""
	for i := 0; i < 2; i++ {
		nextLetter := util.GetNeighborsDiagonal(inputArr, x, y, direction)
		finalStr += nextLetter[1][1]
		x += direction[0]
		y += direction[1]
	}
	return finalStr == "AS"
}
