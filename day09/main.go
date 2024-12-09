package main

import (
	adventofcode "aoc"
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
	denseFormat := parseInput(f)
	diskFormat := DiskMap(denseFormat)
	compacted := Compact(diskFormat)

	fmt.Println("disk compacted", IsCompacted(compacted), ", checksum: ", Checksum(compacted))

}

func parseInput(input io.Reader) []int {
	nums := make([]int, 0)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		nums = append(nums, adventofcode.Int(scanner.Text()))
	}
	return nums
}

const EMPTY = -1

func DiskMap(dense []int) []int {
	m := make([]int, 0)
	for i := 0; i < len(dense); i++ {
		isBlock := i%2 == 0
		if isBlock {
			blockID := i / 2
			m = append(m, intRepeat(blockID, dense[i])...)
			continue
		}

		m = append(m, intRepeat(EMPTY, dense[i])...)
	}
	return m
}

func intRepeat(num int, count int) []int {
	r := make([]int, count)
	for i := 0; i < count; i++ {
		r[i] = num
	}
	return r
}

func firstFreeSpaceIndex(diskmap []int) int {
	for i, ch := range diskmap {
		if ch == EMPTY {
			return i
		}
	}
	return -1
}

func lastDataBlockIndex(diskmap []int) int {
	for i := len(diskmap) - 1; i >= 0; i-- {
		if diskmap[i] != EMPTY {
			return i
		}
	}
	return -1
}

func Switch(diskmap []int, from int, to int) {
	tmp := diskmap[to]
	diskmap[to] = diskmap[from]
	diskmap[from] = tmp
}

func Compact(diskmap []int) []int {
	disk := make([]int, len(diskmap))
	copy(disk, diskmap)

	for {
		ldbi := lastDataBlockIndex(disk)
		ffsi := firstFreeSpaceIndex(disk)
		noFreeSpace := ffsi > ldbi
		if noFreeSpace {
			break
		}
		Switch(disk, ldbi, ffsi)
	}

	return disk
}

func Checksum(diskmap []int) int {
	sum := 0
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i] == EMPTY {
			continue
		}
		sum += i * diskmap[i]
	}
	return sum
}

func IsCompacted(diskmap []int) bool {
	mustFindEmpty := false
	for _, v := range diskmap {
		if mustFindEmpty && v != EMPTY {
			return false
		}
		if v == EMPTY {
			mustFindEmpty = true
		}
	}
	return true
}

func DiskMapToString(diskmap []int) string {
	s := ""
	for i := 0; i < len(diskmap); i++ {
		if diskmap[i] == EMPTY {
			s += "."
			continue

		}
		s += fmt.Sprintf("%d", diskmap[i])
	}
	return s
}
