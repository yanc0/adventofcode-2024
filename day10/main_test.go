package main

import (
	. "aoc"
	"strings"
	"testing"
)

var testInput = `0123
1234
8765
9876`

func TestParseInput(t *testing.T) {
	topomap := parseInput(strings.NewReader(testInput))
	Assert(topomap[0][0] == 0)
	Assert(topomap[1][1] == 2)
	Assert(topomap[3][3] == 6)

	trailheads := topomap.FindTrailHeads()
	Assert(len(trailheads) == 1)
	Assert(trailheads[0].x == 0 && trailheads[0].y == 0)
}

func TestCountReachableSummits(t *testing.T) {
	topomap := parseInput(strings.NewReader(testInput))
	summitsCount := topomap.ReachableSummits(Position{0, 0})
	Assert(summitsCount == 1)
}

func TestCountHikingTrails(t *testing.T) {
	topomap := parseInput(strings.NewReader(testInput))
	distinctHikingTrailsCount := topomap.CountDistinctHikingTrails(Position{0, 0})
	Assert(distinctHikingTrailsCount == 16)
}
