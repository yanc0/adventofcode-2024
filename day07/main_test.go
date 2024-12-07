package main

import (
	adventofcode "aoc"
	"log"
	"testing"
)

func TestInput(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	problems := parseInput(f)
	adventofcode.Assert(problems[0].solution == 190, "expected first solution is 190")
	adventofcode.Assert(len(problems) == 9, "there should be 9 problems")
	adventofcode.Assert(len(problems[8].numbers) == 4, "there should be 4 number for the 9th problem")
}

func TestProblemResolution(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	problems := parseInput(f)
	adventofcode.Assert(problems[0].IsSolvable(true), "problem 1 should be solvable")
	adventofcode.Assert(problems[1].IsSolvable(true), "problem 1 should be solvable")
	adventofcode.Assert(problems[8].IsSolvable(true), "problem 1 should be solvable")
}
