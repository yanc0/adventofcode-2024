package main

import (
	"fmt"
	"strings"
	"testing"
)

var testInput = `......#....#
...#....0...
....#0....#.
..#....0....
....0....#..
.#....A.....
...#........
#......#....
........A...
.........A..
..........#.
..........#.`

func TestInput(t *testing.T) {
	antennas := parseInput(strings.NewReader(testInput))
	_ = antennas
}

func TestAntennasAntinodes(t *testing.T) {
	antennas := parseInput(strings.NewReader(testInput))
	antinodes := NewAntinodesMap(len(antennas))
	antinodes.Mark(antennas.CalculateAntinodesForFreqV1('0')...)
	fmt.Println(antinodes.Count())
}
