package main

import (
	. "aoc"
	"bufio"
	"fmt"
	"io"
	"slices"
)

type RayPoint uint8

const (
	Same RayPoint = iota
	VoidToFull
	FullToVoid
	Unknown
)

func main() {
	m := parseInput(InputFile())

	price := 0
	for _, region := range m.GetRegions() {
		price += region.Area() * region.Perimeter()
	}
	fmt.Println("price for fences with perimeter:", price)

	price = 0
	for _, region := range m.GetRegions() {
		price += region.Area() * region.Sides()
	}
	fmt.Println("price for fences with sides:", price)
}

func parseInput(input io.Reader) Map {
	m := make(Map, 0)
	scanner := bufio.NewScanner(input)
	x := 0
	for scanner.Scan() {
		row := make([]Plot, len(scanner.Text()))
		for y, c := range scanner.Text() {
			row[y] = Plot{
				Type:     c,
				Position: Position{X: x, Y: y},
				RegionID: 0,
			}
		}
		m = append(m, row)
		x++
	}
	m.markRegions()
	return m
}

type Map [][]Plot
type Region []Plot
type Plot struct {
	Type     rune
	Position Position
	RegionID int
}

var VOID Plot = Plot{
	Type: ' ',
}

func (m Map) GetPlot(p Position) Plot {
	if !p.ExistsOnMatrix(len(m)) {
		return VOID
	}
	return m[p.X][p.Y]
}

func (p Plot) RegionMarked() bool {
	return p.RegionID > 0
}

func (p Plot) Void() bool {
	return p.Type == VOID.Type
}

func (m Map) markRegions() {
	regionID := 1
	m.ForEach(func(p Plot) {
		if m.markRegionID(p.Position, regionID) {
			regionID++
		}
	})
}

func (m Map) ForEach(f func(p Plot)) {
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m); y++ {
			f(m.GetPlot(Position{X: x, Y: y}))
		}
	}
}

func (m Map) RegionCount() int {
	return len(m.GetRegionIDs())
}

func (m Map) GetRegionIDs() []int {
	ids := make([]int, 0)
	m.ForEach(func(p Plot) {
		if !slices.Contains(ids, p.RegionID) {
			ids = append(ids, p.RegionID)
		}
	})
	return ids
}

func (m Map) GetRegions() []Region {
	regions := make([]Region, 0)
	for _, rid := range m.GetRegionIDs() {
		regions = append(regions, m.GetRegion(rid))
	}
	return regions
}

func (m Map) GetRegion(id int) Region {
	r := make(Region, 0)
	m.ForEach(func(p Plot) {
		if p.RegionID == id {
			r = append(r, p)
		}
	})
	return r
}

func (r Region) Area() int {
	return len(r)
}

func (r Region) PositionExists(p Position) bool {
	for _, plot := range r {
		if plot.Position.X == p.X && plot.Position.Y == p.Y {
			return true
		}
	}
	return false
}

func (r Region) GetPlot(p Position) Plot {
	for _, plot := range r {
		if plot.Position.X == p.X && plot.Position.Y == p.Y {
			return plot
		}
	}
	return VOID
}

func (r Region) Perimeter() int {
	perimeter := 0
	for _, plot := range r {
		if !r.PositionExists(plot.Position.Up()) {
			perimeter++
		}
		if !r.PositionExists(plot.Position.Right()) {
			perimeter++
		}
		if !r.PositionExists(plot.Position.Down()) {
			perimeter++
		}
		if !r.PositionExists(plot.Position.Left()) {
			perimeter++
		}
	}
	return perimeter
}

func (r Region) MaxRow() int {
	max := -1
	for _, plot := range r {
		if plot.Position.X > max {
			max = plot.Position.X
		}
	}
	Assert(max >= 0, "max row must return a valid x position")
	return max
}

func (r Region) MinRow() int {
	min := r.MaxRow()
	for _, plot := range r {
		if plot.Position.X < min {
			min = plot.Position.X
		}
	}
	Assert(min >= 0, "min row must return a valid x position")
	return min
}

func (r Region) MaxCol() int {
	max := -1
	for _, plot := range r {
		if plot.Position.Y > max {
			max = plot.Position.Y
		}
	}
	Assert(max >= 0, "max row must return a valid y position")
	return max
}

func (r Region) MinCol() int {
	min := r.MaxCol()
	for _, plot := range r {
		if plot.Position.Y < min {
			min = plot.Position.Y
		}
	}
	Assert(min >= 0, "min row must return a valid y position")
	return min
}

func (r Region) Sides() int {
	sides := 0

	for x := -1; x <= r.MaxRow(); x++ {
		sides += r.horizontalSideScan(x)
	}

	for y := 0; y <= r.MaxCol()+1; y++ {
		sides += r.verticalSideScan(y)
	}

	return sides
}

func (m Map) markRegionID(p Position, id int) bool {
	Assert(id > 0, "region must be greater than 0")
	if m.GetPlot(p).RegionMarked() {
		return false
	}

	m.setRegion(p, id)

	if m.GetPlot(p.Up()).Type == m.GetPlot(p).Type {
		m.markRegionID(p.Up(), id)
	}
	if m.GetPlot(p.Right()).Type == m.GetPlot(p).Type {
		m.markRegionID(p.Right(), id)
	}
	if m.GetPlot(p.Down()).Type == m.GetPlot(p).Type {
		m.markRegionID(p.Down(), id)
	}
	if m.GetPlot(p.Left()).Type == m.GetPlot(p).Type {
		m.markRegionID(p.Left(), id)
	}

	return true
}

func (m Map) setRegion(p Position, id int) {
	Assert(p.ExistsOnMatrix(len(m)))
	m[m.GetPlot(p).Position.X][m.GetPlot(p).Position.Y].RegionID = id
}

func (r Region) verticalSideScan(y int) (sides int) {
	ray := make([]RayPoint, 0)
	for x := r.MinRow(); x <= r.MaxRow(); x++ {
		currentPos := Position{X: x, Y: y + r.MinCol()}
		ray = append(ray, ternary(r.GetPlot(currentPos), r.GetPlot(currentPos.Left())))
	}

	return sequences(ray)
}

func (r Region) horizontalSideScan(x int) (sides int) {
	ray := make([]RayPoint, 0)
	for y := r.MinCol(); y <= r.MaxCol(); y++ {
		currentPos := Position{X: x + r.MinRow(), Y: y}
		ray = append(ray, ternary(r.GetPlot(currentPos), r.GetPlot(currentPos.Down())))
	}

	return sequences(ray)
}

func sequences(ray []RayPoint) (num int) {
	previousSide := Unknown
	for i := 0; i < len(ray); i++ {
		if previousSide != ray[i] && (ray[i] == VoidToFull || ray[i] == FullToVoid) {
			num++
		}
		previousSide = ray[i]
	}

	return num
}

func ternary(p1, p2 Plot) RayPoint {
	if p1.Void() && p2.Void() || !p1.Void() && !p2.Void() {
		return Same
	}
	if p1.Void() && !p2.Void() {
		return VoidToFull
	}
	return FullToVoid
}
