package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testInput = `AAAA
BBCD
BBCC
EEEC`

func TestParseInput(t *testing.T) {
	m := parseInput(strings.NewReader(testInput))
	Assert(m[0][0].Type == 'A')
	Assert(m[0][3].Type == 'A')
	Assert(m[2][1].Type == 'B')
	Assert(m[3][3].Type == 'C')
}

func TestRegions(t *testing.T) {
	m := parseInput(strings.NewReader(testInput))
	m.markRegions()
	Assert(m[0][3].RegionID == 1)
	Assert(m[1][0].RegionID == 2)

	Assert(len(m.GetRegionIDs()) == 5)
}

func TestAreaPerimeter(t *testing.T) {
	m := parseInput(strings.NewReader(testInput))
	m.markRegions()

	Assert(m.GetRegion(1).Area() == 4)
	Assert(m.GetRegion(1).Perimeter() == 10)

	Assert(m.GetRegion(3).Area() == 4)

	Assert(m.GetRegion(4).Area() == 1)
	Assert(m.GetRegion(4).Perimeter() == 4)
}

func TestDirectionnalSideScan(t *testing.T) {
	m := parseInput(strings.NewReader(testInput))

	Assert(m.GetRegion(1).verticalSideScan(0) == 1)
	Assert(m.GetRegion(1).verticalSideScan(1) == 0)
	Assert(m.GetRegion(1).verticalSideScan(2) == 0)
	Assert(m.GetRegion(1).verticalSideScan(3) == 0)
	Assert(m.GetRegion(1).verticalSideScan(4) == 1)

	Assert(m.GetRegion(3).verticalSideScan(0) == 1)
	Assert(m.GetRegion(3).verticalSideScan(1) == 2)
	Assert(m.GetRegion(3).verticalSideScan(2) == 1)

	Assert(m.GetRegion(1).horizontalSideScan(-1) == 1)
	Assert(m.GetRegion(1).horizontalSideScan(0) == 1)

	Assert(m.GetRegion(3).horizontalSideScan(-1) == 1)
	Assert(m.GetRegion(3).horizontalSideScan(0) == 1)
	Assert(m.GetRegion(3).horizontalSideScan(1) == 1)
	Assert(m.GetRegion(3).horizontalSideScan(2) == 1)
}

func TestSideScan(t *testing.T) {
	m := parseInput(strings.NewReader(testInput))

	Assert(m.GetRegion(1).Sides() == 4)
	Assert(m.GetRegion(2).Sides() == 4)
	Assert(m.GetRegion(3).Sides() == 8)
	Assert(m.GetRegion(4).Sides() == 4)
	Assert(m.GetRegion(5).Sides() == 4)
}

func TestSideScanBug1(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	m := parseInput(strings.NewReader(input))

	Assert(m.GetRegion(1).Sides() == 10)
	Assert(m.GetRegion(3).Sides() == 22)
}

func TestSideScanBug2(t *testing.T) {
	input := `AAAAAA
AAA**A
AAA**A
A**AAA
A**AAA
AAAAAA`
	m := parseInput(strings.NewReader(input))

	sides := m.GetRegion(1).Sides()
 	Assert(sides == 12)
	Assert(m.GetRegion(2).Sides() == 4)
	Assert(m.GetRegion(3).Sides() == 4)
}
