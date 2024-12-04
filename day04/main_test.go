package main

import (
	adventofcode "aoc"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	grid := parseInput(input)
	adventofcode.Assert(len(grid) == 10, "grid should contain 10 rows")
	adventofcode.Assert(len(grid[0]) == 10, "row should contain 10 columns")
	adventofcode.Assert(grid[0][0] == 'M', "first letter is M")
	adventofcode.Assert(grid[9][9] == 'X', "first letter is X")
}

func TestReadXMAS(t *testing.T) {
	input := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	grid := parseInput(input)

	adventofcode.Assert(grid.SearchString(HorizontalForward, "XMAS", -1, -1) == false, "should not panic")
	adventofcode.Assert(grid.SearchString(HorizontalForward, "XMAS", 0, 5))
	adventofcode.Assert(grid.SearchString(HorizontalBackward, "XMAS", 1, 4))
	adventofcode.Assert(grid.SearchString(VerticalForward, "XMAS", 3, 9))
	adventofcode.Assert(grid.SearchString(VerticalBackward, "XMAS", 9, 9))
	adventofcode.Assert(grid.SearchString(RightDiagonalForward, "XMAS", 0, 4))
	adventofcode.Assert(grid.SearchString(RightDiagonalBackward, "XMAS", 9, 3))
	adventofcode.Assert(grid.SearchString(LeftDiagonalBackward, "XMAS", 5, 0))
	adventofcode.Assert(grid.SearchString(LeftDiagonalForward, "XMAS", 3, 9))
}
