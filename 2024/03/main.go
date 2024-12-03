package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	util.PrintSolution(1, p1(input))
	util.PrintSolution(2, p2(input))
}

func p1(inputStr string) string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAll([]byte(inputStr), -1)

	total := 0
	for _, match := range matches {
		re = regexp.MustCompile(`\d+`)
		nums := re.FindAll(match, -1)
		asInts := covertToInts(nums)

		for i := 0; i < len(asInts); i += 2 {
			total += asInts[0] * asInts[i+1]
		}

	}

	return strconv.Itoa(total)
}

func p2(inputStr string) string {
	instructions := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don\'t\(\))`)

	matches := instructions.FindAll([]byte(inputStr), -1)

	operationsEnabled := true

	total := 0
	for _, match := range matches {
		asStr := string(match)[:3]
		switch asStr {
		case "do(":
			operationsEnabled = true
		case "don":
			operationsEnabled = false
		case "mul":
			if !operationsEnabled {
				continue
			}
			re := regexp.MustCompile(`\d+`)
			nums := re.FindAll(match, -1)
			asInts := covertToInts(nums)

			for i := 0; i < len(asInts); i += 2 {
				total += asInts[0] * asInts[i+1]
			}
		}

	}
	return strconv.Itoa(total)
}

func covertToInts(bytes [][]byte) []int {
	asInts := make([]int, 0)
	for _, num := range bytes {
		x, _ := strconv.Atoi(string(num))
		asInts = append(asInts, x)
	}

	return asInts
}
