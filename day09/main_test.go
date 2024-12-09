package main

import (
	adventofcode "aoc"
	"log"
	"testing"
)

func TestChecksumPart1(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	denseMap := parseInput(f)

	diskmap := DiskMap(denseMap)
	expected := "00...111...2...333.44.5555.6666.777.888899"
	got := DiskMapToString(diskmap)
	adventofcode.Assert(
		got == expected, got,
	)

	expected = "0099811188827773336446555566.............."
	adventofcode.Assert(DiskMapToString(Compact(diskmap)) == expected)

	compacted := Compact(diskmap)
	adventofcode.Assert(IsCompacted(compacted))
	adventofcode.Assert(Checksum(Compact(diskmap)) == 1928)
}
