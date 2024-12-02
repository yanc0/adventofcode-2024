package main

import (
	adventofcode "aoc"
	"bufio"
	"embed"
	"fmt"
	"io"
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

	safeReports := 0
	reports := parseReports(file)

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Println("number of safe reports:", safeReports)
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

type ReportWay string

const (
	Ascending  ReportWay = "asc"
	Descending ReportWay = "desc"
	Unknown    ReportWay = "unkn"
)

func isReportSafe(report []int) bool {
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
				return false
			}
		}

		if way == Ascending && report[i-1] >= level {
			return false
		}

		if way == Descending && report[i-1] <= level {
			return false
		}

		if way == Ascending && level-report[i-1] > 3 {
			return false
		}

		if way == Descending && report[i-1]-level > 3 {
			return false
		}
	}
	return true
}
