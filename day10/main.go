package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"io"
)

func main() {
	topomap := parseInput(InputFile())

	// part 1
	totalReachableSummits := 0
	for _, starting := range topomap.FindTrailHeads() {
		totalReachableSummits += topomap.ReachableSummits(starting)
	}
	fmt.Println("total of reachable summits:", totalReachableSummits)

	// part 2
	totalDistinctHikingTrails := 0
	for _, starting := range topomap.FindTrailHeads() {
		totalDistinctHikingTrails += topomap.CountDistinctHikingTrails(starting)
	}
	fmt.Println("total of distrinct hiking trails:", totalDistinctHikingTrails)
}

type TopographicMap [][]int

func parseInput(input io.Reader) TopographicMap {
	topomap := make([][]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		latitudes := make([]int, 0)
		for _, altitude := range scanner.Text() {
			if altitude == '.' {
				latitudes = append(latitudes, -1)
				continue
			}
			latitudes = append(latitudes, Int(string(altitude)))
		}
		topomap = append(topomap, latitudes)
	}
	return topomap
}

func (t TopographicMap) FindTrailHeads() []Position {
	positions := make([]Position, 0)
	for x := 0; x < len(t); x++ {
		for y := 0; y < len(t); y++ {
			if t[x][y] == 0 {
				positions = append(positions, Position{X: x, Y: y})
			}
		}
	}
	return positions
}

func (t TopographicMap) MaxElevation() int {
	maxElevation := 0
	for x := 0; x < len(t); x++ {
		for y := 0; y < len(t); y++ {
			if t[x][y] > maxElevation {
				maxElevation = t[x][y]
			}
		}
	}
	return maxElevation
}

func (t TopographicMap) ReachableSummits(starting Position) int {
	summits := make(map[string]bool)
	t.MarkReachableSummits(starting, summits)
	return len(summits)
}

func (t TopographicMap) MarkReachableSummits(starting Position, summits map[string]bool) {
	startingIsSummit := t.Elevation(starting) == t.MaxElevation()
	if startingIsSummit {
		summits[starting.String()] = true
		return
	}

	if starting.Up().ExistsOnMatrix(len(t)) && t.Elevation(starting.Up()) == t.Elevation(starting)+1 {
		t.MarkReachableSummits(starting.Up(), summits)
	}
	if starting.Right().ExistsOnMatrix(len(t)) && t.Elevation(starting.Right()) == t.Elevation(starting)+1 {
		t.MarkReachableSummits(starting.Right(), summits)
	}
	if starting.Down().ExistsOnMatrix(len(t)) && t.Elevation(starting.Down()) == t.Elevation(starting)+1 {
		t.MarkReachableSummits(starting.Down(), summits)
	}
	if starting.Left().ExistsOnMatrix(len(t)) && t.Elevation(starting.Left()) == t.Elevation(starting)+1 {
		t.MarkReachableSummits(starting.Left(), summits)
	}
}

func (t TopographicMap) CountDistinctHikingTrails(starting Position) int {
	startingIsSummit := t.Elevation(starting) == t.MaxElevation()
	if startingIsSummit {
		return 1
	}

	countDistinctHikingTrails := 0

	if starting.Up().ExistsOnMatrix(len(t)) && t.Elevation(starting.Up()) == t.Elevation(starting)+1 {
		countDistinctHikingTrails += t.CountDistinctHikingTrails(starting.Up())
	}
	if starting.Right().ExistsOnMatrix(len(t)) && t.Elevation(starting.Right()) == t.Elevation(starting)+1 {
		countDistinctHikingTrails += t.CountDistinctHikingTrails(starting.Right())
	}
	if starting.Down().ExistsOnMatrix(len(t)) && t.Elevation(starting.Down()) == t.Elevation(starting)+1 {
		countDistinctHikingTrails += t.CountDistinctHikingTrails(starting.Down())
	}
	if starting.Left().ExistsOnMatrix(len(t)) && t.Elevation(starting.Left()) == t.Elevation(starting)+1 {
		countDistinctHikingTrails += t.CountDistinctHikingTrails(starting.Left())
	}

	return countDistinctHikingTrails
}

func (t TopographicMap) Elevation(at Position) int {
	Assert(at.ExistsOnMatrix(len(t)))

	return t[at.X][at.Y]
}
