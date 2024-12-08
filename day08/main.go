package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
)

//go:embed *.txt
var input embed.FS

func main() {
	f, err := input.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	antennas := parseInput(f)
	// part 1
	fmt.Println(antennas.AntinodesV1())
	fmt.Printf("there are %d antinodes on the map with model v1\n", antennas.AntinodesV1().Count())

	// part 2
	fmt.Println(antennas.AntinodesV2())
	fmt.Printf("there are %d antinodes on the map with model v2\n", antennas.AntinodesV2().Count())
}

type antinodes [][]bool

func NewAntinodesMap(size int) antinodes {
	atnd := make(antinodes, size)
	for i := 0; i < size; i++ {
		atnd[i] = make([]bool, size)
	}
	return atnd
}

func (a antinodes) String() string {
	s := ""
	for _, line := range a {
		for _, isAntinode := range line {
			if isAntinode {
				s += "#"
				continue
			}
			s += "."
		}
		s += "\n"
	}
	return s
}

func (a antinodes) Count() int {
	count := 0
	for _, line := range a {
		for _, isPosAntinode := range line {
			if isPosAntinode {
				count++
			}
		}
	}
	return count
}

func (a antinodes) Mark(positions ...vec) {
	for _, p := range positions {
		xOutOfBand := p.x < 0 || p.x >= len(a)
		yOutOfBand := p.y < 0 || p.y >= len(a)
		if xOutOfBand || yOutOfBand {
			continue
		}

		a[p.x][p.y] = true
	}
}

type antennas [][]rune

func (a antennas) AntinodesV1() antinodes {
	antinodes := NewAntinodesMap(len(a))

	for _, freq := range a.AllFrequencies() {
		antinodes.Mark(a.CalculateAntinodesForFreqV1(freq)...)
	}

	return antinodes
}

func (a antennas) AntinodesV2() antinodes {
	antinodes := NewAntinodesMap(len(a))

	for _, freq := range a.AllFrequencies() {
		antinodes.Mark(a.CalculateAntinodesForFreqV2(freq)...)
	}

	return antinodes
}

type vec struct {
	x int
	y int
}

func (v vec) Distance(to vec) (dist vec) {
	dist = vec{
		x: to.x - v.x,
		y: to.y - v.y,
	}

	return dist
}

func (v vec) AbsDistance(to vec) (dist vec) {
	dist = v.Distance(to)

	if dist.x < 0 {
		dist.x *= -1
	}
	if dist.y < 0 {
		dist.y *= -1
	}
	return dist
}

func (v vec) Add(toadd vec) (newvec vec) {
	return vec{
		x: v.x + toadd.x,
		y: v.y + toadd.y,
	}
}

func (v vec) Mul(tomul vec) (newvec vec) {
	return vec{
		x: v.x * tomul.x,
		y: v.y * tomul.y,
	}
}

func (v vec) Sub(tosub vec) (newvec vec) {
	return vec{
		x: v.x - tosub.x,
		y: v.y - tosub.y,
	}
}

func (v vec) InsideMatrix(matrixLen int) bool {
	if v.x < 0 || v.x >= matrixLen {
		return false
	}
	if v.y < 0 || v.y >= matrixLen {
		return false
	}
	return true
}

func parseInput(input io.Reader) antennas {
	antennas := make(antennas, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		antennas = append(antennas, []rune(line))
	}
	return antennas
}

func (a antennas) WalkAllPos(f func(freq rune, p vec)) {
	for x := 0; x < len(a); x++ {
		for y := 0; y < len(a); y++ {
			f(a[x][y], vec{x, y})
		}
	}
}

func (a antennas) AllFrequencies() []rune {
	freqmap := make(map[rune]bool)
	freqs := make([]rune, 0)
	a.WalkAllPos(func(freq rune, p vec) {
		if freq != '#' && freq != '.' {
			freqmap[freq] = true
		}
	})

	for freq := range freqmap {
		freqs = append(freqs, freq)
	}

	return freqs
}

func (a antennas) AllFrequencyPos(freq rune) []vec {
	positions := make([]vec, 0)
	a.WalkAllPos(func(f rune, p vec) {
		if freq == f {
			positions = append(positions, p)
		}
	})
	return positions
}

func (a antennas) CalculateAntinodesForFreqV1(freq rune) []vec {
	antinodesPos := make([]vec, 0)
	freqantennas := a.AllFrequencyPos(freq)
	for x := 0; x < len(freqantennas)-1; x++ {
		for y := x + 1; y < len(freqantennas); y++ {
			dist := freqantennas[x].Distance(freqantennas[y])
			antinodesPos = append(antinodesPos, freqantennas[x].Sub(dist))
			antinodesPos = append(antinodesPos, freqantennas[y].Add(dist))
		}
	}
	return antinodesPos
}

func (a antennas) CalculateAntinodesForFreqV2(freq rune) []vec {
	antinodesPos := make([]vec, 0)
	freqantennas := a.AllFrequencyPos(freq)
	for x := 0; x < len(freqantennas)-1; x++ {
		for y := x + 1; y < len(freqantennas); y++ {
			dist := freqantennas[x].Distance(freqantennas[y])
			antinodesPos = append(antinodesPos, freqantennas[x], freqantennas[y])
			for i := 1; i < len(a); i++ {
				backantinode := freqantennas[x].Sub(dist.Mul(vec{i, i}))
				frontantinode := freqantennas[y].Add(dist.Mul(vec{i, i}))
				if !backantinode.InsideMatrix(len(a)) && !frontantinode.InsideMatrix(len(a)) {
					break
				}
				antinodesPos = append(antinodesPos, backantinode, frontantinode)
			}
		}
	}
	return antinodesPos
}
