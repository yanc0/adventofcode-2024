package adventofcode

import "fmt"

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("%d|%d", p.X, p.Y)
}

func (p Position) Up() Position {
	return Position{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p Position) Right() Position {
	return Position{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Position) Down() Position {
	return Position{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p Position) Left() Position {
	return Position{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p Position) ExistsOnMatrix(size int) bool {
	xOutOfBand := p.X < 0 || p.X >= size
	yOutOfBand := p.Y < 0 || p.Y >= size
	return !xOutOfBand && !yOutOfBand
}
