package main

import (
	. "aoc"
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

	Assert(isReportSafe(reports[0]) == true, "first report should be safe")
	Assert(isReportSafe(reports[1]) == false, "second report should be unsafe")
	Assert(isReportSafe(reports[2]) == false, "third report should be unsafe")
	Assert(isReportSafe(reports[3]) == false, "fourth report should be unsafe")
	Assert(isReportSafe(reports[4]) == false, "fifth report should be unsafe")
	Assert(isReportSafe(reports[5]) == true, "sixth report should be safe")
}

func TestBadLevel(t *testing.T) {
	input := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	reports := parseReports(input)
	Assert(badLevel(reports[0]) == 0, "first report should be 0")
	Assert(badLevel(reports[1]) == 2, "second report should be 2")
	Assert(badLevel(reports[2]) == 3, "third report should be 3")
	Assert(badLevel(reports[3]) == 2, "fourth report should be 2")
	Assert(badLevel(reports[4]) == 3, "fifth report should be 3")
	Assert(badLevel(reports[5]) == 0, "sixth report should be 0")
}

func TestIsReportSafeRemovingOneLevel(t *testing.T) {
	input := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

	reports := parseReports(input)

	Assert(isReportSafeRemovingOneLevel(reports[0]) == true, "first report should be safe")
	Assert(isReportSafeRemovingOneLevel(reports[1]) == false, "second report should be unsafe")
	Assert(isReportSafeRemovingOneLevel(reports[2]) == false, "third report should be unsafe")
	Assert(isReportSafeRemovingOneLevel(reports[3]) == true, "fourth report should be safe")
	Assert(isReportSafeRemovingOneLevel(reports[4]) == true, "fifth report should be safe")
	Assert(isReportSafeRemovingOneLevel(reports[5]) == true, "sixth report should be safe")
}
