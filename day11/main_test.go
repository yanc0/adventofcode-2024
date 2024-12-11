package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testInput = `0 1 10 99 999`

func TestParseInput(t *testing.T) {
	stones := parseInput(strings.NewReader(testInput))
	Assert(stones[0] == 0)
	Assert(stones[1] == 1)
	Assert(stones[2] == 10)
	Assert(stones[3] == 99)
	Assert(stones[4] == 999)
}

func TestSplitInHalf(t *testing.T) {
	left, right := splitInHalf("1010")
	Assert(Int(left) == 10 && Int(right) == 10)

	left, right = splitInHalf("010010")
	Assert(Int(left) == 10 && Int(right) == 10)
}

func TestRecursiveBlinkN(t *testing.T) {
	stones := parseInput(strings.NewReader("125 17"))
	stonesCount := RecursiveBlinkN(stones, 6)
	Assert(stonesCount == 22)
}

func BenchmarkRecursiveBlinkN(b *testing.B) {
	stones := parseInput(strings.NewReader("125 17"))

	b.ResetTimer()
	b.SetParallelism(1)
	for n := 0; n < b.N; n++ {
		Assert(RecursiveBlinkN(stones, 6) == 22)
	}

}
