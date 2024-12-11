package main

import (
	. "aoc"
	"fmt"
	"io"
	"strings"
)

func main() {
	stones := parseInput(InputFile())

	for _, blinks := range []int{25, 75} {
		fmt.Printf("There are %d stones after %d blinks\n", RecursiveBlinkN(stones, blinks), blinks)
	}
}

func parseInput(input io.Reader) Stones {
	stones := make([]int, 0)
	line, err := io.ReadAll(input)
	Assert(err == nil, "failed to read input")
	for _, stone := range strings.Split(string(line), " ") {
		stones = append(stones, Int(stone))
	}
	return stones
}

type Stones []int

func (s Stones) String() string {
	str := ""
	for _, stone := range s {
		str += fmt.Sprintf("%d ", stone)
	}
	return strings.TrimSpace(str)
}

func splitInHalf(s string) (string, string) {
	return s[0 : len(s)/2], s[len(s)/2:]
}

type cache map[string]int

func (c cache) set(stone int, num int, result int) {
	c[fmt.Sprintf("%d|%d", stone, num)] = result
}

func (c cache) get(stone int, num int) (int, bool) {
	i, exists := c[fmt.Sprintf("%d|%d", stone, num)]
	return i, exists
}

func RecursiveBlink(c cache, stone int, num int) int {
	if num == 0 {
		return 1
	}

	var res int
	var ok bool
	if res, ok = c.get(stone, num); ok {
		return res
	}

	stoneStr := fmt.Sprintf("%d", stone)
	stoneHasEvenNumberOfDigits := len(stoneStr)%2 == 0

	switch {
	case stone == 0:
		res = RecursiveBlink(c, 1, num-1)
	case stoneHasEvenNumberOfDigits:
		left, right := splitInHalf(stoneStr)
		res = RecursiveBlink(c, Int(left), num-1) + RecursiveBlink(c, Int(right), num-1)
	default:
		res = RecursiveBlink(c, stone*2024, num-1)
	}

	c.set(stone, num, res)

	return res
}

func RecursiveBlinkN(stones Stones, n int) int {
	c := make(cache)
	count := 0
	for _, s := range stones {
		count += RecursiveBlink(c, s, n)
	}
	return count
}
