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
	left, right := parseInput(file)
	slices.Sort(left)
	slices.Sort(right)

	fmt.Println("distances sum is:", sumDistances(left, right))
}

func sumDistances(left, right []int) int {
	adventofcode.Assert(len(left) == len(right), "left and right slices have different sizes")
	sum := 0
	for i := 0; i < len(left); i++ {
		distance := right[i] - left[i]
		if distance < 0 {
			distance *= -1
		}
		sum += distance
	}
	return sum
}

func parseInput(file io.Reader) (left []int, right []int) {
	left = make([]int, 0)
	right = make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		adventofcode.Assert(len(numbers) == 2, "line must contains 2 numbers")
		leftNum, err := strconv.Atoi(numbers[0])
		adventofcode.Assert(err == nil, "leftNum must be parsed correctly", err)
		rightNum, err := strconv.Atoi(numbers[1])
		adventofcode.Assert(err == nil, "rightNum must be parsed correctly", err)
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	adventofcode.Assert(len(left) == len(right), "left and right slices have different sizes")
	return left, right
}
