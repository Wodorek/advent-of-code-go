package util

import (
	"fmt"
	"os"
	"strings"
)

func PrintSolution(part int, solution string) {
	pth, _ := os.Getwd()
	split := strings.Split(pth, "/")

	fmt.Printf("the solution for year %s day %s part %d is: %s\n", split[len(split)-1], split[len(split)-2], part, solution)
}
