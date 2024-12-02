package main

import (
	adventofcode "aoc"
	"bufio"
	"embed"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input embed.FS

func main() {
	file, err := input.Open("input.txt")
	if err != nil {
		panic(err) // should not panic
	}

	reports := parseReports(file)

	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Println("number of safe reports:", safeReports)

	safeReports = 0
	for _, report := range reports {
		if isReportSafeRemovingOneLevel(report) {
			safeReports++
		}
	}
	fmt.Println("number of safe reports removing one bad level:", safeReports)

}

func parseReports(file io.Reader) [][]int {
	reports := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportLine := scanner.Text()
		levelsStr := strings.Split(reportLine, " ")
		adventofcode.Assert(len(levelsStr) >= 2, "report line must at least 2 levels")

		levels := make([]int, 0)
		for _, levelStr := range levelsStr {
			level, err := strconv.Atoi(levelStr)
			adventofcode.Assert(err == nil, "level must be int:", err)
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	return reports
}

type ReportWay int

const (
	Ascending ReportWay = iota
	Descending
	Unknown
)

// badLevel returns the position of the first bad level
// found in a report. Returns 0 if no bad level are found
func badLevel(report []int) int {
	way := Unknown
	for i, level := range report {
		if i == 0 {
			continue
		}

		if i == 1 {
			if report[i-1] > level {
				way = Descending
			}
			if report[i-1] < level {
				way = Ascending
			}
			if report[i-1] == level {
				return i
			}
		}

		if way == Ascending && report[i-1] >= level {
			return i
		}

		if way == Descending && report[i-1] <= level {
			return i
		}

		if way == Ascending && level-report[i-1] > 3 {
			return i
		}

		if way == Descending && report[i-1]-level > 3 {
			return i
		}
	}
	return 0
}

func isReportSafe(report []int) bool {
	return badLevel(report) == 0
}

func isReportSafeRemovingOneLevel(report []int) bool {
	if index := badLevel(report); index > 0 {
		if isReportSafe(deleteLevel(report, index-1)) {
			return true
		}

		if isReportSafe(deleteLevel(report, index)) {
			return true
		}

		return false
	}
	return true
}

func deleteLevel(report []int, index int) []int {
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)
	return slices.Delete(reportCopy, index, index+1)
}
