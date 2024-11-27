package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := prepareInput()

	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(arr []string) string {
	totalNice := 0

	for _, str := range arr {
		if testStringNice1(str) {
			totalNice++
		}
	}

	return strconv.Itoa(totalNice)
}

func p2(arr []string) string {
	return ""
}

func prepareInput() []string {
	return strings.Split(input, "\n")
}

func testStringNice1(str string) bool {
	vowels := 0
	hasDouble := false
	notForbidden := false

	split := strings.Split(str, "")

	for i := range split {
		switch split[i] {
		case "a", "e", "i", "o", "u":
			vowels++
		}

		if i < len(split)-1 {
			if split[i] == split[i+1] {
				hasDouble = true
			}
			switch split[i] + split[i+1] {
			case "ab", "cd", "pq", "xy":
				notForbidden = true
			}
		}
	}

	return vowels >= 3 && hasDouble && !notForbidden
}
