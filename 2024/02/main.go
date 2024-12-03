package main

import (
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	inputArr := prepareInput(input)
	util.PrintSolution(1, p1(inputArr))
	util.PrintSolution(2, p2(inputArr))
}

func prepareInput(inputStr string) [][]int {
	reports := make([][]int, 0)

	split := strings.Split(inputStr, "\n")

	for _, report := range split {
		splitReport := strings.Split(report, " ")
		parsedReport := make([]int, 0)
		for _, num := range splitReport {
			converted, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			parsedReport = append(parsedReport, converted)
		}
		reports = append(reports, parsedReport)
	}

	return reports
}

func p1(inputArr [][]int) string {
	totalSafe := 0

	for _, report := range inputArr {

		reportType := "ASC"
		if report[0] > report[1] {
			reportType = "DESC"
		}

		if checkIsSafe(report, reportType) {
			totalSafe++
		}
	}

	return strconv.Itoa(totalSafe)
}

func p2(inputArr [][]int) string {
	totalSafe := 0

	for _, report := range inputArr {

		reportType := "ASC"
		if report[0] > report[1] {
			reportType = "DESC"
		}

		isSafe := checkIsSafe(report, reportType)

		if isSafe {
			totalSafe++
		} else {
			safeSubreports := 0
			copies := make([][]int, 0)
			for j := range report {
				copied := slices.Clone(report)
				copied = slices.Delete(copied, j, j+1)
				copies = append(copies, copied)
			}

			for _, subreport := range copies {
				reportType := "ASC"
				if subreport[0] > subreport[1] {
					reportType = "DESC"
				}
				if checkIsSafe(subreport, reportType) {
					safeSubreports++
				}
			}
			if safeSubreports >= 1 {
				totalSafe++
			}
		}
	}

	return strconv.Itoa(totalSafe)
}

func checkIsSafe(report []int, reportType string) bool {
	for i := 0; i < len(report)-1; i++ {
		switch reportType {
		case "ASC":
			if report[i] > report[i+1] || (report[i+1]-report[i] > 3) || (report[i+1]-report[i] < 1) {
				return false
			}
		case "DESC":
			if report[i] < report[i+1] || (report[i]-report[i+1] > 3) || (report[i]-report[i+1] < 1) {
				return false
			}
		}
	}

	return true
}
