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

	type diskBlock struct {
		size           int
		value          int
		positionOnDisk int
		spacesLeft     int
		insertedVals   []int
		used           bool
	}

	emptyBlocks := []diskBlock{}
	valueBlocks := []diskBlock{}

	currId := 0

	for i, val := range inputArr {
		if i%2 == 0 {
			valueBlocks = append(valueBlocks, diskBlock{size: val, value: currId, positionOnDisk: i, spacesLeft: 0})
			currId++
		} else {
			emptyBlocks = append(emptyBlocks, diskBlock{size: val, value: -1, positionOnDisk: i, spacesLeft: val, insertedVals: make([]int, 0), used: false})
		}
	}

	for i := len(valueBlocks) - 1; i >= 0; i-- {
		currVal := &valueBlocks[i]

		for j := 0; j <= len(emptyBlocks)-1; j++ {
			currEmpty := &emptyBlocks[j]

			if currEmpty.spacesLeft >= currVal.size {
				(*currEmpty).spacesLeft -= currVal.size
				(*currEmpty).insertedVals = append((*currEmpty).insertedVals, slices.Repeat([]int{currVal.value}, currVal.size)...)
				(*currVal).value = -1
				break
			}
		}
	}

	newBuffer := make([]int, 0)

	valIdx := 0
	emptyIdx := 0

	for i := 0; i <= (len(emptyBlocks) + len(valueBlocks) - 2); i++ {
		if i%2 == 0 {
			newBuffer = append(newBuffer, slices.Repeat([]int{valueBlocks[valIdx].value}, valueBlocks[valIdx].size)...)
			valIdx++
		} else {
			newBuffer = append(newBuffer, emptyBlocks[emptyIdx].insertedVals...)
			if emptyBlocks[emptyIdx].spacesLeft > 0 {
				newBuffer = append(newBuffer, slices.Repeat([]int{-1}, emptyBlocks[emptyIdx].spacesLeft)...)
			}
			emptyIdx++
		}
	}

	total := 0

	for i, val := range newBuffer {
		if val == -1 {
			continue
		}

		total += i * val
	}

	return strconv.Itoa(total)

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
