package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestInput(t *testing.T) {
	problems := parseInput(strings.NewReader(testInput))
	Assert(problems[0].solution == 190, "expected first solution is 190")
	Assert(len(problems) == 9, "there should be 9 problems")
	Assert(len(problems[8].numbers) == 4, "there should be 4 number for the 9th problem")
}

func TestProblemResolution(t *testing.T) {
	problems := parseInput(strings.NewReader(testInput))
	Assert(problems[0].IsSolvable(true), "problem 1 should be solvable")
	Assert(problems[1].IsSolvable(true), "problem 1 should be solvable")
	Assert(problems[8].IsSolvable(true), "problem 1 should be solvable")
}
