package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)

	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(inputArr []int) string {
	buffer := make([]int, 0)

	currId := 0

	for i, val := range inputArr {
		if i%2 == 0 {
			buffer = append(buffer, slices.Repeat([]int{currId}, val)...)
			currId++
		} else {
			buffer = append(buffer, slices.Repeat([]int{-1}, val)...)
		}
	}

	left := 0
	right := len(buffer) - 1

	for {
		if left > right {
			break
		}

		if buffer[left] == -1 {
			if buffer[right] != -1 {
				buffer[left] = buffer[right]
				buffer[right] = -1
				right--
				left++
			} else {
				right--
			}
		} else {
			left++
		}
	}

	total := 0

	idx := 0
	for {
		if buffer[idx] == -1 {
			break
		}
		total += idx * buffer[idx]
		idx++
	}

	return strconv.Itoa(total)
}

func p2(inputArr []int) string {
	return ""
}

func parseInput(inputString string) []int {
	split := strings.Split(inputString, "")
	nums := make([]int, 0)
	for _, char := range split {
		converted, err := strconv.Atoi(char)
		if err != nil {
			return []int{}
		}
		nums = append(nums, converted)
	}

	return nums
}
