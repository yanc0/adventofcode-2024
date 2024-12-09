package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testInput = "2333133121414131402"

func TestChecksumPart1(t *testing.T) {
	denseMap := parseInput(strings.NewReader(testInput))

	diskmap := DiskMap(denseMap)
	expected := "00...111...2...333.44.5555.6666.777.888899"
	got := DiskMapToString(diskmap)
	Assert(
		got == expected, got,
	)

	expected = "0099811188827773336446555566.............."
	Assert(DiskMapToString(CompactV1(diskmap)) == expected)

	compacted := CompactV1(diskmap)
	Assert(IsCompacted(compacted))
	Assert(Checksum(CompactV1(diskmap)) == 1928)
}

func TestFileAndBlockOperations(t *testing.T) {
	denseMap := parseInput(strings.NewReader(testInput))

	diskmap := DiskMap(denseMap)
	expected := "00...111...2...333.44.5555.6666.777.888899"
	got := DiskMapToString(diskmap)
	Assert(got == expected)

	idx, size := findFile(diskmap, 0)
	Assert(idx == 0 && size == 2)

	idx, size = findFile(diskmap, 1)
	Assert(idx == 5 && size == 3)

	idx, size = findFile(diskmap, 8)
	Assert(idx == 36 && size == 4)

	idx, size = findFile(diskmap, 9)
	Assert(idx == 40 && size == 2)

	idx, size = findFile(diskmap, 10)
	Assert(idx == -1 && size == 0)

	Assert(FirstFreeContiguousSpace(diskmap, 1) == 2)
	Assert(FirstFreeContiguousSpace(diskmap, 2) == 2)
	Assert(FirstFreeContiguousSpace(diskmap, 3) == 2)
	Assert(FirstFreeContiguousSpace(diskmap, 4) == -1)
}

func TestChecksumPart2(t *testing.T) {
	denseMap := parseInput(strings.NewReader(testInput))

	diskmap := DiskMap(denseMap)
	expected := "00992111777.44.333....5555.6666.....8888.."
	got := DiskMapToString(CompactV2(diskmap))
	Assert(got == expected, got)
}
