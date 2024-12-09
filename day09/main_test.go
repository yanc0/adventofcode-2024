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
	adventofcode.Assert(DiskMapToString(CompactV1(diskmap)) == expected)

	compacted := CompactV1(diskmap)
	adventofcode.Assert(IsCompacted(compacted))
	adventofcode.Assert(Checksum(CompactV1(diskmap)) == 1928)
}

func TestFileAndBlockOperations(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	denseMap := parseInput(f)

	diskmap := DiskMap(denseMap)
	expected := "00...111...2...333.44.5555.6666.777.888899"
	got := DiskMapToString(diskmap)
	adventofcode.Assert(got == expected)

	idx, size := findFile(diskmap, 0)
	adventofcode.Assert(idx == 0 && size == 2)

	idx, size = findFile(diskmap, 1)
	adventofcode.Assert(idx == 5 && size == 3)

	idx, size = findFile(diskmap, 8)
	adventofcode.Assert(idx == 36 && size == 4)

	idx, size = findFile(diskmap, 9)
	adventofcode.Assert(idx == 40 && size == 2)

	idx, size = findFile(diskmap, 10)
	adventofcode.Assert(idx == -1 && size == 0)

	adventofcode.Assert(FirstFreeContiguousSpace(diskmap, 1) == 2)
	adventofcode.Assert(FirstFreeContiguousSpace(diskmap, 2) == 2)
	adventofcode.Assert(FirstFreeContiguousSpace(diskmap, 3) == 2)
	adventofcode.Assert(FirstFreeContiguousSpace(diskmap, 4) == -1)
}

func TestChecksumPart2(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	denseMap := parseInput(f)

	diskmap := DiskMap(denseMap)
	expected := "00992111777.44.333....5555.6666.....8888.."
	got := DiskMapToString(CompactV2(diskmap))
	adventofcode.Assert(got == expected, got)
}
