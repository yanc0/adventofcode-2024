package main

import (
	adventofcode "aoc"
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

//go:embed *.txt
var input embed.FS

func main() {
	f, err := input.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	problems := parseInput(f)

	totalSolutions := 0
	withPart2 := false
	for _, problem := range problems {
		if problem.IsSolvable(withPart2) {
			totalSolutions += problem.solution
		}
	}
	fmt.Println("part 1: sum of solution that can possibly be true:", totalSolutions)

	totalSolutions = 0
	withPart2 = true
	for _, problem := range problems {
		if problem.IsSolvable(withPart2) {
			totalSolutions += problem.solution
		}
	}
	fmt.Println("part 2: sum of solution that can possibly be true:", totalSolutions)
}

type Problem struct {
	solution int
	numbers  []int
}

func parseInput(input io.Reader) []Problem {
	problems := make([]Problem, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		splittedProblem := strings.Split(line, ":")
		adventofcode.Assert(len(splittedProblem) == 2, "wrong line format")
		problems = append(problems,
			Problem{
				solution: Int(splittedProblem[0]),
				numbers:  Ints(strings.Split(strings.TrimSpace(splittedProblem[1]), " ")),
			},
		)
	}
	return problems
}

func Int(str string) int {
	num, err := strconv.Atoi(str)
	adventofcode.Assert(err == nil, str, "is not an int:", err)
	return num
}

func Ints(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i] = Int(str)
	}
	return ints
}

func (p Problem) IsSolvable(withPart2 bool) bool {
	return isSolvable(p.solution, p.numbers[1:], p.numbers[0], withPart2)
}

func isSolvable(solution int, nums []int, num int, withPart2 bool) bool {
	if len(nums) == 0 {
		return num == solution
	}
	// returns early as num can only grows
	if num > solution {
		return false
	}
	if isSolvable(solution, nums[1:], num+nums[0], withPart2) {
		return true
	}
	if isSolvable(solution, nums[1:], num*nums[0], withPart2) {
		return true
	}
	if withPart2 && isSolvable(solution, nums[1:], combine(num, nums[0]), withPart2) {
		return true
	}

	return false
}

func combine(a, b int) int {
	num, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	adventofcode.Assert(err == nil, err)
	return num
}
