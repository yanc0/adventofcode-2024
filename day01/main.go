package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := parseInput(InputFile())
	slices.Sort(left)
	slices.Sort(right)

	fmt.Println("distances sum is:", sumDistances(left, right))
	fmt.Println("similarity score is:", scoreSimilarity(left, right))
}

func sumDistances(left, right []int) int {
	Assert(len(left) == len(right), "left and right slices have different sizes")
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
		Assert(len(numbers) == 2, "line must contains 2 numbers")
		leftNum, err := strconv.Atoi(numbers[0])
		Assert(err == nil, "leftNum must be parsed correctly", err)
		rightNum, err := strconv.Atoi(numbers[1])
		Assert(err == nil, "rightNum must be parsed correctly", err)
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	Assert(len(left) == len(right), "left and right slices have different sizes")
	return left, right
}

func countOccurences(num int, s []int) (count int) {
	for _, value := range s {
		if num == value {
			count++
		}
	}
	return count
}

func scoreSimilarity(left, right []int) (score int) {
	for _, leftNum := range left {
		score += countOccurences(leftNum, right) * leftNum
	}
	return score
}
