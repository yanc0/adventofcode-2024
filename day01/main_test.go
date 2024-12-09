package main

import (
	. "aoc"
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

	Assert(len(left) == 3, "left slice length must be 3")
	Assert(len(right) == 3, "right slice length must be 3")
}

func TestSumDistances(t *testing.T) {
	left := []int{1, 2, 4}
	right := []int{2, 3, 3}

	Assert(sumDistances(left, right) == 3, "distance must be 3")
}

func TestSimilarityScore(t *testing.T) {
	left := []int{1, 3, 4}
	right := []int{4, 3, 3}

	Assert(scoreSimilarity(left, right) == 10, "distance must be 10")
}
