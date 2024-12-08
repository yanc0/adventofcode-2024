package main

import (
	"fmt"
	"log"
	"testing"
)

func TestInput(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	antennas := parseInput(f)
	_ = antennas
}

func TestAntennasAntinodes(t *testing.T) {
	f, err := input.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	antennas := parseInput(f)
	antinodes := NewAntinodesMap(len(antennas))
	antinodes.Mark(antennas.CalculateAntinodesForFreqV1('0')...)
	fmt.Println(antinodes.Count())
}
