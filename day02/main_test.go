package main

import (
	adventofcode "aoc"
	"strings"
	"testing"
)

func TestIsReportSafe(t *testing.T) {
	input := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	reports := parseReports(input)

	adventofcode.Assert(isReportSafe(reports[0]) == true, "first report should be safe")
	adventofcode.Assert(isReportSafe(reports[1]) == false, "first report should be unsafe")
	adventofcode.Assert(isReportSafe(reports[2]) == false, "first report should be unsafe")
	adventofcode.Assert(isReportSafe(reports[3]) == false, "first report should be unsafe")
	adventofcode.Assert(isReportSafe(reports[4]) == false, "first report should be unsafe")
	adventofcode.Assert(isReportSafe(reports[5]) == true, "first report should be safe")

}
