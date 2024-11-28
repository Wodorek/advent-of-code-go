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
	totalNice := 0

	for _, str := range arr {
		if testStringNice2(str) {
			totalNice++
		}
	}

	return strconv.Itoa(totalNice)
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

func testStringNice2(str string) bool {
	//screw regexp
	hasSandwich := false
	hasDouble := false

	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			hasSandwich = true
		}
	}

	for i := 0; i < len(str)-1; i++ {
		if strings.Contains(str[i+2:], str[i:i+2]) {
			hasDouble = true
		}
	}

	return hasDouble && hasSandwich

}
