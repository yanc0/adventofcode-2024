package main

import (
	. "aoc"
	"fmt"
	"os"
)

func main() {

	parser := NewParser(NewScanner(InputFile()))

	total := int64(0)
	for {
		num, err := parser.Parse()
		if err != nil {
			if err.Error() == "eof" {
				fmt.Println("total:", total)
				os.Exit(0)
			}
			continue
		}
		total += int64(num)
	}
}
