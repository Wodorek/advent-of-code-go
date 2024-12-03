package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

type wire struct {
	operation   string
	provider1   string
	provider2   string
	output      string
	signalValue int
}

//go:embed input.txt
var input string

func main() {
	inputArr := parseInput(input)

	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func p1(wires []wire) string {

	circuit := make(map[string]int)

	for _, wire := range wires {
		fmt.Println(wire)
		switch wire.operation {
		case "SEND":
			circuit[wire.output] = wire.signalValue
		case "SEND FROM":
			fmt.Println(wire)
			circuit[wire.output] = circuit[wire.provider1]
		case "AND":
			circuit[wire.output] = circuit[wire.provider1] & circuit[wire.provider2]
		case "OR":
			circuit[wire.output] = circuit[wire.provider1] | circuit[wire.provider2]
		case "LSHIFT":
			shift, _ := strconv.Atoi(wire.provider2)
			circuit[wire.output] = circuit[wire.provider1] << shift
		case "RSHIFT":
			shift, _ := strconv.Atoi(wire.provider2)
			circuit[wire.output] = circuit[wire.provider1] >> shift

		case "NOT":
			circuit[wire.output] = 65535 ^ circuit[wire.provider1]
		}
	}

	return strconv.Itoa(circuit["a"])
}

func p2(wires []wire) string {
	return ""
}

func parseInput(inputStr string) []wire {
	split := strings.Split(inputStr, "\n")

	wires := make([]wire, 0)

	for _, line := range split {
		wire := wire{}
		splitLine := strings.Split(line, " ")

		val, err := strconv.Atoi(splitLine[0])
		if err == nil {
			wire.operation = "SEND"
			wire.signalValue = val
			wire.output = splitLine[2]
			wires = append(wires, wire)
			continue

		}

		if len(splitLine) == 3 {
			wire.operation = "SEND FROM"
			wire.provider1 = splitLine[0]
			wire.output = splitLine[2]

		} else if splitLine[0] == "NOT" {
			wire.operation = "NOT"
			wire.provider1 = splitLine[1]
			wire.output = splitLine[3]
		} else {

			wire.operation = splitLine[1]
			wire.provider1 = splitLine[0]
			wire.provider2 = splitLine[2]
			wire.output = splitLine[4]
		}

		wires = append(wires, wire)
	}

	return wires

}
