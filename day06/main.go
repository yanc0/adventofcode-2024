package main

import (
	. "aoc"
	"bufio"
	"errors"
	"fmt"
	"io"
)

func main() {
	originalMap := parseInput(InputFile())

	m := originalMap.Clone()
	for {
		err := m.NextFrame()
		if err == ErrGuardInfiniteLoop {
			panic("infinite guard loop detected, should not happend on part 1")
		}
		if err == ErrGuardExited {
			break
		}
	}
	fmt.Printf("number of visited tiles: %d\n", len(m.visitedTile))

	fmt.Printf("searching for infinite loop among %d scenarios...\n", len(m.visitedTile))
	countLoopScenarios := 0
	probableTiles := m.visitedTile
	for x := 0; x < len(m.grid); x++ {
		for y := 0; y < len(m.grid); y++ {
			_, isProbableTile := probableTiles[fmt.Sprintf("x:%d,y:%d", x, y)]
			if !isProbableTile {
				continue
			}
			m = originalMap.Clone()
			currentElement := m.GetElement(x, y)
			if currentElement.IsGuard() || currentElement == Obstacle {
				continue
			}
			m.SetElement(x, y, Obstacle)
			if isGuardInALoop(m) {
				countLoopScenarios++
			}
		}
	}
	fmt.Printf("there are %d infinite looped scenarios where to put an obstacle\n", countLoopScenarios)
}

func isGuardInALoop(m Map) bool {
	for {
		err := m.NextFrame()
		if err != nil {
			return err == ErrGuardInfiniteLoop
		}
	}
}

func (m Map) Clone() Map {
	grid := make([][]MapElement, len(m.grid))
	visitedTile := make(map[string]MapElement)

	for i := 0; i < len(m.grid); i++ {
		grid[i] = make([]MapElement, len(m.grid[i]))
		copy(grid[i], m.grid[i])
	}

	return Map{
		grid:        grid,
		visitedTile: visitedTile,
	}
}

func (m Map) SetElement(x int, y int, e MapElement) {
	xOutOfBound := x < 0 || x >= len(m.grid)
	yOutOfBound := y < 0 || y >= len(m.grid)
	if xOutOfBound || yOutOfBound {
		panic("cannot set element on out of bound coordinates")
	}

	m.grid[x][y] = e
}

func parseInput(input io.Reader) Map {
	m := Map{
		grid:        make([][]MapElement, 0),
		visitedTile: make(map[string]MapElement),
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// rules ends, jumps to updates
			break
		}
		row := make([]MapElement, len(line))
		for i, c := range line {
			row[i] = MapElement(c)
		}
		m.grid = append(m.grid, row)
	}

	// visit the initial tile where the guard is
	guardX, guardY := m.FindGuard()
	facing := m.GetElement(guardX, guardY)
	m.MarkVisited(guardX, guardY, facing)

	return m
}

type MapElement rune

const (
	Void             MapElement = 'x'
	Free             MapElement = '.'
	Obstacle         MapElement = '#'
	GuardFacingUP    MapElement = '^'
	GuardFacingRight MapElement = '>'
	GuardFacingDown  MapElement = 'v'
	GuardFacingLeft  MapElement = '<'
)

func (e MapElement) IsGuard() bool {
	return e == GuardFacingUP ||
		e == GuardFacingRight ||
		e == GuardFacingDown ||
		e == GuardFacingLeft
}

type Map struct {
	guardXCache int
	guardYCache int
	grid        [][]MapElement
	visitedTile map[string]MapElement
}

func (m Map) String() string {
	s := ""
	for x := 0; x < len(m.grid); x++ {
		s += string(m.grid[x]) + "\n"
	}
	return s
}

func (m Map) GetElement(x, y int) MapElement {
	xOutOfBound := x < 0 || x >= len(m.grid)
	yOutOfBound := y < 0 || y >= len(m.grid)
	if xOutOfBound || yOutOfBound {
		return Void
	}

	return m.grid[x][y]
}

func (m Map) FindGuard() (x, y int) {
	if m.guardXCache != 0 || m.guardYCache != 0 {
		return m.guardXCache, m.guardYCache
	}
	for x := 0; x < len(m.grid); x++ {
		for y := 0; y < len(m.grid); y++ {
			if m.grid[x][y].IsGuard() {
				return x, y
			}
		}
	}
	panic("unreachable, expected a guard on the map")
}

var ErrGuardInfiniteLoop = errors.New("guard infinite loop detected")
var ErrGuardExited = errors.New("guard has left the map")

func (m Map) MarkVisited(x, y int, facing MapElement) error {
	id := fmt.Sprintf("x:%d,y:%d", x, y)
	previousFacing, alreadyVisited := m.visitedTile[id]
	if alreadyVisited && previousFacing == facing {
		return ErrGuardInfiniteLoop
	}
	m.visitedTile[id] = facing
	return nil
}

func (m Map) IsTileVisited(x, y int) bool {
	id := fmt.Sprintf("x:%d,y:%d", x, y)
	_, visited := m.visitedTile[id]
	return visited
}

func (m Map) NextFrame() error {
	guardX, guardY := m.FindGuard()
	switch m.grid[guardX][guardY] {
	case GuardFacingUP:
		nextTile := m.GetElement(guardX-1, guardY)
		switch nextTile {
		case Void:
			return ErrGuardExited
		case Obstacle:
			m.grid[guardX][guardY] = GuardFacingRight
			return m.NextFrame()
		case Free:
			m.grid[guardX][guardY] = Free
			m.grid[guardX-1][guardY] = GuardFacingUP
			m.guardXCache, m.guardYCache = guardX-1, guardY
			return m.MarkVisited(guardX-1, guardY, GuardFacingUP)
		}
	case GuardFacingRight:
		nextTile := m.GetElement(guardX, guardY+1)
		switch nextTile {
		case Void:
			return ErrGuardExited
		case Obstacle:
			m.grid[guardX][guardY] = GuardFacingDown
			return m.NextFrame()
		case Free:
			m.grid[guardX][guardY] = Free
			m.grid[guardX][guardY+1] = GuardFacingRight
			m.guardXCache, m.guardYCache = guardX, guardY+1
			return m.MarkVisited(guardX, guardY+1, GuardFacingRight)
		}
	case GuardFacingDown:
		nextTile := m.GetElement(guardX+1, guardY)
		switch nextTile {
		case Void:
			return ErrGuardExited
		case Obstacle:
			m.grid[guardX][guardY] = GuardFacingLeft
			return m.NextFrame()
		case Free:
			m.grid[guardX][guardY] = Free
			m.grid[guardX+1][guardY] = GuardFacingDown
			m.guardXCache, m.guardYCache = guardX+1, guardY
			return m.MarkVisited(guardX+1, guardY, GuardFacingDown)
		}
	case GuardFacingLeft:
		nextTile := m.GetElement(guardX, guardY-1)
		switch nextTile {
		case Void:
			return ErrGuardExited
		case Obstacle:
			m.grid[guardX][guardY] = GuardFacingUP
			return m.NextFrame()
		case Free:
			m.grid[guardX][guardY] = Free
			m.grid[guardX][guardY-1] = GuardFacingLeft
			m.guardXCache, m.guardYCache = guardX, guardY-1
			return m.MarkVisited(guardX, guardY-1, GuardFacingLeft)
		}
	}
	return nil
}
