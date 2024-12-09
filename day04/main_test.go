package main

import (
	. "aoc"
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
	Assert(len(grid) == 10, "grid should contain 10 rows")
	Assert(len(grid[0]) == 10, "row should contain 10 columns")
	Assert(grid[0][0] == 'M', "first letter is M")
	Assert(grid[9][9] == 'X', "first letter is X")
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

	Assert(grid.SearchString(HorizontalForward, "XMAS", -1, -1) == false, "should not panic")
	Assert(grid.SearchString(HorizontalForward, "XMAS", 0, 5))
	Assert(grid.SearchString(HorizontalBackward, "XMAS", 1, 4))
	Assert(grid.SearchString(VerticalForward, "XMAS", 3, 9))
	Assert(grid.SearchString(VerticalBackward, "XMAS", 9, 9))
	Assert(grid.SearchString(RightDiagonalForward, "XMAS", 0, 4))
	Assert(grid.SearchString(RightDiagonalBackward, "XMAS", 9, 3))
	Assert(grid.SearchString(LeftDiagonalBackward, "XMAS", 5, 0))
	Assert(grid.SearchString(LeftDiagonalForward, "XMAS", 3, 9))
}
