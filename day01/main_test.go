package main

import (
	adventofcode "aoc"
	"slices"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `5   2
3   6
1   4`

	left, right := parseInput(strings.NewReader(input))
	slices.Sort(left)
	slices.Sort(right)

	adventofcode.Assert(len(left) == 3, "left slice length must be 3")
	adventofcode.Assert(len(right) == 3, "right slice length must be 3")
}

func TestSumDistances(t *testing.T) {
	left := []int{1, 2, 4}
	right := []int{2, 3, 3}

	adventofcode.Assert(sumDistances(left, right) == 3, "distance must be 3")
}
