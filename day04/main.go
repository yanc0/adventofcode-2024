package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
)

//go:embed input.txt
var input embed.FS

func main() {
	f, err := input.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := parseInput(f)
	count := 0
	orders := []Order{
		HorizontalForward,
		HorizontalBackward,
		VerticalForward,
		VerticalBackward,
		RightDiagonalForward,
		RightDiagonalBackward,
		LeftDiagonalForward,
		LeftDiagonalBackward,
	}
	for _, order := range orders {
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid); y++ {
				if grid.SearchString(order, "XMAS", x, y) {
					count++
				}
			}
		}
	}
	fmt.Printf("we counted %d XMAS\n", count)

	count = 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid); y++ {
			if grid.SearchString(RightDiagonalForward, "MAS", x, y) || grid.SearchString(RightDiagonalForward, "SAM", x, y) {
				if grid.SearchString(LeftDiagonalForward, "MAS", x, y+2) || grid.SearchString(LeftDiagonalForward, "SAM", x, y+2) {
					count++
				}
			}
		}
	}

	fmt.Printf("we counted %d X-MAS\n", count)
}

type Grid [][]rune

type Order int

const (
	HorizontalForward Order = iota
	HorizontalBackward
	VerticalForward
	VerticalBackward
	RightDiagonalForward
	RightDiagonalBackward
	LeftDiagonalForward
	LeftDiagonalBackward
)

func (g Grid) Read(x, y int) rune {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g) {
		return '.'
	}
	return g[x][y]
}

func (g Grid) SearchString(o Order, s string, x int, y int) bool {
	str := ""
	switch o {
	case HorizontalForward:
		for j := y; j < y+len(s); j++ {
			str += string(g.Read(x, j))
		}
	case HorizontalBackward:
		for j := y; j > y-len(s); j-- {
			str += string(g.Read(x, j))
		}
	case VerticalForward:
		for i := x; i < x+len(s); i++ {
			str += string(g.Read(i, y))
		}
	case VerticalBackward:
		for i := x; i > x-len(s); i-- {
			str += string(g.Read(i, y))
		}
	case RightDiagonalForward:
		for i := 0; i < len(s); i++ {
			str += string(g.Read(x+i, y+i))
		}
	case RightDiagonalBackward:
		for i := 0; i < len(s); i++ {
			str += string(g.Read(x-i, y-i))
		}
	case LeftDiagonalForward:
		for i := 0; i < len(s); i++ {
			str += string(g.Read(x+i, y-i))
		}
	case LeftDiagonalBackward:
		for i := 0; i < len(s); i++ {
			str += string(g.Read(x-i, y+i))
		}
	default:
		return false
	}
	return str == s
}

func parseInput(input io.Reader) Grid {
	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := make([]rune, 0)
		line := scanner.Text()
		for _, ch := range line {
			row = append(row, ch)
		}
		grid = append(grid, row)
	}
	return grid
}
